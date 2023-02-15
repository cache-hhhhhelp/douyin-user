package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"microservice/user/common/xerr"
	"microservice/user/rpc/userclient"
	"time"

	"microservice/user/api/internal/svc"
	"microservice/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginUserLogic) LoginUser(req *types.RegisterOrLoginReq) (resp *types.DataReply, err error) {
	login, err := l.svcCtx.UserRpc.LoginUser(l.ctx, &userclient.RegisterOrLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return &types.DataReply{
			StatusCode: 102,
			StatusMsg:  "用户不存在或密码不正确",
		}, nil
	}

	now := time.Now().Unix()
	login.Token, err = l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, login.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(5000, "token生成失败"), "loginName: %s,err:%v", req, err)
	}

	return &types.DataReply{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     login.UserId,
		Token:      login.Token,
	}, err
}

func (l *LoginUserLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
