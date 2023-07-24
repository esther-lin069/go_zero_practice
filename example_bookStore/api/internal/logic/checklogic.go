package logic

import (
	"context"

	"go_zero/demo/bookstore/api/internal/svc"
	"go_zero/demo/bookstore/api/internal/types"
	"go_zero/demo/bookstore/rpc/check/check"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req *types.CheckReq) (resp *types.CheckResp, err error) {
	r, err := l.svcCtx.Checker.Check(l.ctx, &check.CheckReq{
		Book: req.Book,
	})
	if err != nil {
		return nil, err
	}

	return &types.CheckResp{
		Found: r.Found,
		Price: r.Price,
	}, nil
}
