package order

import (
	"net/http"

	"air-grating-pms-backend/api/order/internal/logic/order"
	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateOrderInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateOrderInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewUpdateOrderInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateOrderInfo(&req)
		response.Response(w, resp, err)
	}
}
