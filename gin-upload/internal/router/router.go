package router

import (
	"gin-upload/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/upload", handlers.UploadFile)
	}
}
