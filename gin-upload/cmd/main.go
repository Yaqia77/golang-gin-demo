package main

import (
	"gin-upload/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetRouter(r)

	r.Run(":8080")
}

//处理文件上传
//1. 先设置路由
//2. 接收文件
//3. 保存文件
//4. 返回响应
