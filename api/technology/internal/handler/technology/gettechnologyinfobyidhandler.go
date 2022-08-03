package technology

import (
	"net/http"

	"air-grating-pms-backend/api/technology/internal/logic/technology"
	"air-grating-pms-backend/api/technology/internal/svc"
	"air-grating-pms-backend/api/technology/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTechnologyInfoByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTechnologyInfoByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := technology.NewGetTechnologyInfoByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetTechnologyInfoById(&req)
		response.Response(w, resp, err)
	}
}
