package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"picture/internal/logic"
	"picture/internal/svc"
	"picture/internal/types"
)

func ReceiveImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewReceiveImgLogic(r.Context(), svcCtx)
		resp, err := l.ReceiveImg(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
