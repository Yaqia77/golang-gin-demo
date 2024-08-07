package logic

import (
	"context"

	"github.com/yaqia77/golang-gin-demo/apps/user/api_swagger/internal/svc"
	"github.com/yaqia77/golang-gin-demo/apps/user/api_swagger/internal/types"
	"github.com/yaqia77/golang-gin-demo/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	auth := l.svcCtx.Config.Auth
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   1,
		Username: "user1",
		Role:     1,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return "", err
	}
	return token, nil
}
