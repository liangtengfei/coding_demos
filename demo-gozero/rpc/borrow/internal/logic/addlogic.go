package logic

import (
	"bookstore/rpc/model/record"
	"context"
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"bookstore/rpc/borrow/borrow"
	"bookstore/rpc/borrow/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *borrow.BorrowReq) (*borrow.BorrowResp, error) {
	book, err := l.svcCtx.BookModel.FindOne(l.ctx, in.Book)
	if err != nil {
		logx.Error("查找书籍：", err)
		return nil, err
	}

	data := &record.Record{
		Book: book.Book,
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{},
	}
	_, err = l.svcCtx.RecordModel.Insert(l.ctx, data)
	if err != nil {
		logx.Error("插入记录：", err)
		return nil, err
	}

	return &borrow.BorrowResp{
		Id:        data.Id,
		Book:      in.Book,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: nil,
	}, nil
}
