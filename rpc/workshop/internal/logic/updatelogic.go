package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/workshop"
	"air-grating-pms-backend/rpc/workshop/internal/svc"
	"air-grating-pms-backend/rpc/workshop/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *pb.WorkshopInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.WorkshopModel.Update(l.ctx, &workshop.Workshop{
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
