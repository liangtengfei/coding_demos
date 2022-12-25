package handler

import (
	"net/http"

	"bookstore/api/internal/logic"
	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BorrowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BorrowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBorrowLogic(r.Context(), svcCtx)
		resp, err := l.Borrow(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
