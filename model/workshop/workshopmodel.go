package workshop

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

var _ WorkshopModel = (*customWorkshopModel)(nil)

var (
	cacheWorkshopEnterpriseListCountPrefix = "cache:workshop:enterprise:list:count:"
	workshopRowsWithPrefixA                = "a." + strings.Join(workshopFieldNames, ",a.")
)

type (
	// WorkshopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWorkshopModel.
	WorkshopModel interface {
		workshopModel
		Insert(ctx context.Context, data *Workshop) (sql.Result, error)
		Update(ctx context.Context, data *Workshop) error
		Delete(ctx context.Context, id int64) error
		FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (WorkshopList, int64, error)
	}

	customWorkshopModel struct {
		*defaultWorkshopModel
	}
)

// NewWorkshopModel returns a model for the database table.
func NewWorkshopModel(conn sqlx.SqlConn, c cache.CacheConf) WorkshopModel {
	return &customWorkshopModel{
		defaultWorkshopModel: newWorkshopModel(conn, c),
	}
}

func cacheWorkshopEnterpriseListCountKey(enterpriseId int64) string {
	return fmt.Sprintf("%s%v", cacheWorkshopEnterpriseListCountPrefix, enterpriseId)
}

func (m *customWorkshopModel) Insert(ctx context.Context, data *Workshop) (sql.Result, error) {
	workshopEnterpriseIdNameKey := fmt.Sprintf("%s%v:%v", cacheWorkshopEnterpriseIdNamePrefix, data.EnterpriseId, data.Name)
	workshopIdKey := fmt.Sprintf("%s%v", cacheWorkshopIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, workshopRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.EnterpriseId, data.Name, data.Address, data.PhoneNumber, data.ManagerId, data.Remark, data.Version)
	}, workshopEnterpriseIdNameKey, workshopIdKey, cacheWorkshopEnterpriseListCountKey(data.EnterpriseId))
	return ret, err
}

func (m *customWorkshopModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	workshopEnterpriseIdNameKey := fmt.Sprintf("%s%v:%v", cacheWorkshopEnterpriseIdNamePrefix, data.EnterpriseId, data.Name)
	workshopIdKey := fmt.Sprintf("%s%v", cacheWorkshopIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, workshopEnterpriseIdNameKey, workshopIdKey, cacheWorkshopEnterpriseListCountKey(data.EnterpriseId))
	return err
}

func (m *customWorkshopModel) Update(ctx context.Context, data *Workshop) error {
	rows, args := partial.Partial(data)

	workshopEnterpriseIdNameKey := fmt.Sprintf("%s%v:%v", cacheWorkshopEnterpriseIdNamePrefix, data.EnterpriseId, data.Name)
	workshopIdKey := fmt.Sprintf("%s%v", cacheWorkshopIdPrefix, data.Id)
	keys := []string{workshopEnterpriseIdNameKey, workshopIdKey}
	if data.EnterpriseId != 0 {
		info, err := m.FindOne(ctx, data.Id)
		if err != nil {
			return err
		}

		keys = append(*partial.UpdateKeys1x1(data.EnterpriseId, info.EnterpriseId, cacheWorkshopEnterpriseListCountKey), keys...)
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rows.StringWithPlaceHolder())
		return conn.ExecCtx(ctx, query, *args.WithId(data.Id)...)
	}, keys...)
	return err
}

func (m *customWorkshopModel) FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (WorkshopList, int64, error) {
	resp := make(WorkshopList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", workshopRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, offset, limit)
		if err != nil {
			logx.Errorf("workshopEnterpriseId.findList error, enterpriseId=%d, offset=%d, limit=%d, err=%s", enterpriseId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ?", m.table)
		return m.QueryRow(&count, cacheWorkshopEnterpriseListCountKey(enterpriseId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId)
		})
	})

	return resp, count, err
}
