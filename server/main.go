package main

import (
	"context"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego-component/eetcd"
	_ "github.com/gotomicro/ego-component/eetcd/conf"
	"github.com/gotomicro/ego-component/eetcd/examples/helloworld"
	"github.com/gotomicro/ego-component/eetcd/registry"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egovernor"
	"github.com/gotomicro/ego/server/egrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	if err := ego.New().
		Invoker(invoker).
		Registry(EtcdRegistry).
		Serve(
			newGRPCServer(),
			newGovernorServer(),
		).Run(); err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}

var (
	EtcdClient   *eetcd.Component
	EtcdRegistry *registry.Component
)

func invoker() error {
	EtcdRegistry = registry.Load("registry").Build(registry.WithClientEtcd(eetcd.Load("etcd").Build()))
	return nil
}

func newGRPCServer() server.Server {
	server := egrpc.Load("server.grpc").Build()
	helloworld.RegisterGreeterServer(server.Server, &Greeter{server: server})
	return server
}

func newGovernorServer() server.Server {
	return egovernor.Load("server.governor").Build()
}

type Greeter struct {
	server *egrpc.Component
}

func (g Greeter) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if request.Name == "error" {
		return nil, status.Error(codes.Unavailable, "error")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		elog.Debug("metadata info",zap.Any("meta",md))
	}

	return &helloworld.HelloReply{
		Message: "Hello EGO, I'm " + g.server.Address(),
	}, nil
}
