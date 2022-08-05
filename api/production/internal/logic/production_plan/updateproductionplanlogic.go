package production_plan

import (
	"context"

	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/rpc/production_plan/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductionPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductionPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductionPlanLogic {
	return &UpdateProductionPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductionPlanLogic) UpdateProductionPlan(req *types.UpdateProductionPlanReq) (resp *types.UpdateProductionPlanReply, err error) {
	_, err = l.svcCtx.ProductionPlanRpc.Update(l.ctx, &pb.ProductionPlanInfo{
		Id:         req.Id,
		WorkshopId: utils.ZeroSelectWorkshopId(l.ctx, req.WorkshopId),
		Remark:     req.Remark,
	})

	return &types.UpdateProductionPlanReply{
		Message: "OK",
	}, err
}
