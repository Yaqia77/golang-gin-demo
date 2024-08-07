// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_gorm/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserCreateRequest = user.UserCreateRequest
	UserCreateesponse = user.UserCreateesponse
	UserInfoRequest   = user.UserInfoRequest
	UserInfoResponse  = user.UserInfoResponse

	User interface {
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateesponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUser) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateesponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserCreate(ctx, in, opts...)
}
