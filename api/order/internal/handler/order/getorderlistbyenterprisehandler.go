package order

import (
	"net/http"

	"air-grating-pms-backend/api/order/internal/logic/order"
	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOrderListByEnterpriseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrderListByEnterpriseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewGetOrderListByEnterpriseLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderListByEnterprise(&req)
		response.Response(w, resp, err)
	}
}
