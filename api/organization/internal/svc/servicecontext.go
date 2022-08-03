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
	StafferRpc       spb.StafferClient
	WorkshopRpc      wpb.WorkshopClient
	EnterpriseRpc    epb.EnterpriseClient
	StafferTarget    string
	WorkshopTarget   string
	EnterpriseTarget string
}

func NewServiceContext(c config.Config) *ServiceContext {
	econn := zrpc.MustNewClient(c.EnterpriseRpc).Conn()
	wconn := zrpc.MustNewClient(c.WorkshopRpc).Conn()
	sconn := zrpc.MustNewClient(c.StafferRpc).Conn()

	return &ServiceContext{
		Config:           c,
		StafferRpc:       spb.NewStafferClient(sconn),
		WorkshopRpc:      wpb.NewWorkshopClient(wconn),
		EnterpriseRpc:    epb.NewEnterpriseClient(econn),
		StafferTarget:    sconn.Target(),
		EnterpriseTarget: econn.Target(),
		WorkshopTarget:   wconn.Target(),
	}
}
