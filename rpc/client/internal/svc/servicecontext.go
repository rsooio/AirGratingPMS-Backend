package svc

import (
	"air-grating-pms-backend/model/client"
	"air-grating-pms-backend/rpc/client/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	ClientModel client.ClientModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:      c,
		ClientModel: client.NewClientModel(conn, c.CacheRedis),
	}
}
