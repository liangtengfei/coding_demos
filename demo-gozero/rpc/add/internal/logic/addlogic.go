package logic

import (
	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/svc"
	"bookstore/rpc/model/book"
	"context"

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

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	_, err := l.svcCtx.Model.Insert(l.ctx, &book.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}

	return &add.AddResp{Ok: true}, nil
}
