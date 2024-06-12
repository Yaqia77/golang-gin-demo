package mq

import (
	"github.com/streadway/amqp"
)

// Consume receives a message from the RabbitMQ queue
// Consume是从RabbitMQ队列接收消息的函数，它接收一个*amqp.Channel和一个队列名作为参数，并返回一个字符串和一个错误。
func Consume(ch *amqp.Channel, queueName string) (string, error) {
	msgs, err := ch.Consume(
		queueName, // queue 队列
		"",        // consumer 消费者
		true,      // auto-ack 自动确认
		false,     // exclusive 独占
		false,     // no-local 不本地
		false,     // no-wait 不等待
		nil,       // args 参数
	)
	if err != nil {
		return "", err
	}
	//从msgs通道中接收消息，并将其赋值给msg变量。
	msg := <-msgs
	return string(msg.Body), nil
}
