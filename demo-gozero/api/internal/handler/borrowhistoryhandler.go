package handler

import (
	"net/http"

	"bookstore/api/internal/logic"
	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BorrowHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BorrowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBorrowHistoryLogic(r.Context(), svcCtx)
		resp, err := l.BorrowHistory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
