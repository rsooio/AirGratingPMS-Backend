package enterprise

import (
	"net/http"

	"air-grating-pms-backend/api/organization/internal/logic/enterprise"
	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/common/response"
)

func DeleteEnterpriseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := enterprise.NewDeleteEnterpriseLogic(r.Context(), svcCtx)
		resp, err := l.DeleteEnterprise()
		response.Response(w, resp, err)
	}
}
