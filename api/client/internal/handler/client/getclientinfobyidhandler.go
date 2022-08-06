package client

import (
	"net/http"

	"air-grating-pms-backend/api/client/internal/logic/client"
	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetClientInfoByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetClientInfoByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := client.NewGetClientInfoByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetClientInfoById(&req)
		response.Response(w, resp, err)
	}
}
