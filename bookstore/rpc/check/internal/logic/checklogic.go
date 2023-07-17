package logic

import (
	"context"
	"errors"

	"go_zero/demo/bookstore/rpc/check/check"
	"go_zero/demo/bookstore/rpc/check/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	if in.Book == "NaN" {
		return nil, errors.New("not found")
	}

	return &check.CheckResp{}, nil
}
