// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	technology "air-grating-pms-backend/api/technology/internal/handler/technology"
	"air-grating-pms-backend/api/technology/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/technology",
				Handler: technology.InsertTechnologyHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/technology",
				Handler: technology.DeleteTechnologyHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/technology",
				Handler: technology.UpdateTechnologyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/technology",
				Handler: technology.GetTechnologyInfoByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/workshop/technology",
				Handler: technology.GetTechnologyListByWorkshopHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/enterprise/technology",
				Handler: technology.GetTechnologyListByEnterpriseHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}