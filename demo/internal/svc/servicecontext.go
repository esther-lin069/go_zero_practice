package svc

import (
	"go_zero/demo/demo/internal/config"
	"go_zero/demo/demo/model/mysql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	// 定義 UserModel 結構體
	UserModel mysql.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 資料庫連線
	sqlConn := sqlx.NewMysql(c.DB.DsnString)

	return &ServiceContext{
		Config: c,
		// new 一個UserModel物件
		UserModel: mysql.NewUserModel(sqlConn),
	}
}
