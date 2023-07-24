package logic

import (
	"context"

	userModel "go-zero/doing/model/user"
	"go-zero/doing/rpc/userctl/internal/svc"
	"go-zero/doing/rpc/userctl/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register rpc Logic
func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 於資料庫中新增此筆資料
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &userModel.User{
		Name:     in.Username,
		Password: in.Password,
		Gender:   in.Gender,
		Mobile:   in.Mobile,
		Nickname: in.Nickname,
	})

	if err != nil {
		return nil, err
	}

	return &pb.RegisterResp{
		Result: "ok",
	}, nil
}
