package product_set

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"air-grating-pms-backend/utils/partial"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductSetModel = (*customProductSetModel)(nil)

var (
	cacheProductSetOrderListCountPrefix = "cache:product_set:order:list:count:"
	productSetRowsWithPrefixA           = "a." + strings.Join(productSetFieldNames, ",a.")
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

func cacheProductSetOrderListKey(orderId int64) string {
	return fmt.Sprintf("%s%v", cacheProductSetOrderListCountPrefix, orderId)
}

func (m *customProductSetModel) Insert(ctx context.Context, data *ProductSet) (sql.Result, error) {
	productSetIdKey := fmt.Sprintf("%s%v", cacheProductSetIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, productSetRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.OrderId, data.Remark, data.Version)
	}, productSetIdKey, cacheProductSetOrderListKey(data.OrderId))
	return ret, err
}

func (m *customProductSetModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	productSetIdKey := fmt.Sprintf("%s%v", cacheProductSetIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, productSetIdKey, cacheProductSetOrderListKey(data.OrderId))
	return err
}

func (m *customProductSetModel) Update(ctx context.Context, data *ProductSet) error {
	rows, args := partial.Partial(data)

	productSetIdKey := fmt.Sprintf("%s%v", cacheProductSetIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rows.StringWithPlaceHolder())
		return conn.ExecCtx(ctx, query, *args.WithId(data.Id)...)
	}, productSetIdKey)
	return err
}

func (m *customProductSetModel) FindListByOrderId(ctx context.Context, orderId int64, offset, limit int32) (ProductSetList, int64, error) {
	resp := make(ProductSetList, 0)
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
		return m.QueryRow(&count, cacheProductSetOrderListKey(orderId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, orderId)
		})
	})

	return resp, count, err
}
