package client

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClientModel = (*customClientModel)(nil)

type (
	// ClientModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClientModel.
	ClientModel interface {
		clientModel
	}

	customClientModel struct {
		*defaultClientModel
	}
)

// NewClientModel returns a model for the database table.
func NewClientModel(conn sqlx.SqlConn, c cache.CacheConf) ClientModel {
	return &customClientModel{
		defaultClientModel: newClientModel(conn, c),
	}
}
