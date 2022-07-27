package staffer

import (
	"net/http"

	"air-grating-pms-backend/common/response"

	"air-grating-pms-backend/api/staffer/internal/logic/staffer"
	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetStafferListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetStafferListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := staffer.NewGetStafferListLogic(r.Context(), svcCtx)
		resp, err := l.GetStafferList(&req)
		response.Response(w, resp, err)
	}
}
