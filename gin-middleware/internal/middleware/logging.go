package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LogginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//记录请求开始时间
		start := time.Now()

		//处理请求
		c.Next()

		//记录日志
		duration := time.Since(start)
		// 日志格式
		log.Printf("Request: %s %s | Status: %d | Duration: %v",
			c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}

}
