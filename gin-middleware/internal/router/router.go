package router

import (
	"gin-middleware/internal/handlers"
	"gin-middleware/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.Use(middleware.LogginMiddleware())

	api := r.Group("/api/v1")

	{
		//公开路由
		api.GET("/public", handlers.PublicEndpoint)

		//私有路由
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/user", handlers.UserEndpoint)
		}
	}

}
