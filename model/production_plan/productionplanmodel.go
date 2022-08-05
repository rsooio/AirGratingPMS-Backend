package production_plan

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

var _ ProductionPlanModel = (*customProductionPlanModel)(nil)

var (
	cacheProductionPlanWorkshopListCountPrefix   = "cache:production_plan:workshop:list:count:"
	cacheProductionPlanEnterpriseListCountPrefix = "cache:production_plan:enterprise:list:count:"
	productionPlanRowsWithPrefixA                = "a." + strings.Join(productionPlanFieldNames, ",a.")
)

type (
	// ProductionPlanModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductionPlanModel.
	ProductionPlanModel interface {
		productionPlanModel
		Insert(ctx context.Context, data *ProductionPlan) (sql.Result, error)
		Update(ctx context.Context, data *ProductionPlan) error
		Delete(ctx context.Context, id int64) error
		FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (ProductionPlanList, int64, error)
		FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (ProductionPlanList, int64, error)
	}

	customProductionPlanModel struct {
		*defaultProductionPlanModel
	}
)

// NewProductionPlanModel returns a model for the database table.
func NewProductionPlanModel(conn sqlx.SqlConn, c cache.CacheConf) ProductionPlanModel {
	return &customProductionPlanModel{
		defaultProductionPlanModel: newProductionPlanModel(conn, c),
	}
}

func cacheProductionPlanWorkshopListCountKey(enterpriseId, workshopId int64) string {
	return fmt.Sprintf("%s%v:%v", cacheProductionPlanWorkshopListCountPrefix, enterpriseId, workshopId)
}

func cacheProductionPlanEnterpriseListCountKey(enterpriseId int64) string {
	return fmt.Sprintf("%s%v", cacheProductionPlanEnterpriseListCountPrefix, enterpriseId)
}

func (m *customProductionPlanModel) Insert(ctx context.Context, data *ProductionPlan) (sql.Result, error) {
	productionPlanIdKey := fmt.Sprintf("%s%v", cacheProductionPlanIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, productionPlanRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.EnterpriseId, data.WorkshopId, data.State, data.Remark, data.Version)
	}, productionPlanIdKey, cacheProductionPlanEnterpriseListCountKey(data.EnterpriseId), cacheProductionPlanWorkshopListCountKey(data.EnterpriseId, data.WorkshopId))
	return ret, err
}

func (m *customProductionPlanModel) Update(ctx context.Context, data *ProductionPlan) error {
	rows, args := partial.Partial(data)

	keys := []string{fmt.Sprintf("%s%v", cacheProductionPlanIdPrefix, data.Id)}
	if data.WorkshopId != 0 || data.EnterpriseId != 0 {
		info, err := m.FindOne(ctx, data.Id)
		if err != nil {
			return err
		}

		keys = append(*partial.UpdateKeys2x1_12(data.EnterpriseId, data.WorkshopId, info.EnterpriseId, info.WorkshopId, cacheProductionPlanEnterpriseListCountKey, cacheProductionPlanWorkshopListCountKey), keys...)
	}
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rows.StringWithPlaceHolder())
		return conn.ExecCtx(ctx, query, *args.WithId(data.Id)...)
	}, keys...)
	return err
}

func (m *customProductionPlanModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	productionPlanIdKey := fmt.Sprintf("%s%v", cacheProductionPlanIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, productionPlanIdKey, cacheProductionPlanEnterpriseListCountKey(data.EnterpriseId), cacheProductionPlanWorkshopListCountKey(data.EnterpriseId, data.WorkshopId))
	return err
}

func (m *customProductionPlanModel) FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (ProductionPlanList, int64, error) {
	resp := make(ProductionPlanList, 0, limit)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", productionPlanRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, offset, limit)
		if err != nil {
			logx.Errorf("prodcutionPlanEnterpriseId.findList error, enterpriseId=%d, offset=%d, limit=%d, err=%s", enterpriseId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ?", m.table)
		return m.QueryRow(&count, cacheProductionPlanEnterpriseListCountKey(enterpriseId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId)
		})
	})

	return resp, count, err
}

func (m *customProductionPlanModel) FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (ProductionPlanList, int64, error) {
	resp := make(ProductionPlanList, 0, limit)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ? LIMIT ?, ?) AS b ON a.id = b.id", productionPlanRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, workshopId, offset, limit)
		if err != nil {
			logx.Errorf("prodcutionPlanWorkshopId.findList error, enterpriseId=%d, workshopId=%d, offset=%d, limit=%d, err=%s", enterpriseId, workshopId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ?", m.table)
		return m.QueryRow(&count, cacheProductionPlanWorkshopListCountKey(enterpriseId, workshopId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId, workshopId)
		})
	})

	return resp, count, err
}
