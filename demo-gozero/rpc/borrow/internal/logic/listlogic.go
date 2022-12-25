package logic

import (
	"bookstore/rpc/borrow/borrow"
	"bookstore/rpc/borrow/internal/svc"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *borrow.BorrowReq, stream borrow.Borrower_ListServer) error {
	resp, err := l.svcCtx.RecordModel.Find(l.ctx, in.Book)
	if err != nil {
		logx.Error("查询借阅记录：", err)
		return err
	}
	logx.Info("借阅记录：", len(*resp))
	for _, r := range *resp {
		err = stream.SendMsg(&borrow.BorrowResp{
			Id:        r.Id,
			Book:      r.Book,
			CreatedAt: timestamppb.New(r.CreatedAt.Time),
			UpdatedAt: timestamppb.New(r.UpdatedAt.Time),
		})
		if err != nil {
			logx.Error("发送记录错误")
			return err
		}
	}
	return err
}
