package router

import (
	"gin-captcha/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/captcha", handlers.GenerateCaptcha)
		api.GET("/captcha/:captcha_id", handlers.CaptchaServe)
		api.POST("/verify", handlers.VerifyCaptcha)
	}
}
