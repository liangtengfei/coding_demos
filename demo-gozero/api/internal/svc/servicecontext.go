package svc

import (
	"bookstore/api/internal/config"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/borrow/borrower"
	"bookstore/rpc/check/checker"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Adder    adder.Adder
	Checker  checker.Checker
	Borrower borrower.Borrower
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Adder:    adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker:  checker.NewChecker(zrpc.MustNewClient(c.Check)),
		Borrower: borrower.NewBorrower(zrpc.MustNewClient(c.Borrow)),
	}
}
