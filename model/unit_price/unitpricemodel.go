package unit_price

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UnitPriceModel = (*customUnitPriceModel)(nil)

type (
	// UnitPriceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUnitPriceModel.
	UnitPriceModel interface {
		unitPriceModel
	}

	customUnitPriceModel struct {
		*defaultUnitPriceModel
	}
)

// NewUnitPriceModel returns a model for the database table.
func NewUnitPriceModel(conn sqlx.SqlConn, c cache.CacheConf) UnitPriceModel {
	return &customUnitPriceModel{
		defaultUnitPriceModel: newUnitPriceModel(conn, c),
	}
}
