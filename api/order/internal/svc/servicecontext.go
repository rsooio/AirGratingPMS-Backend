package svc

import (
	"air-grating-pms-backend/api/order/internal/config"
	opb "air-grating-pms-backend/rpc/order/pb"
	ppb "air-grating-pms-backend/rpc/product/pb"
	spb "air-grating-pms-backend/rpc/product_set/pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	OrderRpc         opb.OrderClient
	ProductRpc       ppb.ProductClient
	ProductSetRpc    spb.ProductsetClient
	OrderTarget      string
	ProductTarget    string
	ProductSetTarget string
}

func NewServiceContext(c config.Config) *ServiceContext {
	oconn := zrpc.MustNewClient(c.OrderRpc).Conn()
	pconn := zrpc.MustNewClient(c.ProductRpc).Conn()
	sconn := zrpc.MustNewClient(c.ProductSetRpc).Conn()

	return &ServiceContext{
		Config:           c,
		OrderRpc:         opb.NewOrderClient(oconn),
		ProductRpc:       ppb.NewProductClient(pconn),
		ProductSetRpc:    spb.NewProductsetClient(sconn),
		OrderTarget:      oconn.Target(),
		ProductTarget:    pconn.Target(),
		ProductSetTarget: sconn.Target(),
	}
}
