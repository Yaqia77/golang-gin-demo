package router

import (
	"gin-cors/controllers"
	"gin-cors/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//使用cors中间件
	r.Use(middlewares.CORSMiddleware())

	//注册路由
	r.GET("/test", controllers.Test)
	return r
}
