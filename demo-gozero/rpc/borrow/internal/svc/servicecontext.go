package svc

import (
	"bookstore/rpc/borrow/internal/config"
	"bookstore/rpc/model/book"
	"bookstore/rpc/model/record"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RecordModel record.RecordModel
	BookModel   book.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RecordModel: record.NewRecordModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
		BookModel:   book.NewBookModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
	}
}
