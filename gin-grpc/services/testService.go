package services

import (
	"context"
	"gin-grpc/proto"
)

// TestServiceServer是TestService服务的服务器API
type TestServiceServer struct {
	//proto.UnimplementedTestServiceServer是TestService服务的未实现接口
	proto.UnimplementedTestServiceServer
}

// SayHello方法实现了TestService服务的SayHello方法
// 它接收一个HelloRequest参数，返回一个HelloResponse参数和一个error
// 这里我们简单地返回一个HelloResponse，其中包含一个问候语和请求的名字
// 注意，我们使用了指针类型作为返回值，这是因为我们需要返回可选的错误信息
//ctx context.Context是请求的上下文，req *proto.HelloRequest是客户端请求的HelloRequest参数
func (s *TestServiceServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Message: "Hello " + req.Name}, nil
}
