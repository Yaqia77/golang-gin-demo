package logic

import (
	"context"

	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/svc"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/types"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	response, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId: uint32(req.ID),
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		UserId:   uint(response.UserId),
		Username: response.Username,
	}, nil
}
