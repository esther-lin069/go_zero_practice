package logic

import (
	"context"
	"go-zero/doing/model/user"
	"go-zero/doing/rpc/userctl/internal/svc"
	"go-zero/doing/rpc/userctl/pb"
	"reflect"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/golang/mock/gomock"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestLoginLogic_Login(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *pb.LoginReq
	}

	ctx := context.TODO()

	// 宣告 redis mock 物件（實際上本 func 中無作用，僅做範例參考用）
	redisMockClient, redisMock := redismock.NewClientMock()

	// 宣告 mysql mock 物件
	dbMockCtl := gomock.NewController(t)
	defer dbMockCtl.Finish()
	dbMock := user.NewMockuserModel(dbMockCtl)

	mockLogic := fields{
		ctx: ctx,
		svcCtx: &svc.ServiceContext{
			RedisClient: redisMockClient,
			UserModel:   dbMock,
		},
		Logger: logx.WithContext(ctx),
	}

	// 設定 redis mock 值（實際上本 func 中無作用，僅做範例參考用）
	redisMock.ExpectGetSet("cache:user:id:1", `{"Nickname":"testUser01"}`)

	// 設定 mysql mock 值
	dbMock.
		EXPECT().
		FindOneByName(ctx, "testUser").
		Return(&user.User{
			Password: "qwe123",
			Nickname: "testUser01",
		}, nil).
		Times(2)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.LoginResp
		wantErr bool
	}{
		{
			name:   "successful login",
			fields: mockLogic,
			args: args{
				in: &pb.LoginReq{
					Username: "testUser",
					Password: "qwe123",
				},
			},
			want: &pb.LoginResp{
				Nickname: "testUser01",
			},
			wantErr: false,
		},
		{
			name:   "wrong password",
			fields: mockLogic,
			args: args{
				in: &pb.LoginReq{
					Username: "testUser",
					Password: "qwe1234",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LoginLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.Login(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginLogic.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginLogic.Login() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
