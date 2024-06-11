package main

import (
	"gin-ratelimit/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}

//使用 Postman 或浏览器访问 http://localhost:8080/test。

//在限流中间件中设置每分钟允许的最大请求数为 5。
//在一分钟内连续发送超过 5 个请求，第 6 个请求应该返回状态码 429 Too Many Requests，并显示错误信息 {"error": "Too many requests"}。
//在 Gin 应用程序中实现基本的限流功能，防止滥用和过载。
