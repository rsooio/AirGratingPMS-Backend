package client

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

var _ ClientModel = (*customClientModel)(nil)

var (
	cacheClientWorkshopListCountPrefix   = "cache:client:workshop:list:count:"
	cacheClientEnterpriseListCountPrefix = "cache:client:enterprise:list:count:"
	clientRowsWithPrefixA                = "a." + strings.Join(clientFieldNames, ",a.")
)

type (
	// ClientModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClientModel.
	ClientModel interface {
		clientModel
		Insert(ctx context.Context, data *Client) (sql.Result, error)
		Update(ctx context.Context, data *Client) error
		Delete(ctx context.Context, id int64) error
		FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (ClientList, int64, error)
		FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (ClientList, int64, error)
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

func cacheClientEnterpriseListCountKey(enterpriseId int64) string {
	return fmt.Sprintf("%s%v", cacheClientEnterpriseListCountPrefix, enterpriseId)
}

func cacheClientWorkshopListCountKey(enterpriseId, workshopId int64) string {
	return fmt.Sprintf("%s%v:%v", cacheClientWorkshopListCountPrefix, enterpriseId, workshopId)
}

func (m *customClientModel) Insert(ctx context.Context, data *Client) (sql.Result, error) {
	clientEnterpriseIdWorkshopIdNameKey := fmt.Sprintf("%s%v:%v:%v", cacheClientEnterpriseIdWorkshopIdNamePrefix, data.EnterpriseId, data.WorkshopId, data.Name)
	clientIdKey := fmt.Sprintf("%s%v", cacheClientIdPrefix, data.Id)
	ret, err := m.ExecCtx(
		ctx,
		func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
			query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, clientRowsExpectAutoSet)
			return conn.ExecCtx(ctx, query, data.EnterpriseId, data.WorkshopId, data.Name, data.PhoneNumber, data.Email, data.Address, data.Remark, data.Version)
		},
		clientEnterpriseIdWorkshopIdNameKey,
		clientIdKey,
		cacheClientEnterpriseListCountKey(data.EnterpriseId),
		cacheClientWorkshopListCountKey(data.EnterpriseId, data.WorkshopId),
	)
	return ret, err
}

func (m *customClientModel) Update(ctx context.Context, data *Client) error {
	rows, args := partial.Partial(data)

	clientEnterpriseIdWorkshopIdNameKey := fmt.Sprintf("%s%v:%v:%v", cacheClientEnterpriseIdWorkshopIdNamePrefix, data.EnterpriseId, data.WorkshopId, data.Name)
	clientIdKey := fmt.Sprintf("%s%v", cacheClientIdPrefix, data.Id)
	keys := []string{clientEnterpriseIdWorkshopIdNameKey, clientIdKey}
	if data.WorkshopId != 0 || data.EnterpriseId != 0 {
		info, err := m.FindOne(ctx, data.Id)
		if err != nil {
			return err
		}

		keys = append(keys, *partial.UpdateKeys2x1_12(data.EnterpriseId, data.WorkshopId, info.EnterpriseId, info.WorkshopId, cacheClientEnterpriseListCountKey, cacheClientWorkshopListCountKey)...)
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rows.StringWithPlaceHolder())
		return conn.ExecCtx(ctx, query, *args.WithId(data.Id)...)
	}, keys...)
	return err
}

func (m *customClientModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	clientEnterpriseIdWorkshopIdNameKey := fmt.Sprintf("%s%v:%v:%v", cacheClientEnterpriseIdWorkshopIdNamePrefix, data.EnterpriseId, data.WorkshopId, data.Name)
	clientIdKey := fmt.Sprintf("%s%v", cacheClientIdPrefix, id)
	_, err = m.ExecCtx(
		ctx,
		func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
			query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
			return conn.ExecCtx(ctx, query, id)
		},
		clientEnterpriseIdWorkshopIdNameKey,
		clientIdKey,
		cacheClientEnterpriseListCountKey(data.EnterpriseId),
		cacheClientWorkshopListCountKey(data.EnterpriseId, data.WorkshopId),
	)
	return err
}

func (m *customClientModel) FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (ClientList, int64, error) {
	resp := make(ClientList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", clientRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, offset, limit)
		if err != nil {
			logx.Errorf("clientEnterpriseId.findList error, enterpriseId=%d, offset=%d, limit=%d, err=%s", enterpriseId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ?", m.table)
		return m.QueryRow(&count, cacheClientEnterpriseListCountKey(enterpriseId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId)
		})
	})

	return resp, count, err
}

func (m *customClientModel) FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (ClientList, int64, error) {
	resp := make(ClientList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ? LIMIT ?, ?) AS b ON a.id = b.id", clientRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, workshopId, offset, limit)
		if err != nil {
			logx.Errorf("clientWorkshopId.findList error, enterpriseId=%d, workshopId=%d, offset=%d, limit=%d, err=%s", enterpriseId, workshopId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ?", m.table)
		return m.QueryRow(&count, cacheClientWorkshopListCountKey(enterpriseId, workshopId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId, workshopId)
		})
	})

	return resp, count, err
}
