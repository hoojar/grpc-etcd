# grpc-etcd
grpc use etcd implement the registry and discovery

# How to test
```bash
make all

grpc-etcd-server -svrAddr=127.0.0.1:3000
grpc-etcd-server -svrAddr=192.168.0.2:3001
grpc-etcd-server -svrAddr=192.168.0.2:3003

grpc-etcd-client
```
