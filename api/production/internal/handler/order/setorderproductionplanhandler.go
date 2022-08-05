package order

import (
	"net/http"

	"air-grating-pms-backend/api/production/internal/logic/order"
	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetOrderProductionPlanHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetOrderProductionPlanReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewSetOrderProductionPlanLogic(r.Context(), svcCtx)
		resp, err := l.SetOrderProductionPlan(&req)
		response.Response(w, resp, err)
	}
}
