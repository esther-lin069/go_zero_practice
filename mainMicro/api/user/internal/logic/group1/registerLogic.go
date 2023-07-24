package group1

import (
	"context"

	"go-zero/doing/api/user/internal/svc"
	"go-zero/doing/api/user/internal/types"
	"go-zero/doing/rpc/userctl/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// call rpc func
	r, err := l.svcCtx.UserRpc.Register(l.ctx, &pb.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Gender:   req.Gender,
		Nickname: req.Nickname,
	})

	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Result: r.Result,
	}, nil
}
