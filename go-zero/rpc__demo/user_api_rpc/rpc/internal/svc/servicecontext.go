package svc

import (
	"github.com/yaqia77/golang-gin-demo/pkg/init_gorm"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/models"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&models.UserModel{})
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
