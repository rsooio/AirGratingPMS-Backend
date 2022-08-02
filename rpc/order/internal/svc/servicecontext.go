package svc

import (
	"air-grating-pms-backend/model/order"
	"air-grating-pms-backend/rpc/order/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel order.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		OrderModel: order.NewOrderModel(conn, c.CacheRedis),
	}
}
