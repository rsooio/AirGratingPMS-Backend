package technology

import (
	"net/http"

	"air-grating-pms-backend/api/technology/internal/logic/technology"
	"air-grating-pms-backend/api/technology/internal/svc"
	"air-grating-pms-backend/api/technology/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTechnologyListByWorkshopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTechnologyListByWorkshopReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := technology.NewGetTechnologyListByWorkshopLogic(r.Context(), svcCtx)
		resp, err := l.GetTechnologyListByWorkshop(&req)
		response.Response(w, resp, err)
	}
}
