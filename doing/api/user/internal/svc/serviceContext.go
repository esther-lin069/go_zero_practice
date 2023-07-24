package svc

import (
	"go-zero/doing/api/user/internal/config"
	"go-zero/doing/rpc/userctl/userctl"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// userCtl rpc
	UserRpc userctl.UserCtl
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userctl.NewUserCtl(zrpc.MustNewClient(c.UserRpcConf)), // new UserCtl 帶入 rpc 連線資訊
	}
}
