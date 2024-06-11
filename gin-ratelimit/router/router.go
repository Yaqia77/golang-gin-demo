package router

import(
	"github.com/gin-gonic/gin"
	"gin-ratelimit/controllers"
	"gin-ratelimit/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用限流中间件
	r.Use(middlewares.RateLimitMiddleware())

	// 配置路由
	r.GET("/test", controllers.Test)

	return r

}