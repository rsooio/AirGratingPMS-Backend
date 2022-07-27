package svc

import (
	"air-grating-pms-backend/api/organization/internal/config"
	"air-grating-pms-backend/model/enterprise"
	"air-grating-pms-backend/model/staffer"
	"air-grating-pms-backend/model/workshop"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	DTM struct {
		Gid string
	}
	RPC struct {
		Target string
	}
)

type ServiceContext struct {
	Config          config.Config
	StafferModel    staffer.StafferModel
	WorkshopModel   workshop.WorkshopModel
	EnterpriseModel enterprise.EnterpriseModel
	DTM             DTM
	EnterpriseRPC   RPC
	StafferRPC      RPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	enterpriseTarget, err := c.EnterpriseRPC.BuildTarget()
	if err != nil {
		panic(err)
	}

	stafferTarget, err := c.StafferRPC.BuildTarget()
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:          c,
		StafferModel:    staffer.NewStafferModel(conn, c.CacheRedis),
		WorkshopModel:   workshop.NewWorkshopModel(conn, c.CacheRedis),
		EnterpriseModel: enterprise.NewEnterpriseModel(conn, c.CacheRedis),
		DTM: DTM{
			Gid: dtmgrpc.MustGenGid(c.DTM.Server),
		},
		EnterpriseRPC: RPC{
			Target: enterpriseTarget,
		},
		StafferRPC: RPC{
			Target: stafferTarget,
		},
	}
}
