package enterprise

import (
	"net/http"

	"air-grating-pms-backend/api/organization/internal/logic/enterprise"
	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PartialChangeEnterpriseInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PartialChangeEnterpriseInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := enterprise.NewPartialChangeEnterpriseInfoLogic(r.Context(), svcCtx)
		resp, err := l.PartialChangeEnterpriseInfo(&req)
		response.Response(w, resp, err)
	}
}
