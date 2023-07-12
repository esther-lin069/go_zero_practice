package svc

import (
	"go_zero/demo/demo/internal/config"
	usermodel "go_zero/demo/demo/model/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 定義 redis 連線物件
	RedisClient *redis.Redis

	// 定義 UserModel 結構體
	UserModel usermodel.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 資料庫連線
	sqlConn := sqlx.NewMysql(c.DB.DsnString)

	// redis 連線
	redisConn := redis.MustNewRedis(c.Redis)

	return &ServiceContext{
		Config: c,
		// redis 連線物件
		RedisClient: redisConn,
		// UserModel db 連線物件
		UserModel: usermodel.NewUserModel(sqlConn),
	}
}
