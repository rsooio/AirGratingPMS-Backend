package svc

import (
	"air-grating-pms-backend/api/client/internal/config"
	"air-grating-pms-backend/rpc/client/pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	ClientRpc pb.ClientClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		ClientRpc: pb.NewClientClient(zrpc.MustNewClient(c.ClientRpc).Conn()),
	}
}
