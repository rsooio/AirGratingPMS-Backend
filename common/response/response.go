package response

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		if len(err.Error()) >= 4 && err.Error()[0] == '#' && err.Error()[3] == '#' {
			code, err := strconv.Atoi(err.Error()[1:3])
			if err != nil {
				body.Code = -1
				body.Msg = err.Error()
			} else {
				body.Code = code
				body.Msg = err.Error()[4:]
			}
		} else {
			body.Code = -1
			body.Msg = err.Error()
		}
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
