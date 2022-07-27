package product_set

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductSetModel = (*customProductSetModel)(nil)

type (
	// ProductSetModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductSetModel.
	ProductSetModel interface {
		productSetModel
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
