// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"github.com/cache-hhhhhelp/douyin-user/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DataResp           = user.DataResp
	ManyUserReq        = user.ManyUserReq
	ManyUserResp       = user.ManyUserResp
	RegisterOrLoginReq = user.RegisterOrLoginReq
	UserReq            = user.UserReq
	UserResp           = user.UserResp

	User interface {
		RegisterUser(ctx context.Context, in *RegisterOrLoginReq, opts ...grpc.CallOption) (*DataResp, error)
		LoginUser(ctx context.Context, in *RegisterOrLoginReq, opts ...grpc.CallOption) (*DataResp, error)
		GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error)
		GetManyUser(ctx context.Context, in *ManyUserReq, opts ...grpc.CallOption) (*ManyUserResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) RegisterUser(ctx context.Context, in *RegisterOrLoginReq, opts ...grpc.CallOption) (*DataResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RegisterUser(ctx, in, opts...)
}

func (m *defaultUser) LoginUser(ctx context.Context, in *RegisterOrLoginReq, opts ...grpc.CallOption) (*DataResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.LoginUser(ctx, in, opts...)
}

func (m *defaultUser) GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) GetManyUser(ctx context.Context, in *ManyUserReq, opts ...grpc.CallOption) (*ManyUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetManyUser(ctx, in, opts...)
}
