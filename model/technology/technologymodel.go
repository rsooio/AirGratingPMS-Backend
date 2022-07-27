package technology

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TechnologyModel = (*customTechnologyModel)(nil)

type (
	// TechnologyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTechnologyModel.
	TechnologyModel interface {
		technologyModel
	}

	customTechnologyModel struct {
		*defaultTechnologyModel
	}
)

// NewTechnologyModel returns a model for the database table.
func NewTechnologyModel(conn sqlx.SqlConn, c cache.CacheConf) TechnologyModel {
	return &customTechnologyModel{
		defaultTechnologyModel: newTechnologyModel(conn, c),
	}
}
