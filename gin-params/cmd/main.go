package main

import (
	"gin-params/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	router.SerRouter(r)

	r.Run(":8080")
}

//处理请求参数，包括路径参数、查询参数和表单参数。
//gin框架支持多种方式处理请求参数，包括路径参数、查询参数、表单参数、JSON参数、XML参数等。