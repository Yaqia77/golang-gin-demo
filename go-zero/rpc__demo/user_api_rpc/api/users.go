package main

import (
	"flag"
	"fmt"

	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/config"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/handler"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}