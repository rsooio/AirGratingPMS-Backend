package workshop

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WorkshopModel = (*customWorkshopModel)(nil)

type (
	// WorkshopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWorkshopModel.
	WorkshopModel interface {
		workshopModel
		FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) ([]*Workshop, error)
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

func workshopRowsWithPrefix(prefix string) string {
	return prefix + "." + strings.Join(workshopFieldNames, ","+prefix+".")
}

func (m *customWorkshopModel) FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) ([]*Workshop, error) {
	resp := make([]*Workshop, 0)
	query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", workshopRowsWithPrefix("a"), m.table, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, offset, limit)
	if err != nil {
		logx.Errorf("workshopEnterpriseId.findList error, enterpriseId=%d, offset=%d, limit=%d, err=%s", enterpriseId, offset, limit, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return resp, nil
}
