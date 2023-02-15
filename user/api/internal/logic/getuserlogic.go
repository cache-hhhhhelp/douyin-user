package logic

import (
	"context"
	"microservice/user/rpc/userclient"

	"microservice/user/api/internal/svc"
	"microservice/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserReply, err error) {
	info, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.UserRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return &types.UserReply{
			StatusCode: 105,
			StatusMsg:  "用户不存在",
		}, nil
	}

	user := &types.User{
		Id:            info.User.Id,
		FollowCount:   *info.User.FollowCount,
		FollowerCount: *info.User.FollowerCount,
		IsFollow:      info.User.IsFollow,
		Name:          info.User.Name,
	}

	return &types.UserReply{
		StatusCode: 200,
		StatusMsg:  "success",
		User:       *user,
	}, err
}
