package logic

import (
	"context"
	"microservice/internal/model"

	"microservice/internal/svc"
	"microservice/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetManyUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetManyUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetManyUserLogic {
	return &GetManyUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func toUserResp(v []model.User) []*user.UserResp {
	ret := make([]*user.UserResp, len(v))
	for i := 0; i < len(v); i++ {
		ret[i] = &user.UserResp{
			Id:   v[i].UserId,
			Name: v[i].Username,
		}
	}
	return ret
}

func (l *GetManyUserLogic) GetManyUser(in *user.ManyUserReq) (*user.ManyUserResp, error) {

	result, err := l.svcCtx.UserModel.FindMany(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &user.ManyUserResp{Users: toUserResp(result)}, nil
}
