package logic

import (
	"context"
	"encoding/json"

	"github.com/yaqia77/golang-gin-demo/apps/user/api_jwt/internal/svc"
	"github.com/yaqia77/golang-gin-demo/apps/user/api_jwt/internal/types"

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

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("user_id").(json.Number)
	uid, _ := userId.Int64()
	username := l.ctx.Value("username").(string)

	return &types.UserInfoResponse{
		UserId:   uint(uid),
		Username: username,
	}, nil
}
