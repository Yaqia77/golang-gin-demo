package main

import (
	"flag"
	"fmt"

	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_gorm/rpc/internal/config"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_gorm/rpc/internal/server"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_gorm/rpc/internal/svc"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_gorm/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
