package response

import (
	"air-grating-pms-backend/utils"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error, ctxs ...context.Context) {
	var body Body
	if err != nil {
		body.Code = utils.SelectErrorCode(ctxs...)
		body.Msg = err.Error()
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
