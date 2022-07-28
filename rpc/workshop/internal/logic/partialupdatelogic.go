package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/workshop"
	"air-grating-pms-backend/rpc/workshop/internal/svc"
	"air-grating-pms-backend/rpc/workshop/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartialUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPartialUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartialUpdateLogic {
	return &PartialUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PartialUpdateLogic) PartialUpdate(in *pb.WorkshopInfo) (*pb.Empty, error) {
	now, err := l.svcCtx.WorkshopModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	out := workshop.Workshop{
		Id: in.GetId(),
	}

	if in.GetEnterpriseId() == out.EnterpriseId {
		out.EnterpriseId = now.EnterpriseId
	} else {
		out.EnterpriseId = in.GetEnterpriseId()
	}
	if in.GetName() == out.Name {
		out.Name = now.Name
	} else {
		out.Name = in.GetName()
	}
	if in.GetAddress() == out.Address.String {
		out.Address = now.Address
	} else {
		out.Address = sql.NullString{String: in.GetAddress(), Valid: in.GetAddress() != ""}
	}
	if in.GetPhoneNumber() == out.PhoneNumber.String {
		out.PhoneNumber = now.PhoneNumber
	} else {
		out.PhoneNumber = sql.NullString{String: in.GetPhoneNumber(), Valid: in.GetPhoneNumber() != ""}
	}
	if in.GetManagerId() == out.ManagerId.Int64 {
		out.ManagerId = now.ManagerId
	} else {
		out.ManagerId = sql.NullInt64{Int64: in.GetManagerId(), Valid: in.GetManagerId() != 0}
	}
	if in.GetRemark() == out.Remark.String {
		out.Remark = now.Remark
	} else {
		out.Remark = sql.NullString{String: in.GetRemark(), Valid: in.GetRemark() != ""}
	}
	if in.GetVersion() == out.Version {
		out.Version = now.Version
	} else {
		out.Version = in.GetVersion()
	}

	return &pb.Empty{}, l.svcCtx.WorkshopModel.Update(l.ctx, &out)
}
