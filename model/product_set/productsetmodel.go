package product_set

import (
	"air-grating-pms-backend/utils/partial"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductSetModel = (*customProductSetModel)(nil)

var (
	cacheProductSetOrderListCount = "cache:product_set:order:list:count"
	productSetRowsWithPrefixA     = "a." + strings.Join(productSetFieldNames, ",a.")
)

type (
	// ProductSetModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductSetModel.
	ProductSetModel interface {
		productSetModel
		Insert(ctx context.Context, data *ProductSet) (sql.Result, error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, data *ProductSet) error
		FindListByOrderId(ctx context.Context, orderId int64, offset, limit int32) (ProductSetList, int64, error)
	}

	customProductSetModel struct {
		*defaultProductSetModel
	}
)

// NewProductSetModel returns a model for the database table.
func NewProductSetModel(conn sqlx.SqlConn, c cache.CacheConf) ProductSetModel {
	return &customProductSetModel{
		defaultProductSetModel: newProductSetModel(conn, c),
	}
}

func (m *customProductSetModel) Insert(ctx context.Context, data *ProductSet) (sql.Result, error) {
	m.DelCacheCtx(ctx, cacheProductSetOrderListCount)
	return m.defaultProductSetModel.Insert(ctx, data)
}

func (m *customProductSetModel) Delete(ctx context.Context, id int64) error {
	m.DelCacheCtx(ctx, cacheProductSetOrderListCount)
	return m.defaultProductSetModel.Delete(ctx, id)
}

func (m *customProductSetModel) Update(ctx context.Context, data *ProductSet) error {
	rows, args := partial.Partial(data)

	productSetIdKey := fmt.Sprintf("%s%v", cacheProductSetIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, partial.RowsWithPlaceHolder(rows))
		return conn.ExecCtx(ctx, query, append(args, data.Id)...)
	}, productSetIdKey)
	return err
}

func (m *customProductSetModel) FindListByOrderId(ctx context.Context, orderId int64, offset, limit int32) (ProductSetList, int64, error) {
	resp := make([]*ProductSet, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `order_id` = ? LIMIT ?, ?) AS b ON a.id = b.id", productSetRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, orderId, offset, limit)
		if err != nil {
			logx.Errorf("productSetOrderId.findList error, orderId=%d, offset=%d, limit=%d, err=%s", orderId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `order_id` = ?", m.table)
		return m.QueryRow(&count, cacheProductSetOrderListCount, func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, orderId)
		})
	})

	return resp, count, err
}
