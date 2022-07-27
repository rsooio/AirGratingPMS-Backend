package svc

import (
	"air-grating-pms-backend/api/staffer/internal/config"
	staffer "air-grating-pms-backend/model/staffer"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	StafferModel staffer.StafferModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:       c,
		StafferModel: staffer.NewStafferModel(conn, c.CacheRedis),
	}
}
