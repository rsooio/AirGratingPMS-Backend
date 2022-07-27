package production_plan

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductionPlanModel = (*customProductionPlanModel)(nil)

type (
	// ProductionPlanModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductionPlanModel.
	ProductionPlanModel interface {
		productionPlanModel
	}

	customProductionPlanModel struct {
		*defaultProductionPlanModel
	}
)

// NewProductionPlanModel returns a model for the database table.
func NewProductionPlanModel(conn sqlx.SqlConn, c cache.CacheConf) ProductionPlanModel {
	return &customProductionPlanModel{
		defaultProductionPlanModel: newProductionPlanModel(conn, c),
	}
}
