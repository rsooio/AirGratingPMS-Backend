package product_set

import (
	"net/http"

	"air-grating-pms-backend/api/order/internal/logic/product_set"
	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductSetListByOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductSetListByOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := product_set.NewGetProductSetListByOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetProductSetListByOrder(&req)
		response.Response(w, resp, err)
	}
}
