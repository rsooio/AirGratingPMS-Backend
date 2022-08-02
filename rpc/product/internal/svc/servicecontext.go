package svc

import (
	"air-grating-pms-backend/model/product"
	"air-grating-pms-backend/rpc/product/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel product.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:       c,
		ProductModel: product.NewProductModel(conn, c.CacheRedis),
	}
}
