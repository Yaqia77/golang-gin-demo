package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Response(r *http.Request, w http.ResponseWriter, res any, err error) {
	if err != nil {
		boyd := Body{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		}

		httpx.WriteJson(w, http.StatusOK, boyd)
	}
	body := Body{
		Code: 200,
		Msg:  "success",
		Data: res,
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
