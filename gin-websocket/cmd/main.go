package main

import (
	"gin-websocket/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.SetupRoutes(r)

	r.Run(":8080")

}

//WebSocket是HTML5一种新的协议。它实现了浏览器与服务器全双工通信（full-duplex communication）的功能。
//WebSocket协议在建立连接时，需要先握手，然后才能进行通信。握手时，服务器和客户端都需要发送一个HTTP请求，并指定一种协议。
//应用场景：
//1.实时通信，如聊天室、股票行情、游戏
//2.实时数据推送，如股票行情、天气预报、新闻推送
//3.在线多人协作，如在线文档协作、在线视频会议

//第三方库github.com/gorilla/websocket来实现
//go get -u github.com/gorilla/websocket

//HTML文件在浏览器中打开。点击“Connect”按钮建立WebSocket连接，然后点击“Send Message”按钮发送消息，你将看到服务器回显的消息。

//通过这种方式，你可以在Gin项目中集成WebSocket，并使用WebSocket客户端或简单的网页进行测试。
