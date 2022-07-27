package statistic

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StatisticModel = (*customStatisticModel)(nil)

type (
	// StatisticModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStatisticModel.
	StatisticModel interface {
		statisticModel
	}

	customStatisticModel struct {
		*defaultStatisticModel
	}
)

// NewStatisticModel returns a model for the database table.
func NewStatisticModel(conn sqlx.SqlConn, c cache.CacheConf) StatisticModel {
	return &customStatisticModel{
		defaultStatisticModel: newStatisticModel(conn, c),
	}
}
