package main

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	proto "grpc-etcd/pb"
)

//rpc服务接口
type rpcServer struct {
	svrAddress string
	proto.UnimplementedRpcServer
}

func (gs *rpcServer) Hello(ctx context.Context, req *proto.RpcReq) (*proto.RpcRes, error) {
	fmt.Printf("Hello 调用传参值: %s\n", req.Msg)
	return &proto.RpcRes{
		Text: "Hello -> " + req.GetMsg(),
		From: gs.svrAddress,
	}, nil
}

func (gs *rpcServer) Md5(ctx context.Context, req *proto.RpcReq) (*proto.RpcRes, error) {
	fmt.Printf("Md5 调用传参值: %s\n", req.GetMsg())
	return &proto.RpcRes{
		Text: req.Msg + " Md5 val is -> " + fmt.Sprintf("%x", md5.Sum([]byte(req.Msg))),
		From: gs.svrAddress,
	}, nil
}

/**
将服务地址注册到etcd中,与etcd建立长连接，并保证连接不断(心跳检测)
*/
func etcdRegister(cli *clientv3.Client, svrName, svrAddr string, ttl int64) {
	ticker := time.NewTicker(time.Second * time.Duration(ttl))
	go func() {
		key := "/etcd/" + svrName + "/" + svrAddr
		for {
			resp, err := cli.Get(context.Background(), key)
			if err != nil {
				fmt.Printf("获取服务地址失败：%s", err)
			} else if resp.Count == 0 { //尚未注册
				//创建租约
				leaseResp, err := cli.Grant(context.Background(), ttl)
				if err != nil {
					fmt.Printf("创建租期失败：%s\n", err)
					continue
				}

				//将服务地址注册到etcd中
				_, err = cli.Put(context.Background(), key, svrAddr, clientv3.WithLease(leaseResp.ID))
				if err != nil {
					fmt.Printf("注册服务失败：%s", err)
					continue
				}

				//建立长连接
				ch, err := cli.KeepAlive(context.Background(), leaseResp.ID)
				if err != nil {
					fmt.Printf("建立长连接失败：%s\n", err)
					continue
				}

				//清空keepAlive返回的channel
				go func() {
					for {
						<-ch //这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步。
						time.Sleep(1)
					}
				}()
			}
			<-ticker.C //执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。
		}
	}()
}

/**
取消注册
*/
func etcdUngegister(cli *clientv3.Client, svrName, svrAddr string) {
	_, err := cli.Delete(context.Background(), "/etcd/"+svrName+"/"+svrAddr)
	if err != nil {
		log.Print("Etcd delete data error: ", err)
	}
}

/* 监听信号处理 */
func disposeSignal(cli *clientv3.Client, svrName, svrAddr string) {
	var ch = make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		var signalData = <-ch //执行该语句时将会阻塞，直到接收到数据并赋值给 signalData 变量
		etcdUngegister(cli, svrName, svrAddr)
		if i, ok := signalData.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()
}

func main() {
	var (
		svrAddr  = flag.String("svrAddr", ":3000", "service address")                 //服务IP:端口
		svrName  = flag.String("svrName", "greetSvr", "service name")                 //服务名称
		etcdAddr = flag.String("etcdAddr", "127.0.0.1:2379", "register etcd address") //etcd的地址
	)
	flag.Parse()

	/*	连接ETCD	*/
	var err error
	var etcdCli *clientv3.Client
	var etcdConfig = clientv3.Config{
		Endpoints:   strings.Split(*etcdAddr, ";"),
		DialTimeout: 15 * time.Second,
	}
	if etcdCli, err = clientv3.New(etcdConfig); err != nil {
		fmt.Printf("连接etcd失败：%s\n", err.Error())
		return
	}

	/* 检测ETCD状态是否正常 */
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = etcdCli.Status(ctx, etcdConfig.Endpoints[0])
	if err != nil {
		fmt.Printf("etcd错误状态：%s\n", err.Error())
		return
	}

	/* 创建服务并监听网络 */
	listener, err := net.Listen("tcp", *svrAddr) //监听网络
	if err != nil {
		fmt.Println("监听网络失败：", err)
		return
	}
	defer listener.Close()

	/* 创建GRPC服务 */
	grpcSvr := grpc.NewServer()                                        //创建grpc句柄
	defer grpcSvr.GracefulStop()                                       //安全停止RPC服务
	proto.RegisterRpcServer(grpcSvr, &rpcServer{svrAddress: *svrAddr}) //将rpcServer结构体注册到grpc服务中
	etcdRegister(etcdCli, *svrName, *svrAddr, 5)                       //将服务地址注册到etcd中
	fmt.Printf("greeting server address: %s\n", *svrAddr)
	if err = grpcSvr.Serve(listener); err != nil {
		fmt.Println("GRPC 服务监听异常：", err)
		return
	}
}
