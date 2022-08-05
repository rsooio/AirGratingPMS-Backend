package logic

import (
	"context"

	"air-grating-pms-backend/model/production_plan"
	"air-grating-pms-backend/rpc/production_plan/internal/svc"
	"air-grating-pms-backend/rpc/production_plan/pb"

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

func (l *UpdateLogic) Update(in *pb.ProductionPlanInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.ProductionPlanModel.Update(l.ctx, &production_plan.ProductionPlan{
		Id:           in.Id,
		EnterpriseId: in.EnterpriseId,
		WorkshopId:   in.WorkshopId,
		State:        in.State,
		Remark:       in.Remark,
	})
}
