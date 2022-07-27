package enterprise

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EnterpriseModel = (*customEnterpriseModel)(nil)

type (
	// EnterpriseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEnterpriseModel.
	EnterpriseModel interface {
		enterpriseModel
	}

	customEnterpriseModel struct {
		*defaultEnterpriseModel
	}
)

// NewEnterpriseModel returns a model for the database table.
func NewEnterpriseModel(conn sqlx.SqlConn, c cache.CacheConf) EnterpriseModel {
	return &customEnterpriseModel{
		defaultEnterpriseModel: newEnterpriseModel(conn, c),
	}
}
