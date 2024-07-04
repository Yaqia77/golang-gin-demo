package svc

import (
	"github.com/yaqia77/golang-gin-demo/apps/user/api_swagger/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
