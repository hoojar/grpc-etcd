package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	pb "grpc-etcd/pb"
)

var etcdClient *clientv3.Client

//etcd解析器
type etcdResolver struct {
	etcdAddr   string
	clientConn resolver.ClientConn
}

//初始化一个etcd解析器
func newResolver(etcdAddr string) resolver.Builder {
	return &etcdResolver{etcdAddr: etcdAddr}
}

func (r *etcdResolver) Scheme() string {
	return "etcd"
}

//watch有变化以后会调用
func (r *etcdResolver) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("Resolve Now")
}

//解析器关闭时调用
func (r *etcdResolver) Close() {
	log.Println("Close")
}

//构建解析器 grpc.Dial()同步调用
func (r *etcdResolver) Build(target resolver.Target, clientConn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error
	if etcdClient == nil { //构建etcd client
		etcdClient, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(r.etcdAddr, ";"),
			DialTimeout: 15 * time.Second,
		})

		if err != nil {
			fmt.Printf("连接etcd失败：%s\n", err)
			return nil, err
		}
	}

	r.clientConn = clientConn
	go r.watch("/" + target.Scheme + "/" + target.Endpoint + "/")
	return r, nil
}

//监听etcd中某个key前缀的服务地址列表的变化
func (r *etcdResolver) watch(keyPrefix string) {
	var addrList []resolver.Address //初始化服务地址列表
	fmt.Println("service name key prefix: ", keyPrefix)
	resp, err := etcdClient.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("获取服务地址列表失败：", err)
	} else {
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(resp.Kvs[i].Key), keyPrefix)})
		}
	}
	_ = r.clientConn.UpdateState(resolver.State{Addresses: addrList})

	//监听服务地址列表的变化
	rch := etcdClient.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
			if ev.Type == 0 { //ETCD CMD: PUT
				if !r.addressExists(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})
					_ = r.clientConn.UpdateState(resolver.State{Addresses: addrList})
				}
				fmt.Printf("watch event put-current: %#v \n", string(ev.Kv.Value))
			} else if ev.Type == 1 { //ETCD CMD: DELETE
				if s, ok := r.addressRemove(addrList, addr); ok {
					_ = r.clientConn.UpdateState(resolver.State{Addresses: s})
				}
				fmt.Printf("watch event delete-current: %#v \n", string(ev.Kv.Value))
			}
		}
	}
}

func (r *etcdResolver) addressExists(l []resolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func (r *etcdResolver) addressRemove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func main() {
	var svrName = flag.String("svrName", "greetSvr", "service name")                 //服务名称
	var etcdHost = flag.String("etcdHost", "127.0.0.1:2379", "etcd service address") //etcd的地址

	flag.Parse()                //参数分析
	r := newResolver(*etcdHost) //注册etcd解析器
	resolver.Register(r)        //负载分析

	conn, err := grpc.Dial(
		r.Scheme()+"://author/"+*svrName,
		grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`),
		grpc.WithInsecure()) //客户端连接服务器(负载均衡：轮询) 会同步调用r.Build()
	if err != nil {
		fmt.Println("连接服务器失败：", err)
	}
	defer conn.Close()

	/* 获得grpc句柄并执行GRPC调用 */
	grpcClient := pb.NewRpcClient(conn)
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		resp1, err := grpcClient.Hello(context.Background(), &pb.RpcReq{Msg: "woods zhang"})
		if err != nil {
			fmt.Println("Hello 调用失败：", err.Error())
		} else {
			fmt.Printf("Hello 响应：%s 来自：%s\n", resp1.Text, resp1.From)
		}

		resp2, err := grpcClient.Md5(context.Background(), &pb.RpcReq{Msg: "woods.zhang"})
		if err != nil {
			fmt.Println("Md5 调用失败：", err.Error())
		} else {
			fmt.Printf("Md5 响应：%s 来自：%s\n", resp2.Text, resp2.From)
		}

		fmt.Println()
	}
}
