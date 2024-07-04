package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/yaqia77/golang-gin-demo/apps/user/api_jwt/internal/config"
	"github.com/yaqia77/golang-gin-demo/apps/user/api_jwt/internal/handler"
	"github.com/yaqia77/golang-gin-demo/apps/user/api_jwt/internal/svc"
	"github.com/yaqia77/golang-gin-demo/pkg/response"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(JwtUnauthorizedResult))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println(err)
	httpx.WriteJson(w, http.StatusOK, response.Body{
		Code: 401,
		Msg:  "Unauthorized",
		Data: nil,
	})
}
