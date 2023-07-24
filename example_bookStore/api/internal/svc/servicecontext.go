package svc

import (
	"go_zero/demo/bookstore/api/internal/config"
	"go_zero/demo/bookstore/rpc/add/adder"
	"go_zero/demo/bookstore/rpc/check/checker"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder
	Checker checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
