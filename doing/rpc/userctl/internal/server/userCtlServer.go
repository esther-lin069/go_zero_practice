// Code generated by goctl. DO NOT EDIT.
// Source: userctl.proto

package server

import (
	"context"

	"go-zero/doing/rpc/userctl/internal/logic"
	"go-zero/doing/rpc/userctl/internal/svc"
	"go-zero/doing/rpc/userctl/pb"
)

type UserCtlServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserCtlServer
}

func NewUserCtlServer(svcCtx *svc.ServiceContext) *UserCtlServer {
	return &UserCtlServer{
		svcCtx: svcCtx,
	}
}

// Login rpc Logic
func (s *UserCtlServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// Register rpc Logic
func (s *UserCtlServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}
