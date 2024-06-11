package controllers

import (
	"context"
	"gin-grpc/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

// GRPCGateway是一个测试函数，用于从Gin框架调用gRPC服务
func GRPCGateway(c *gin.Context) {
	// 连接gRPC服务

	// 注意：这里的连接方式是不安全的，仅用于测试，在生产环境中，请使用安全的连接方式
	//DialContext方法用于连接gRPC服务
	//WithInsecure()方法用于创建不安全的连接
	//"localhost:50051"是gRPC服务的地址
	conn, err := grpc.DialContext(context.Background(), "localhost:50051", grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//关闭连接
	defer conn.Close()
	// 调用gRPC服务
	//NewTestServiceClient是由proto文件生成的客户端
	client := proto.NewTestServiceClient(conn)
	//调用SayHello方法
	//context.Background()是上下文，用于控制连接的生命周期
	res, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "Gin"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": res.Message})
}
