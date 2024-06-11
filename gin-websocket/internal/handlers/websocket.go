package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 定义一个全局的upgrader
// Upgrader是websocket的协议升级器，它提供了一个Upgrade方法，用于将HTTP连接升级为WebSocket连接。
var upGrader = websocket.Upgrader{
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocketHandler 处理websocket连接
func WebSocketHandler(c *gin.Context) {
	//升级websocket连接
	//Upgrade方法将HTTP连接升级为WebSocket连接，并返回一个新的*websocket.Conn类型。
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	//
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %v", err)
		return
	}
	//用于确保在函数退出时关闭WebSocket连接
	defer conn.Close()

	//循环读取消息并写入消息
	for {
		//读取消息
		//ReadMessage方法从WebSocket连接读取一条消息，并返回三个值：消息类型（int），消息（[]byte）和错误（error）。
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		//打印消息
		log.Printf("recv: %s", message)
		//写入消息
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
