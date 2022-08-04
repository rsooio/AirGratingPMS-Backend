package staffer

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

var _ StafferModel = (*customStafferModel)(nil)

var (
	cacheStafferWorkshopListCountPrefix   = "cache:order:enterprise:list:count:"
	cacheStafferEnterpriseListCountPrefix = "cache:order:enterprise:list:count:"
	stafferRowsWithPrefixA                = "a." + strings.Join(stafferFieldNames, ",a.")
)

type (
	// StafferModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStafferModel.
	StafferModel interface {
		stafferModel
		Insert(ctx context.Context, data *Staffer) (sql.Result, error)
		Update(ctx context.Context, data *Staffer) error
		Delete(ctx context.Context, id int64) error
		FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (*StafferList, int64, error)
		FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (*StafferList, int64, error)
	}

	customStafferModel struct {
		*defaultStafferModel
	}
)

// NewStafferModel returns a model for the database table.
func NewStafferModel(conn sqlx.SqlConn, c cache.CacheConf) StafferModel {
	return &customStafferModel{
		defaultStafferModel: newStafferModel(conn, c),
	}
}

func cacheStafferEnterpriseListCountKey(enterpriseId int64) string {
	return fmt.Sprintf("%s%v", cacheStafferEnterpriseListCountPrefix, enterpriseId)
}

func cacheStafferWorkshopListCountKey(enterpriseId, workshopId int64) string {
	return fmt.Sprintf("%s%v:%v", cacheStafferWorkshopListCountPrefix, enterpriseId, workshopId)
}

func (m *customStafferModel) Insert(ctx context.Context, data *Staffer) (sql.Result, error) {
	stafferEnterpriseIdUsernameKey := fmt.Sprintf("%s%v:%v", cacheStafferEnterpriseIdUsernamePrefix, data.EnterpriseId, data.Username)
	stafferIdKey := fmt.Sprintf("%s%v", cacheStafferIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, stafferRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.EnterpriseId, data.WorkshopId, data.Username, data.Role, data.Name, data.HashedPassword, data.Gender, data.PhoneNumber, data.Email, data.Address, data.ExpireTime, data.Remark, data.Version)
	}, stafferEnterpriseIdUsernameKey, stafferIdKey, cacheStafferEnterpriseListCountKey(data.EnterpriseId), cacheStafferWorkshopListCountKey(data.EnterpriseId, data.WorkshopId))
	return ret, err
}

func (m *customStafferModel) Update(ctx context.Context, data *Staffer) error {
	rows, args := partial.Partial(data)

	keys := []string{fmt.Sprintf("%s%v", cacheStafferIdPrefix, data.Id)}
	if data.WorkshopId != 0 || data.EnterpriseId != 0 {
		info, err := m.FindOne(ctx, data.Id)
		if err != nil {
			return err
		}

		keys = append(*partial.UpdateKeys2x1_12(data.EnterpriseId, data.WorkshopId, info.EnterpriseId, info.WorkshopId, cacheStafferEnterpriseListCountKey, cacheStafferWorkshopListCountKey), keys...)
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rows.StringWithPlaceHolder())
		return conn.ExecCtx(ctx, query, *args.WithId(data.Id)...)
	}, keys...)
	return err
}

func (m *customStafferModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	stafferEnterpriseIdUsernameKey := fmt.Sprintf("%s%v:%v", cacheStafferEnterpriseIdUsernamePrefix, data.EnterpriseId, data.Username)
	stafferIdKey := fmt.Sprintf("%s%v", cacheStafferIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, stafferEnterpriseIdUsernameKey, stafferIdKey, cacheStafferEnterpriseListCountKey(data.EnterpriseId), cacheStafferWorkshopListCountKey(data.EnterpriseId, data.WorkshopId))
	return err
}

func (m *customStafferModel) FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (*StafferList, int64, error) {
	resp := make(StafferList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", stafferRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, offset, limit)
		if err != nil {
			logx.Errorf("stafferEnterpriseId.findList error, enterpriseId=%d, offset=%d, limit=%d, err=%s", enterpriseId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ?", m.table)
		return m.QueryRow(&count, cacheStafferEnterpriseListCountKey(enterpriseId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId)
		})
	})

	return &resp, count, err
}

func (m *customStafferModel) FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (*StafferList, int64, error) {
	resp := make(StafferList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", stafferRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, workshopId, offset, limit)
		if err != nil {
			logx.Errorf("stafferWorkshopId.findList error, enterpriseId=%d, workshopId=%d, offset=%d, limit=%d, err=%s", enterpriseId, workshopId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ?", m.table)
		return m.QueryRow(&count, cacheStafferWorkshopListCountKey(enterpriseId, workshopId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId)
		})
	})

	return &resp, count, err
}
