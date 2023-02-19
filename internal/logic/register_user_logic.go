package logic

import (
	"context"
	"github.com/cache-hhhhhelp/douyin-user/internal/model"

	"github.com/cache-hhhhhelp/douyin-user/internal/svc"
	"github.com/cache-hhhhhelp/douyin-user/types/user"

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

func (l *RegisterUserLogic) RegisterUser(in *user.RegisterOrLoginReq) (*user.DataResp, error) {

	pwd, err := PasswordHash(in.Password)
	if err != nil {
		return nil, err
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: in.Username,
		Password: pwd,
	})
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &user.DataResp{UserId: id}, nil
}
