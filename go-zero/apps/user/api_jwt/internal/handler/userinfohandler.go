package handler

import (
	"net/http"

	"github.com/yaqia77/golang-gin-demo/apps/user/api_jwt/internal/logic"
	"github.com/yaqia77/golang-gin-demo/apps/user/api_jwt/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
