package production_plan

import (
	"net/http"

	"air-grating-pms-backend/api/production/internal/logic/production_plan"
	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CompleteProductionPlanHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CompleteProductionPlanReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := production_plan.NewCompleteProductionPlanLogic(r.Context(), svcCtx)
		resp, err := l.CompleteProductionPlan(&req)
		response.Response(w, resp, err)
	}
}
