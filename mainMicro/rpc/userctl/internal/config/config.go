package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 資料庫連線
	DB struct {
		Dsn string
	}

	// 快取設定
	Cache cache.CacheConf
}
