package logic

import (
	"context"

	"github.com/yaqia77/golang-gin-demo/mysql/user_gorm/models"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/rpc/internal/svc"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/rpc/types/user"

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

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (pd *user.UserCreateesponse, err error) {
	// todo: add your logic here and delete this line
	pd = new(user.UserCreateesponse)
	var model models.UserModel
	err = l.svcCtx.DB.Take(&model, "username = ?", in.Username).Error
	if err == nil {
		pd.Err = "该用户已存在"
		return
	}
	model = models.UserModel{
		Username: in.Username,
		Password: in.Password,
	}

	err = l.svcCtx.DB.Create(&model).Error
	if err != nil {
		logx.Error(err)
		pd.Err = err.Error()
		err = nil
		return

	}
	pd.UserId = uint32(model.ID)

	return &user.UserCreateesponse{}, nil
}
