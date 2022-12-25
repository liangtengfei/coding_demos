package logic

import (
	"bookstore/rpc/borrow/borrower"
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BorrowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BorrowLogic {
	return &BorrowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BorrowLogic) Borrow(req *types.BorrowReq) (resp *types.BorrowResp, err error) {
	if req == nil {
		req = &types.BorrowReq{Book: "西游记"}
	}
	r, err := l.svcCtx.Borrower.Add(l.ctx, &borrower.BorrowReq{
		Book: req.Book,
	})
	if err != nil {
		logx.Error("新增借阅错误: ", err)
		return nil, err
	}

	return &types.BorrowResp{
		Ok:       true,
		BorrowId: r.Id,
	}, nil
}
