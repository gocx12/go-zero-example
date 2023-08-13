// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"demo/service/user/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq  = user.LoginReq
	LoginResp = user.LoginResp

	UserService interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
