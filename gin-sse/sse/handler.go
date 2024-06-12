package sse

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// HandleSSE handles Server-Sent Events
func HandleSSE(c *gin.Context) {
	// Set headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	// Flush the headers
	//刷新头部
	c.Writer.Flush()

	//发送事件
	//select是Go语言的一种控制结构，用于在多个通道上进行选择。
	//在这里，我们使用select来监听两个通道：一个是客户端连接断开的通道，另一个是2秒钟过去的通道。
	//如果客户端连接断开，则退出循环；如果2秒钟过去了，则发送一个事件。
	for {
		select {
		case <-c.Request.Context().Done():
			// Done()方法返回一个通道，当请求处理完成或被取消时，该通道会关闭。
			// Client disconnected
			fmt.Println("Client disconnected")
			return
		case <-time.After(2 * time.Second):
			// Send an event
			fmt.Fprintf(c.Writer, "data: %s\n\n", time.Now().Format(time.RFC3339))
			c.Writer.Flush()
		}
	}
}
