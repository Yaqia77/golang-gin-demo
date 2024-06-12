package main

//Server-Sent Events 集成到 Gin 框架中

import (
	"net/http"

	"gin-sse/sse"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//HandleSSE 处理 SSE 请求
	router.GET("/events", sse.HandleSSE)
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(":8080")
}
