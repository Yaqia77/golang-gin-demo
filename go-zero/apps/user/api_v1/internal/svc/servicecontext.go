package svc

import (
	"github.com/yaqia77/golang-gin-demo/apps/user/api_v1/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
