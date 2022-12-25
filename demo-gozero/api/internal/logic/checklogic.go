package logic

import (
	"bookstore/rpc/check/checker"
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req *types.CheckReq) (resp *types.CheckResp, err error) {
	r, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{
		Book: req.Book,
	})
	if err != nil {
		logx.Error("查询数据错误: ", err)
		return &types.CheckResp{}, err
	}

	return &types.CheckResp{
		Found: r.Found,
		Price: r.Price,
	}, nil
}
