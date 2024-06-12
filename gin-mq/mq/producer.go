package mq

import (
    "github.com/streadway/amqp"
)

// Produce sends a message to the RabbitMQ queue
//Produce是将消息发送到RabbitMQ队列的函数
//amqp.Channel是RabbitMQ的通道，queueName是队列的名称，body是要发送的消息
func Produce(ch *amqp.Channel, queueName, body string) error {
    err := ch.Publish(
        "",       // 交换
        queueName, // 队列名
        false,    // 强制
        false,    // 立即
				// 这里的Publishing结构体是用来设置消息属性的，比如ContentType、Body等等。
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    return err
}
