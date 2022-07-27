package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/workshop"
	"air-grating-pms-backend/rpc/workshop/internal/svc"
	"air-grating-pms-backend/rpc/workshop/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeLogic {
	return &ChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeLogic) Change(in *types.WorkshopInfoWithId) (*types.Empty, error) {
	return &types.Empty{}, l.svcCtx.WorkshopModel.Update(l.ctx, &workshop.Workshop{
		Id:           in.GetId(),
		EnterpriseId: in.GetEnterpriseId(),
		Name:         in.GetName(),
		Address:      sql.NullString{String: in.GetAddress(), Valid: true},
		PhoneNumber:  sql.NullString{String: in.GetPhoneNumber(), Valid: true},
		ManagerId:    sql.NullInt64{Int64: in.GetManagerId(), Valid: true},
		Remark:       sql.NullString{String: in.GetRemark(), Valid: true},
		Version:      in.GetVersion(),
	})
}
