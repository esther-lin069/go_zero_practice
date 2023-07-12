package logic

import (
	"context"
	"database/sql"
	"fmt"

	"go_zero/demo/demo/internal/svc"
	"go_zero/demo/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	result, err := l.svcCtx.UserModel.FindOneByName(l.ctx, sql.NullString{
		String: req.Name,
		Valid:  true,
	})

	if result.Password != req.Password {
		return
	}

	// 存入redis
	err = l.svcCtx.RedisClient.Set(fmt.Sprintf("%d", result.Id), result.Name.String)
	resp = &types.LoginResponse{
		Token: fmt.Sprintf("%s@%d", result.Name.String, result.Id),
	}

	return
}
