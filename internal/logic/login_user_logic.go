package logic

import (
	"context"

	"github.com/cache-hhhhhelp/douyin-user/internal/svc"
	"github.com/cache-hhhhhelp/douyin-user/types/user"

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

func (l *LoginUserLogic) LoginUser(in *user.RegisterOrLoginReq) (*user.DataResp, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}

	err = PasswordVerify(in.Password, result.Password)
	if err != nil {
		return nil, err
	}

	return &user.DataResp{UserId: result.UserId}, nil
}
