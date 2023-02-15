package logic

import (
	"context"
	"microservice/user/rpc/internal/utils"
	"microservice/user/rpc/models"

	"microservice/user/rpc/internal/svc"
	"microservice/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginUserLogic) LoginUser(in *user.RegisterOrLoginRequest) (*user.DataResponse, error) {
	var userNew models.User
	err := l.svcCtx.DbEngine.Where("username = ?", in.Username).First(&userNew)

	if utils.CheckPassword(in.Password, userNew.Password) {
		return &user.DataResponse{
			UserId: int64(userNew.ID),
		}, nil
	}
	return nil, err.Error

}
