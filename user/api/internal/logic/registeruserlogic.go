package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"microservice/user/api/internal/svc"
	"microservice/user/api/internal/types"
	"microservice/user/common/xerr"
	"microservice/user/rpc/userclient"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterUserLogic) RegisterUser(req *types.RegisterOrLoginReq) (resp *types.DataReply, err error) {
	// todo: add your logic here and delete this line
	if len(req.Username) > 32 || len(req.Password) > 32 {
		return nil, nil
	}
	user, err := l.svcCtx.UserRpc.RegisterUser(l.ctx, &userclient.RegisterOrLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return &types.DataReply{
			StatusCode: 101,
			StatusMsg:  "用户名已注册",
		}, nil
	}

	now := time.Now().Unix()
	user.Token, err = l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(user.UserId))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(5000, "token生成失败"), "loginName: %s,err:%v", req, err)
	}

	return &types.DataReply{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     user.UserId,
		Token:      user.Token,
	}, nil
}

func (l *RegisterUserLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
