package svc

import (
	"air-grating-pms-backend/model/product_set"
	"air-grating-pms-backend/rpc/product_set/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	ProductSetModel product_set.ProductSetModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:          c,
		ProductSetModel: product_set.NewProductSetModel(conn, c.CacheRedis),
	}
}
