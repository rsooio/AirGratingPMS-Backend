package svc

import (
	"air-grating-pms-backend/rpc/authentication/internal/config"
	"fmt"

	"github.com/casbin/casbin/v2"
)

type ServiceContext struct {
	Config config.Config
	Casbin *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Print()
	e, err := casbin.NewEnforcer("etc/model.conf", "etc/policy.csv")
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		Casbin: e,
	}
}
