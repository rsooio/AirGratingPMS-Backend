package product

import (
	"net/http"

	"air-grating-pms-backend/api/order/internal/logic/product"
	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InsertProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := product.NewInsertProductLogic(r.Context(), svcCtx)
		resp, err := l.InsertProduct(&req)
		response.Response(w, resp, err)
	}
}
