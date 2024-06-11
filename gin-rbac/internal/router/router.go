package router

import (
	"gin-rbac/internal/handlers"
	"gin-rbac/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/public", handlers.PublicEndpoint)

		//需要身份验证的路由
		protected := api.Group("/")
		//middleware.AuthMiddleware() 身份验证中间件
		protected.Use(middleware.AuthMiddleware())
		{
			//需要admin权限的路由
			protected.GET("/user", handlers.UserEndpoint)
			protected.GET("/admin", middleware.RBACMiddleware("admin"), handlers.AdminEndpoint)
		}
	}
}
