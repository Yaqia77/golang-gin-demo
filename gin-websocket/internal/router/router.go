package router

import (
	"github.com/gin-gonic/gin"
	"gin-websocket/internal/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	{
		api.GET("/ws",handlers.WebSocketHandler)
	}
}
