package logic

import (
	"context"
	"errors"

	"go-zero/doing/rpc/userctl/internal/svc"
	"go-zero/doing/rpc/userctl/pb"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login rpc Logic
func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 取出資料
	r, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}

	// 確認帳號密碼吻合
	if r.Password != in.Password {
		return nil, errors.New("wrong username or password")
	}

	// 回傳
	return &pb.LoginResp{
		Nickname: r.Nickname, // 暫時回傳暱稱，可用來驗證緩存
	}, nil
}

// generate uuid
func getUUIDToken() string {
	return uuid.NewString()
}
