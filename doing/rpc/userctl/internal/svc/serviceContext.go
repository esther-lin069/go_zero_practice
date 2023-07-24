package svc

import (
	userModel "go-zero/doing/model/user"
	"go-zero/doing/rpc/userctl/internal/config"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// redis 連線物件
	RedisClient *redis.Client
	// db UserModel 物件
	UserModel userModel.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// DB 連線
	sqlConn := sqlx.NewMysql(c.DB.Dsn)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Cache[0].Host, // 拿快取的第一個 host 來用
			Password: "",
		}),
		UserModel: userModel.NewUserModel(sqlConn, c.Cache), // new model 帶入 sql 連線與快取設定
	}
}
