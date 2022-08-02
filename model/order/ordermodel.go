package order

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

var (
	cacheOrderEnterpriseListCount = "cache:order:enterprise:list:count"
	orderRowsWithPrefixA          = "a." + strings.Join(orderFieldNames, ",a.")
)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		Insert(ctx context.Context, data *Order) (sql.Result, error)
		Delete(ctx context.Context, id int64) error
		FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (OrderList, int64, error)
		FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (OrderList, int64, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c),
	}
}

func (m *customOrderModel) Insert(ctx context.Context, data *Order) (sql.Result, error) {
	m.DelCacheCtx(ctx, cacheOrderEnterpriseListCount)
	return m.defaultOrderModel.Insert(ctx, data)
}

func (m *customOrderModel) Delete(ctx context.Context, id int64) error {
	m.DelCacheCtx(ctx, cacheOrderEnterpriseListCount)
	return m.defaultOrderModel.Delete(ctx, id)
}

func (m *customOrderModel) FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (OrderList, int64, error) {
	resp := make([]*Order, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", orderRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, offset, limit)
		if err != nil {
			logx.Errorf("orderEnterpriseId.findList error, enterpriseId=%d, offset=%d, limit=%d, err=%s", enterpriseId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ?", m.table)
		return m.QueryRow(&count, cacheOrderEnterpriseListCount, func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId)
		})
	})

	return resp, count, err
}

func (m *customOrderModel) FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (OrderList, int64, error) {
	resp := make([]*Order, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ? LIMIT ?, ?) AS b ON a.id = b.id", orderRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, workshopId, offset, limit)
		if err != nil {
			logx.Errorf("orderWorkshopId.findList error, enterpriseId=%d, workshopId=%d, offset=%d, limit=%d, err=%s", enterpriseId, workshopId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ?", m.table)
		return m.QueryRow(&count, cacheOrderEnterpriseListCount, func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId, workshopId)
		})
	})

	return resp, count, err
}
