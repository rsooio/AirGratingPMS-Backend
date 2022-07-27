package svc

import (
	"air-grating-pms-backend/model/workshop"
	"air-grating-pms-backend/rpc/workshop/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	WorkshopModel workshop.WorkshopModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:        c,
		WorkshopModel: workshop.NewWorkshopModel(conn, c.CacheRedis),
	}
}
