package handler

import (
	"net/http"

	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/logic"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/svc"
	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
