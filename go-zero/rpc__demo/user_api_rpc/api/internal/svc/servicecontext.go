package svc

import (
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/config"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
