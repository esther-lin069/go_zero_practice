package svc

import (
	"go_zero/demo/bookstore/internal/config"
	"go_zero/demo/bookstore/rpc/add/adder"
)

type ServiceContext struct {
	Config config.Config
	Adder  adder.Adder
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
