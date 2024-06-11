package router

import (
	"gin-logger/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/users", handlers.GetUsers)
	}
}

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/users", handlers.GetUsers)
	}
}
