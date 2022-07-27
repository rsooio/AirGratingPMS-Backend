package operation

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OperationModel = (*customOperationModel)(nil)

type (
	// OperationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOperationModel.
	OperationModel interface {
		operationModel
	}

	customOperationModel struct {
		*defaultOperationModel
	}
)

// NewOperationModel returns a model for the database table.
func NewOperationModel(conn sqlx.SqlConn, c cache.CacheConf) OperationModel {
	return &customOperationModel{
		defaultOperationModel: newOperationModel(conn, c),
	}
}
