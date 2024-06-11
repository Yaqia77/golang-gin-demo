package main

import (
	"gin-middleware/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	router.SetRouter(r)

	r.Run(":8080")
}

//中间件可以用于日志记录、身份验证、错误处理等。
