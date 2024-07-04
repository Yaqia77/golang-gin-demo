package logic

import (
	"context"
	"errors"

	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/rpc/internal/svc"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/rpc/types/user"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_gorm/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	var model models.UserModel
	err := l.svcCtx.DB.Take(&model, in.UserId).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user.UserInfoResponse{
		UserId:   uint32(model.ID),
		Username: model.Username,
	}, nil
}
