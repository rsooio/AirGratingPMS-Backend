package svc

import (
	"air-grating-pms-backend/api/staffer/internal/config"
	"air-grating-pms-backend/rpc/staffer/pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	StafferRPC    pb.StafferClient
	StafferTarget string
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(c.StafferRPC).Conn()

	return &ServiceContext{
		Config:        c,
		StafferRPC:    pb.NewStafferClient(conn),
		StafferTarget: conn.Target(),
	}
}
