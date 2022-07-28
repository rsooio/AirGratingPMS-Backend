package staffer

import (
	"net/http"

	"air-grating-pms-backend/api/staffer/internal/logic/staffer"
	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetStafferListByWorkshopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetStafferListByWorkshopReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := staffer.NewGetStafferListByWorkshopLogic(r.Context(), svcCtx)
		resp, err := l.GetStafferListByWorkshop(&req)
		response.Response(w, resp, err)
	}
}
