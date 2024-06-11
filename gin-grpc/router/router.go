package router

import (
	"gin-grpc/controllers"
	"gin-grpc/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用 CORS 中间件
	r.Use(middlewares.CORSMiddleware())

	// 配置 REST 路由
	r.GET("/test", controllers.Test)

	// 配置 gRPC-Gateway 路由
	r.GET("/grpc/test", controllers.GRPCGateway)

	return r
}
