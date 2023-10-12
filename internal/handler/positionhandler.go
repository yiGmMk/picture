package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"picture/internal/logic"
	"picture/internal/svc"
	"picture/internal/types"
)

func PositionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PositionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPositionLogic(r.Context(), svcCtx)
		resp, err := l.Position(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
