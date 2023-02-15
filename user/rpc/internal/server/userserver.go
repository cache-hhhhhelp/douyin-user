// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"microservice/user/rpc/internal/logic"
	"microservice/user/rpc/internal/svc"
	"microservice/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) RegisterUser(ctx context.Context, in *user.RegisterOrLoginRequest) (*user.DataResponse, error) {
	l := logic.NewRegisterUserLogic(ctx, s.svcCtx)
	return l.RegisterUser(in)
}

func (s *UserServer) LoginUser(ctx context.Context, in *user.RegisterOrLoginRequest) (*user.DataResponse, error) {
	l := logic.NewLoginUserLogic(ctx, s.svcCtx)
	return l.LoginUser(in)
}

func (s *UserServer) GetUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}