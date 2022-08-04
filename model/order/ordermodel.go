package order

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

var _ OrderModel = (*customOrderModel)(nil)

var (
	cacheOrderWorkshopListCountPrefix   = "cache:order:workshop:list:count:"
	cacheOrderEnterpriseListCountPrefix = "cache:order:enterprise:list:count:"
	orderRowsWithPrefixA                = "a." + strings.Join(orderFieldNames, ",a.")
)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		Insert(ctx context.Context, data *Order) (sql.Result, error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, data *Order) error
		FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (OrderList, int64, error)
		FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (OrderList, int64, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c),
	}
}

func cacheOrderWorkshopListCountKey(enterpriseId, workshopId int64) string {
	return fmt.Sprintf("%s%v:%v", cacheOrderWorkshopListCountPrefix, enterpriseId, workshopId)
}

func cacheOrderEnterpriseListCountKey(enterpriseId int64) string {
	return fmt.Sprintf("%s%v", cacheOrderEnterpriseListCountPrefix, enterpriseId)
}

func (m *customOrderModel) Insert(ctx context.Context, data *Order) (sql.Result, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.EnterpriseId, data.WorkshopId, data.ClientId, data.ProductionPlanId, data.State, data.Address, data.Linkman, data.PhoneNumber, data.Email, data.CorrespondingCode, data.Remark, data.Version)
	}, orderIdKey, cacheOrderEnterpriseListCountKey(data.EnterpriseId), cacheOrderWorkshopListCountKey(data.EnterpriseId, data.WorkshopId))
	return ret, err
}

func (m *customOrderModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, orderIdKey, cacheOrderEnterpriseListCountKey(data.EnterpriseId), cacheOrderWorkshopListCountKey(data.EnterpriseId, data.WorkshopId))
	return err
}

func (m *customOrderModel) Update(ctx context.Context, data *Order) error {
	rows, args := partial.Partial(data)

	keys := []string{fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)}
	if data.WorkshopId != 0 || data.EnterpriseId != 0 {
		info, err := m.FindOne(ctx, data.Id)
		if err != nil {
			return err
		}

		keys = append(*partial.UpdateKeys2x1_12(data.EnterpriseId, data.WorkshopId, info.EnterpriseId, info.WorkshopId, cacheOrderEnterpriseListCountKey, cacheOrderWorkshopListCountKey), keys...)
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rows.StringWithPlaceHolder())
		return conn.ExecCtx(ctx, query, *args.WithId(data.Id)...)
	}, keys...)
	return err
}

func (m *customOrderModel) FindListByEnterprise(ctx context.Context, enterpriseId int64, offset, limit int32) (OrderList, int64, error) {
	resp := make(OrderList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? ORDER BY `id` desc LIMIT ?, ?) AS b ON a.id = b.id", orderRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, offset, limit)
		if err != nil {
			logx.Errorf("orderEnterpriseId.findList error, enterpriseId=%d, offset=%d, limit=%d, err=%s", enterpriseId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ?", m.table)
		return m.QueryRow(&count, cacheOrderEnterpriseListCountKey(enterpriseId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId)
		})
	})

	return resp, count, err
}

func (m *customOrderModel) FindListByWorkshop(ctx context.Context, enterpriseId int64, workshopId int64, offset, limit int32) (OrderList, int64, error) {
	resp := make(OrderList, 0)
	var count int64

	err := mr.Finish(func() error {
		query := fmt.Sprintf("SELECT %s FROM %s AS a INNER JOIN (SELECT `id` FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ? LIMIT ?, ?) AS b ON a.id = b.id", orderRowsWithPrefixA, m.table, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, enterpriseId, workshopId, offset, limit)
		if err != nil {
			logx.Errorf("orderWorkshopId.findList error, enterpriseId=%d, workshopId=%d, offset=%d, limit=%d, err=%s", enterpriseId, workshopId, offset, limit, err.Error())
			if err == sqlx.ErrNotFound {
				return ErrNotFound
			}
			return err
		}
		return nil
	}, func() error {
		query := fmt.Sprintf("SELECT count(*) FROM %s WHERE `enterprise_id` = ? AND `workshop_id` = ?", m.table)
		return m.QueryRow(&count, cacheOrderWorkshopListCountKey(enterpriseId, workshopId), func(conn sqlx.SqlConn, v interface{}) error {
			return conn.QueryRowCtx(ctx, v, query, enterpriseId, workshopId)
		})
	})

	return resp, count, err
}
