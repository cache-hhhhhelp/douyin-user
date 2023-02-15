package logic

import (
	"context"
	"microservice/user/common/xerr"
	"microservice/user/rpc/models"

	"microservice/user/rpc/internal/svc"
	"microservice/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterUserLogic) RegisterUser(in *user.RegisterOrLoginRequest) (*user.DataResponse, error) {
	var userRes models.User
	err := l.svcCtx.DbEngine.Where("username = ?", in.Username).First(&userRes)

	if err.Error == nil {
		return nil, xerr.NewErrCodeMsg(103, "用户名已存在")
	}

	userNew := models.User{
		Username: in.Username,
		Password: in.Password,
	}

	res := l.svcCtx.DbEngine.Create(&userNew)

	if res.Error != nil {
		userRes := user.DataResponse{
			StatusCode: 100,
		}
		userRes.SetStatusMsg(res.Error.Error())
		return &userRes, nil
	} else {
		userRes := user.DataResponse{
			StatusCode: 0,
			UserId:     int64(userNew.Model.ID),
		}
		userRes.SetStatusMsg("success")
		return &userRes, nil
	}

}
