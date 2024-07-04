// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/yaqia77/golang-gin-demo/rpc__demo/user_api_rpc/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: userCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: userInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/users"),
	)
}
