package logic

import (
	"context"
	"microservice/user/common/xerr"
	"microservice/user/rpc/models"

	"microservice/user/rpc/internal/svc"
	"microservice/user/rpc/types/user"

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

func (l *GetUserLogic) GetUser(in *user.UserRequest) (*user.UserResponse, error) {
	userNew := models.User{}
	res := l.svcCtx.DbEngine.First(&userNew, in.UserId)

	if res.Error != nil {
		return nil, xerr.NewErrCodeMsg(105, "用户不存在")
	}

	UserOV := &user.UserData{
		Id:            int64(userNew.ID),
		Name:          userNew.Username,
		FollowCount:   &userNew.FollowCount,
		FollowerCount: &userNew.FollowerCount,
	}
	// todo 查询是否关注
	//loginId := l.ctx.Value("userId")
	//l.svcCtx.DbEngine.Where()
	return &user.UserResponse{
		User: UserOV,
	}, nil
}
