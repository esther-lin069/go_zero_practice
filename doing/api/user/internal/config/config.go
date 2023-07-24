package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	// app
	rest.RestConf

	// user rpc
	UserRpcConf zrpc.RpcClientConf
}
