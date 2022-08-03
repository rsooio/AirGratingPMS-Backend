package technology

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

var _ TechnologyModel = (*customTechnologyModel)(nil)

var (
	cacheTechnologyEnterpriseListCountPrefix = "cache:technology:enterprise:list:count:"
	cacheTechnologyWorkshopListCountPrefix   = "cache:technology:workshop:list:count:"
	technologyRowsWithPrefixA                = "a." + strings.Join(technologyFieldNames, ",a.")
)

type (
	// TechnologyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTechnologyModel.
	TechnologyModel interface {
		technologyModel
		Insert(ctx context.Context, data *Technology) (sql.Result, error)
		Update(ctx context.Context, data *Technology) error
		Delete(ctx context.Context, id int64) error
		FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (TechnologyList, int64, error)
		FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (TechnologyList, int64, error)
	}

	customTechnologyModel struct {
		*defaultTechnologyModel
	}
)

// NewTechnologyModel returns a model for the database table.
func NewTechnologyModel(conn sqlx.SqlConn, c cache.CacheConf) TechnologyModel {
	return &customTechnologyModel{
		defaultTechnologyModel: newTechnologyModel(conn, c),
	}
}

func cacheTechnologyEnterpriseListCountKey(enterpriseId int64) string {
	return fmt.Sprintf("%s%v", cacheTechnologyEnterpriseListCountPrefix, enterpriseId)
}

func cacheTechnologyWorkshopListCountKey(enterpriseId, workshopId int64) string {
	return fmt.Sprintf("%s%v:%v", cacheTechnologyWorkshopListCountPrefix, enterpriseId, workshopId)
}

func (m *customTechnologyModel) Insert(ctx context.Context, data *Technology) (sql.Result, error) {
	technologyEnterpriseIdWorkshopIdNameKey := fmt.Sprintf("%s%v:%v:%v", cacheTechnologyEnterpriseIdWorkshopIdNamePrefix, data.EnterpriseId, data.WorkshopId, data.Name)
	technologyIdKey := fmt.Sprintf("%s%v", cacheTechnologyIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, technologyRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.EnterpriseId, data.WorkshopId, data.Name, data.Info, data.Remark, data.Version)
	}, technologyEnterpriseIdWorkshopIdNameKey, technologyIdKey, cacheTechnologyEnterpriseListCountKey(data.EnterpriseId), cacheTechnologyWorkshopListCountKey(data.EnterpriseId, data.WorkshopId))
	return ret, err
}

func (m *customTechnologyModel) Update(ctx context.Context, data *Technology) error {
	rows, args := partial.Partial(data)

	technologyEnterpriseIdWorkshopIdNameKey := fmt.Sprintf("%s%v:%v:%v", cacheTechnologyEnterpriseIdWorkshopIdNamePrefix, data.EnterpriseId, data.WorkshopId, data.Name)
	technologyIdKey := fmt.Sprintf("%s%v", cacheTechnologyIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rows.StringWithPlaceHolder())
		return conn.ExecCtx(ctx, query, *args.WithId(data.Id)...)
	}, technologyEnterpriseIdWorkshopIdNameKey, technologyIdKey, cacheTechnologyEnterpriseListCountKey(data.EnterpriseId), cacheTechnologyWorkshopListCountKey(data.EnterpriseId, data.WorkshopId))
	return err
}

func (m *customTechnologyModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	technologyEnterpriseIdWorkshopIdNameKey := fmt.Sprintf("%s%v:%v:%v", cacheTechnologyEnterpriseIdWorkshopIdNamePrefix, data.EnterpriseId, data.WorkshopId, data.Name)
	technologyIdKey := fmt.Sprintf("%s%v", cacheTechnologyIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, technologyEnterpriseIdWorkshopIdNameKey, technologyIdKey, cacheTechnologyEnterpriseListCountKey(data.EnterpriseId), cacheTechnologyWorkshopListCountKey(data.EnterpriseId, data.WorkshopId))
	return err
}

func (m *customTechnologyModel) FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (TechnologyList, int64, error) {
	resp := make(TechnologyList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", technologyRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, offset, limit)
		if err != nil {
			logx.Errorf("technologyEnterpriseId.findList error, enterpriseId=%d, offset=%d, limit=%d, err=%s", enterpriseId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ?", m.table)
		return m.QueryRow(&count, cacheTechnologyEnterpriseListCountKey(enterpriseId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId)
		})
	})

	return resp, count, err
}

func (m *customTechnologyModel) FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (TechnologyList, int64, error) {
	resp := make(TechnologyList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ? LIMIT ?, ?) AS b ON a.id = b.id", technologyRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, workshopId, offset, limit)
		if err != nil {
			logx.Errorf("technologyWorkshopId.findList error, enterpriseId=%d, workshopId=%d, offset=%d, limit=%d, err=%s", enterpriseId, workshopId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ?", m.table)
		return m.QueryRow(&count, cacheTechnologyWorkshopListCountKey(enterpriseId, workshopId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId, workshopId)
		})
	})

	return resp, count, err
}
