package svc

import (
	"github.com/yaqia77/golang-gin-demo/apps/user/rpc/userclient"
	"github.com/yaqia77/golang-gin-demo/apps/video/api/internal/config"
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
