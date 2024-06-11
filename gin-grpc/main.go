package main

import (
	"gin-grpc/proto"
	"gin-grpc/router"
	"gin-grpc/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 启动 gRPC 服务器
	grpcServer := grpc.NewServer()
	testService := &services.TestServiceServer{}
	proto.RegisterTestServiceServer(grpcServer, testService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 启动 Gin 服务器
	r := router.SetupRouter()
	r.Run(":8080")
}

//gRPC 和 HTTP/REST 接口的调用：
//1. 启动 gRPC 服务器
//2. 启动 Gin 服务器
//3. 调用 gRPC 接口：http://localhost:8080/grpc/test
//4. 调用 REST 接口：http://localhost:8080/test

//测试REST 接口：
//http://localhost:8080/test
//响应：{"message": "Hello, World!"}

//测试gRPC接口：
//http://localhost:8080/grpc/test
//响应：{"message": "Hello  Gin!"}

//集成 gRPC 和 Gin 框架可以让你在同一个应用程序中同时使用 HTTP/REST 和 gRPC 服务。
