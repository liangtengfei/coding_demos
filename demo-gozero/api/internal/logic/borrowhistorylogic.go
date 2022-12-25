package logic

import (
	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/rpc/borrow/borrow"
	"context"
	"io"

	"github.com/zeromicro/go-zero/core/logx"
)

type BorrowHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBorrowHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BorrowHistoryLogic {
	return &BorrowHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BorrowHistoryLogic) BorrowHistory(req *types.BorrowReq) (*types.HistoryResp, error) {
	stream, err := l.svcCtx.Borrower.List(l.ctx, &borrow.BorrowReq{Book: req.Book})
	if err != nil {
		logx.Error("获取记录错误")
		return nil, err
	}
	resp := &types.HistoryResp{
		List: []types.Record{},
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logx.Error("接收记录错误")
			return resp, err
		}
		resp.List = append(resp.List, types.Record{
			Id:        res.Id,
			Book:      res.Book,
			CreatedAt: res.CreatedAt.String(),
		})
	}

	return resp, nil
	//logx.Info("请求成功")
	//for {
	//	res, err := r.Recv()
	//	if err == io.EOF {
	//		logx.Error("没有更多数据")
	//		return resp, nil
	//	}
	//	if err != nil {
	//		logx.Error("接收记录错误")
	//		return resp, err
	//	}
	//	logx.Infof("接收到记录：%v", res)
	//	resp.List = append(resp.List, types.Record{
	//		Id:        res.Id,
	//		Book:      res.Book,
	//		CreatedAt: res.CreatedAt.String(),
	//	})
	//}
}
