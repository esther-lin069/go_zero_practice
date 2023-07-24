package group1

import (
	"context"
	"go-zero/doing/api/user/internal/svc"
	"go-zero/doing/api/user/internal/types"
	"go-zero/doing/rpc/userctl/pb"
	"go-zero/doing/rpc/userctl/userctl"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestLoginLogic_Login(t *testing.T) {
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		req *types.LoginReq
	}

	ctx := context.TODO()

	// 宣告 mock 物件
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRPC := userctl.NewMockUserCtl(ctrl)

	mockLogicFields := fields{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: &svc.ServiceContext{
			UserRpc: mockRPC,
		},
	}

	// 定義 mock 值
	mockRPC.
		EXPECT().
		Login(ctx, &pb.LoginReq{
			Username: "testUser",
			Password: "qwe123",
		}).
		Return(&pb.LoginResp{
			Nickname: "testUser01",
		}, nil)

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.LoginResp
		wantErr  bool
	}{
		{
			name:   "successful login",
			fields: mockLogicFields,
			args: args{
				req: &types.LoginReq{
					Username: "testUser",
					Password: "qwe123",
				},
			},
			wantResp: &types.LoginResp{
				AccessToken: "testUser01",
			},
			wantErr: false,
		},
		{
			name:   "wrong username",
			fields: mockLogicFields,
			args: args{
				req: &types.LoginReq{
					Username: "test",
					Password: "qwe123",
				},
			},
			wantResp: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LoginLogic{
				Logger: tt.fields.Logger,
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
			}
			gotResp, err := l.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginLogic.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("LoginLogic.Login() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
