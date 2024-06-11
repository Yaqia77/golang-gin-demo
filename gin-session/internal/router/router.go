package router

import (
	"gin-session/internal/handlers"
	"gin-session/internal/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	api := r.Group("/api/v1")
	{
		api.POST("/login", handlers.Login)
		api.GET("/logout", handlers.Logout)
		api.GET("/profile", middleware.AuthMiddleware(), handlers.Profile)
	}
}
