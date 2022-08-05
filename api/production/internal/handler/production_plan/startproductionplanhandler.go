package production_plan

import (
	"net/http"

	"air-grating-pms-backend/api/production/internal/logic/production_plan"
	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StartProductionPlanHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StartProductionPlanReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := production_plan.NewStartProductionPlanLogic(r.Context(), svcCtx)
		resp, err := l.StartProductionPlan(&req)
		response.Response(w, resp, err)
	}
}
