package logic

import (
	"context"

	"github.com/cache-hhhhhelp/douyin-user/internal/svc"
	"github.com/cache-hhhhhelp/douyin-user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.UserReq) (*user.UserResp, error) {

	result, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &user.UserResp{
		Id:   result.UserId,
		Name: result.Username,
	}, nil
}
