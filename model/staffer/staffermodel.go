package staffer

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StafferModel = (*customStafferModel)(nil)

type (
	// StafferModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStafferModel.
	StafferModel interface {
		stafferModel
		FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) ([]*Staffer, error)
		FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) ([]*Staffer, error)
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

func stafferRowsWithPrefix(prefix string) string {
	return prefix + "." + strings.Join(stafferFieldNames, ","+prefix+".")
}

func (m *customStafferModel) FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) ([]*Staffer, error) {
	resp := make([]*Staffer, 0)
	query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", stafferRowsWithPrefix("a"), m.table, m.table)
	// query := fmt.Sprintf("select %s from %s where `enterprise_id` = ? limit 1", stafferRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, offset, limit)
	if err != nil {
		logx.Errorf("staffersEnterpriseId.findList error, enterpriseId=%d, offset=%d, limit=%d, err=%s", enterpriseId, offset, limit, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return resp, nil
}

func (m *customStafferModel) FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) ([]*Staffer, error) {
	resp := make([]*Staffer, 0)
	query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ? LIMIT ?, ?) AS b ON a.id = b.id", stafferRowsWithPrefix("a"), m.table, m.table)
	// query := fmt.Sprintf("select %s from %s where `enterprise_id` = ? limit 1", stafferRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, workshopId, offset, limit)
	if err != nil {
		logx.Errorf("staffersWorkshopId.findList error, enterpriseId=%d, workshopId=%d, offset=%d, limit=%d, err=%s", enterpriseId, workshopId, offset, limit, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return resp, nil
}
