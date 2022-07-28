package enterprise

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EnterpriseModel = (*customEnterpriseModel)(nil)

type (
	// EnterpriseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEnterpriseModel.
	EnterpriseModel interface {
		enterpriseModel
		SoftDelete(ctx context.Context, id int64) error
	}

	customEnterpriseModel struct {
		*defaultEnterpriseModel
		Deleted DeletedEnterpriseModel
	}
)

// NewEnterpriseModel returns a model for the database table.
func NewEnterpriseModel(conn sqlx.SqlConn, c cache.CacheConf) EnterpriseModel {
	return &customEnterpriseModel{
		defaultEnterpriseModel: newEnterpriseModel(conn, c),
		Deleted:                newDeletedEnterpriseModel(conn, c),
	}
}

func (m *customEnterpriseModel) SoftDelete(ctx context.Context, id int64) error {
	return m.Transact(func(s sqlx.Session) error {
		retrive := fmt.Sprintf("select %s from %s where `id` = ? limit 1", enterpriseRows, m.table)
		insert := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.Deleted.tableName(), deletedEnterpriseRowsExpectAutoSet)
		delete := fmt.Sprintf("delete from %s where `id` = ?", m.table)

		var resp Enterprise
		err := s.QueryRowCtx(ctx, &resp, retrive, id)
		if err != nil {
			return err
		}

		_, err = s.ExecCtx(ctx, insert, resp.Id, resp.Name, resp.Address, resp.CreateTime, resp.Remark, resp.Version)
		if err != nil {
			return err
		}

		_, err = s.ExecCtx(ctx, delete, resp.Id)
		return err
	})
}
