package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/enterprise"
	"air-grating-pms-backend/rpc/enterprise/internal/svc"
	"air-grating-pms-backend/rpc/enterprise/pb"

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

func (l *UpdateLogic) Update(in *pb.EnterpriseInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.EnterpriseModel.Update(l.ctx, &enterprise.Enterprise{
		Id:      in.GetId(),
		Name:    in.GetName(),
		Address: sql.NullString{String: in.GetAddress(), Valid: true},
		Remark:  sql.NullString{String: in.GetRemark(), Valid: true},
		Version: in.GetVersion(),
	})
}
