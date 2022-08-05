package svc

import (
	"air-grating-pms-backend/model/production_plan"
	"air-grating-pms-backend/rpc/production_plan/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	ProductionPlanModel production_plan.ProductionPlanModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:              c,
		ProductionPlanModel: production_plan.NewProductionPlanModel(conn, c.CacheRedis),
	}
}
