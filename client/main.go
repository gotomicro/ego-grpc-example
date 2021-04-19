package main

import (
	"context"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/core/elog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/gotomicro/ego-component/eetcd"
	"github.com/gotomicro/ego-component/eetcd/examples/helloworld"
	"github.com/gotomicro/ego-component/eetcd/registry"
)

func main() {
	if err := ego.New().Invoker(
		invokerGrpc,
		callGrpc,
	).Run(); err != nil {
		elog.Error("startup", elog.FieldErr(err))
	}
}

var grpcComp helloworld.GreeterClient

func invokerGrpc() error {
	registry.Load("registry").Build(registry.WithClientEtcd(eetcd.Load("etcd").Build()))
	grpcConn := egrpc.Load("grpc.test").Build()
	grpcComp = helloworld.NewGreeterClient(grpcConn.ClientConn)
	return nil
}

func callGrpc() error {
	var header metadata.MD
	var trailer metadata.MD
	_, err := grpcComp.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "i am client",
	},grpc.Header(&header),grpc.Trailer(&trailer))
	if err != nil {
		return err
	}

	//spew.Dump(header)
	//spew.Dump(trailer)

	_, err = grpcComp.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "error",
	})
	if err != nil {
		return err
	}
	return nil
}