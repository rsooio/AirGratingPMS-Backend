package svc

import (
	"air-grating-pms-backend/model/enterprise"
	"air-grating-pms-backend/rpc/enterprise/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	EnterpriseModel enterprise.EnterpriseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:          c,
		EnterpriseModel: enterprise.NewEnterpriseModel(conn, c.CacheRedis),
	}
}
