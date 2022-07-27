package workshop

import (
	"net/http"

	"air-grating-pms-backend/api/organization/internal/logic/workshop"
	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RetrieveWorkshopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RetrieveWorkshopReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := workshop.NewRetrieveWorkshopLogic(r.Context(), svcCtx)
		resp, err := l.RetrieveWorkshop(&req)
		response.Response(w, resp, err)
	}
}
