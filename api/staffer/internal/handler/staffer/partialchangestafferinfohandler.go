package staffer

import (
	"net/http"

	"air-grating-pms-backend/common/response"

	"air-grating-pms-backend/api/staffer/internal/logic/staffer"
	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PartialChangeStafferInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PartialChangeStafferInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := staffer.NewPartialChangeStafferInfoLogic(r.Context(), svcCtx)
		resp, err := l.PartialChangeStafferInfo(&req)
		response.Response(w, resp, err)
	}
}
