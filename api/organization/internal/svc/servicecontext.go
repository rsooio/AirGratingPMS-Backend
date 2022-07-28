package svc

import (
	"air-grating-pms-backend/api/organization/internal/config"
	epb "air-grating-pms-backend/rpc/enterprise/pb"
	spb "air-grating-pms-backend/rpc/staffer/pb"
	wpb "air-grating-pms-backend/rpc/workshop/pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	StafferRPC       spb.StafferClient
	WorkshopRPC      wpb.WorkshopClient
	EnterpriseRPC    epb.EnterpriseClient
	StafferTarget    string
	WorkshopTarget   string
	EnterpriseTarget string
}

func NewServiceContext(c config.Config) *ServiceContext {
	econn := zrpc.MustNewClient(c.EnterpriseRPC).Conn()
	wconn := zrpc.MustNewClient(c.WorkshopRPC).Conn()
	sconn := zrpc.MustNewClient(c.StafferRPC).Conn()

	return &ServiceContext{
		Config:           c,
		StafferRPC:       spb.NewStafferClient(sconn),
		WorkshopRPC:      wpb.NewWorkshopClient(wconn),
		EnterpriseRPC:    epb.NewEnterpriseClient(econn),
		StafferTarget:    sconn.Target(),
		EnterpriseTarget: econn.Target(),
		WorkshopTarget:   wconn.Target(),
	}
}
