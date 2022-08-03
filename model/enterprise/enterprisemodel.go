package enterprise

import (
	"air-grating-pms-backend/utils/partial"
	"context"
	"database/sql"
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
		Update(ctx context.Context, data *Enterprise) error
		Delete(ctx context.Context, id int64) error
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

func (m *customEnterpriseModel) Update(ctx context.Context, data *Enterprise) error {
	rows, args := partial.Partial(data)

	enterpriseIdKey := fmt.Sprintf("%s%v", cacheEnterpriseIdPrefix, data.Id)
	enterpriseNameKey := fmt.Sprintf("%s%v", cacheEnterpriseNamePrefix, data.Name)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rows.StringWithPlaceHolder())
		return conn.ExecCtx(ctx, query, *args.WithId(data.Id)...)
	}, enterpriseIdKey, enterpriseNameKey)
	return err
}

func (m *customEnterpriseModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	enterpriseIdKey := fmt.Sprintf("%s%v", cacheEnterpriseIdPrefix, id)
	enterpriseNameKey := fmt.Sprintf("%s%v", cacheEnterpriseNamePrefix, data.Name)
	return m.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		if err := m.DelCacheCtx(ctx, enterpriseIdKey, enterpriseNameKey); err != nil {
			return err
		}

		retrive := fmt.Sprintf("select %s from %s where `id` = ? limit 1", enterpriseRows, m.table)
		insert := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.Deleted.tableName(), deletedEnterpriseRowsExpectAutoSet)
		delete := fmt.Sprintf("delete from %s where `id` = ?", m.table)

		var resp Enterprise
		if err = s.QueryRowCtx(ctx, &resp, retrive, id); err != nil {
			return err
		}

		if _, err = s.ExecCtx(ctx, insert, resp.Id, resp.Name, resp.Address, resp.CreateTime, resp.Remark, resp.Version); err != nil {
			return err
		}

		_, err = s.ExecCtx(ctx, delete, resp.Id)
		return err
	})
}
