package enterprise

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeletedEnterpriseModel = (*customDeletedEnterpriseModel)(nil)

type (
	// DeletedEnterpriseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeletedEnterpriseModel.
	DeletedEnterpriseModel interface {
		deletedEnterpriseModel
		tableName() string
	}

	customDeletedEnterpriseModel struct {
		*defaultDeletedEnterpriseModel
	}
)

// NewDeletedEnterpriseModel returns a model for the database table.
func NewDeletedEnterpriseModel(conn sqlx.SqlConn, c cache.CacheConf) DeletedEnterpriseModel {
	return &customDeletedEnterpriseModel{
		defaultDeletedEnterpriseModel: newDeletedEnterpriseModel(conn, c),
	}
}

func (m *customDeletedEnterpriseModel) tableName() string {
	return m.table
}
