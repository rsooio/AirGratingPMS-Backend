package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/enterprise"
	"air-grating-pms-backend/rpc/enterprise/internal/svc"
	"air-grating-pms-backend/rpc/enterprise/pb"

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

func (l *PartialUpdateLogic) PartialUpdate(in *pb.EnterpriseInfo) (*pb.Empty, error) {
	now, err := l.svcCtx.EnterpriseModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	out := enterprise.Enterprise{
		Id: in.GetId(),
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

	return &pb.Empty{}, l.svcCtx.EnterpriseModel.Update(l.ctx, &out)
}
