package main

import (
	"gin-jwt/internal/auth"
	"gin-jwt/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 注册路由
	authMiddleware := auth.NewAuthMiddleware()

	r.GET("/login", authMiddleware.LoginHandler)
	r.GET("protected", authMiddleware.AuthRequired(), handlers.ProtectedHandler)

	r.Run(":8080")
}
