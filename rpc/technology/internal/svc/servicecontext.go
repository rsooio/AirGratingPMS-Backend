package svc

import (
	"air-grating-pms-backend/model/technology"
	"air-grating-pms-backend/rpc/technology/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	TechnologyModel technology.TechnologyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:          c,
		TechnologyModel: technology.NewTechnologyModel(conn, c.CacheRedis),
	}
}
