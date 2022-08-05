package svc

import (
	"air-grating-pms-backend/api/production/internal/config"
	opb "air-grating-pms-backend/rpc/order/pb"
	ppb "air-grating-pms-backend/rpc/production_plan/pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	OrderRpc          opb.OrderClient
	ProductionPlanRpc ppb.ProductionPlanClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		OrderRpc:          opb.NewOrderClient(zrpc.MustNewClient(c.OrderRpc).Conn()),
		ProductionPlanRpc: ppb.NewProductionPlanClient(zrpc.MustNewClient(c.ProductionPlanRpc).Conn()),
	}
}
