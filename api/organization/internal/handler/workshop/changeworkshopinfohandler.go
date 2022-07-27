package workshop

import (
	"net/http"

	"air-grating-pms-backend/api/organization/internal/logic/workshop"
	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChangeWorkshopInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangeWorkshopInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := workshop.NewChangeWorkshopInfoLogic(r.Context(), svcCtx)
		resp, err := l.ChangeWorkshopInfo(&req)
		response.Response(w, resp, err)
	}
}
