package product_set

import (
	"net/http"

	"air-grating-pms-backend/api/order/internal/logic/product_set"
	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteProductSetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProductSetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := product_set.NewDeleteProductSetLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProductSet(&req)
		response.Response(w, resp, err)
	}
}
