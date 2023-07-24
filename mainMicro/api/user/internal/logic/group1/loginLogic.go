package group1

import (
	"context"
	"errors"

	"go-zero/doing/api/user/internal/svc"
	"go-zero/doing/api/user/internal/types"
	"go-zero/doing/rpc/userctl/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if req.Username == "test" {
		return nil, errors.New("wrong username")
	}

	r, err := l.svcCtx.UserRpc.Login(l.ctx, &pb.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken: r.Nickname,
	}, nil
}
