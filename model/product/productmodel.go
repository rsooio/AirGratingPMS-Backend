package product

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

var _ ProductModel = (*customProductModel)(nil)

var (
	cacheProductProductSetListCountPrefix = "cache:product:product_set:list:count:"
	productRowsWithPrefixA                = "a." + strings.Join(productFieldNames, ",a.")
)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		Insert(ctx context.Context, data *Product) (sql.Result, error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, data *Product) error
		FindListByProductSetId(ctx context.Context, productSetId int64, offset, limit int32) (ProductList, int64, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c),
	}
}

func cacheProductProductSetListCountKey(productSetId int64) string {
	return fmt.Sprintf("%s%v", cacheProductProductSetListCountPrefix, productSetId)
}

func (m *customProductModel) Insert(ctx context.Context, data *Product) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, productRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProductSetId, data.TechnologyId, data.Length, data.Width, data.UnitPrice, data.Quantity, data.Remark, data.Version)
	}, productIdKey, cacheProductProductSetListCountKey(data.ProductSetId))
	return ret, err
}

func (m *customProductModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, productIdKey, cacheProductProductSetListCountKey(data.ProductSetId))
	return err
}

func (m *customProductModel) Update(ctx context.Context, data *Product) error {
	rows, args := partial.Partial(data)

	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rows.StringWithPlaceHolder())
		return conn.ExecCtx(ctx, query, *args.WithId(data.Id)...)
	}, productIdKey)
	return err
}

func (m *customProductModel) FindListByProductSetId(ctx context.Context, productSetId int64, offset, limit int32) (ProductList, int64, error) {
	resp := make(ProductList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `product_set_id` = ? LIMIT ?, ?) AS b ON a.id = b.id", productRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, productSetId, offset, limit)
		if err != nil {
			logx.Errorf("productProductSetId.findList error, productSetId=%d, offset=%d, limit=%d, err=%s", productSetId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `product_set_id` = ?", m.table)
		return m.QueryRow(&count, cacheProductProductSetListCountKey(productSetId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, productSetId)
		})
	})

	return resp, count, err
}
