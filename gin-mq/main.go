package main

import (
	"log"
	"net/http"

	"gin-mq/mq"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {
	router := gin.Default()

	// RabbitMQ connection
	// Dial是建立连接的函数，参数是连接地址，返回值是连接对象，如果连接失败会返回错误。
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	// Channel是信道，每个连接都有唯一的信道，通过信道可以完成消息的发送、接收、确认等操作。
	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	// QueueDeclare是声明队列的函数，参数是队列名、是否持久化、是否自动删除、是否独占、是否阻塞、其他参数，返回值是队列对象和错误。
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Route to send a message
	router.POST("/send", func(c *gin.Context) {
		body := c.PostForm("message")
		//Produce是发送消息的函数，参数是信道、队列名、消息内容，返回值是错误。
		err = mq.Produce(ch, q.Name, body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed to send message"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "message sent"})
	})

	// Route to receive a message
	router.GET("/receive", func(c *gin.Context) {
		//Consume是接收消息的函数，参数是信道、队列名，返回值是消息内容和错误。
		msg, err := mq.Consume(ch, q.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed to receive message"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": msg})
	})

	router.Run(":8080")
}
