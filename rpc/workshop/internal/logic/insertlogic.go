package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/workshop"
	"air-grating-pms-backend/rpc/workshop/internal/svc"
	"air-grating-pms-backend/rpc/workshop/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLogic) Insert(in *pb.WorkshopInfo) (*pb.Empty, error) {
	_, err := l.svcCtx.WorkshopModel.Insert(l.ctx, &workshop.Workshop{
		EnterpriseId: in.GetEnterpriseId(),
		Name:         in.GetName(),
		Address:      sql.NullString{String: in.GetAddress(), Valid: in.GetAddress() != ""},
		PhoneNumber:  sql.NullString{String: in.GetPhoneNumber(), Valid: in.GetPhoneNumber() != ""},
		ManagerId:    sql.NullInt64{Int64: in.GetManagerId(), Valid: in.GetManagerId() != 0},
		Remark:       sql.NullString{String: in.GetRemark(), Valid: in.GetRemark() != ""},
		Version:      in.GetVersion(),
	})

	return &pb.Empty{}, err
}
