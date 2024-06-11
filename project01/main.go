package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	//创建路由
	r := gin.Default()
	//设置路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	//启动服务
	r.Run(":8080")

	
}