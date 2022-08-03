package svc

import (
	"air-grating-pms-backend/api/technology/internal/config"
	"air-grating-pms-backend/rpc/technology/pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	TechnologyRpc    pb.TechnologyClient
	TechnologyTarget string
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(c.TechnologyRpc).Conn()

	return &ServiceContext{
		Config:           c,
		TechnologyRpc:    pb.NewTechnologyClient(conn),
		TechnologyTarget: conn.Target(),
	}
}
