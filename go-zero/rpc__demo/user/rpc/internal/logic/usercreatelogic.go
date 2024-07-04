package logic

import (
	"context"

	"github.com/yaqia77/golang-gin-demo/rpc__demo/user/rpc/internal/svc"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (*user.UserCreateesponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserCreateesponse{}, nil
}
