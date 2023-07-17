package handler

import (
	"net/http"

	"go_zero/demo/bookstore/api/internal/logic"
	"go_zero/demo/bookstore/api/internal/svc"
	"go_zero/demo/bookstore/api/internal/types"
	pkg "go_zero/demo/pkg"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCheckLogic(r.Context(), svcCtx)
		resp, err := l.Check(&req)
		if err != nil {
			pkg.Response(w, nil, err)
		} else {
			pkg.Response(w, resp, nil)
		}
	}
}
