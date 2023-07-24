package logic

import (
	"context"
	"database/sql"
	"fmt"

	"go_zero/demo/demo/internal/svc"
	"go_zero/demo/demo/internal/types"
	usermodel "go_zero/demo/demo/model/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignUpLogic {
	return &UserSignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignUpLogic) UserSignUp(req *types.SignUpRequest) (resp *types.SignUpResponse, err error) {
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &usermodel.User{
		Name: sql.NullString{
			String: req.Name,
			Valid:  true,
		},
		Password: req.Password,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Gender:   req.Gender,
	})

	resp = &types.SignUpResponse{
		Message: fmt.Sprintf("Name: %s Added", req.Name),
	}

	return
}
