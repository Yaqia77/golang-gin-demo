package main

import (
	"fmt"
	"gin-logger/internal/router"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = f

	// 使用Logger中间件，默认使用Logger()函数
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义日志格式
		return fmt.Sprintf("[GIN] %v | %3d | %13v |%15s | %-7s %#v\n%s",
			param.TimeStamp.Format(time.RFC1123),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))

	// 使用Recovery中间件，防止panic
	r.Use(gin.Recovery())

	router.SetupRoutes(r)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

//使用Gin框架的Logger中间件可以记录HTTP请求的信息，这对调试和监控非常有用。
//默认情况下，Gin框架的Logger中间件会将日志输出到标准输出，但也可以通过gin.DefaultWriter参数指定日志输出的位置。
