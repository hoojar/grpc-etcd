TARGET=grpc-etcd
outPath=./pb
binPath=${GOPATH}/bin/
protoPath=-I ./pb
all: clean build

clean:
	rm -rf $(TARGET)-client
	rm -rf $(TARGET)-server

build:
	go build -o ${TARGET}-client client.go
	go build -o $(TARGET)-server server.go

proto:
	echo "test"
	protoc ${protoPath} --go_out=$(outPath) --plugin=protoc-gen-go=${binPath}protoc-gen-go RpcApi.proto
	protoc ${protoPath} --go-grpc_out=$(outPath) --plugin=protoc-gen-go-grpc=${binPath}protoc-gen-go-grpc RpcApi.proto