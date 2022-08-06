package unit_price

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/mr"
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

func (m *customUnitPriceModel) SetUnitPriceByClientId(ctx context.Context, enterpriseId, workshopId, clientId int64, unitPricePairList []*UnitPricePair) error {
	return m.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		var fns []func() error
		for _, pair := range unitPricePairList {
			query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?) on duplicate key update `unit_price` = ? and `remark` = ?", m.table, unitPriceRowsExpectAutoSet)
			_, err := s.ExecCtx(ctx, query, enterpriseId, workshopId, pair.TechnologyId, clientId, pair.UnitPrice, pair.Remark, nil, pair.UnitPrice, pair.Remark)
			return err
		}
		return mr.Finish(fns...)
	})
}

func (m *customUnitPriceModel) GetUnitPriceByClientId(ctx context.Context, enterpriseId, workshopId, clientId int64)
