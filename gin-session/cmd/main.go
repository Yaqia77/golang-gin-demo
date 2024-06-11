package main

import (
	"gin-session/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.SetRouter(r)

	r.Run(":8080")
}

//会话管理
//1. 基于cookie的会话管理
//2. 基于token的会话管理
