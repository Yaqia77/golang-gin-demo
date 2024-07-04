package handler

import (
	"net/http"

	"github.com/yaqia77/golang-gin-demo/apps/video/api/internal/logic"
	"github.com/yaqia77/golang-gin-demo/apps/video/api/internal/svc"
	"github.com/yaqia77/golang-gin-demo/apps/video/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetVideoLogic(r.Context(), svcCtx)
		resp, err := l.GetVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
