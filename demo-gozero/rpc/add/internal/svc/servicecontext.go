package svc

import (
	"bookstore/rpc/add/internal/config"
	"bookstore/rpc/model/book"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  book.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  book.NewBookModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
	}
}
