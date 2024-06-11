package main

import (
	"gin-cors/router"
)

//cors跨域
//cors安装
//go get -u github.com/gin-contrib/cors

//选择 OPTIONS 方法。
//点击 "Headers" 选项卡，添加 Origin 头，值为 https://example.com。
//发送 OPTIONS 请求，响应头中会包含 Access-Control-Allow-Origin 头，值为 https://example.com。
//响应头 Access-Control-Allow-Origin 是 https://example.com。
//响应头 Access-Control-Allow-Methods 包含 GET, POST, PUT, PATCH, DELETE, OPTIONS。

func main() {
	r := router.SetupRouter()

	r.Run(":8080")
}
