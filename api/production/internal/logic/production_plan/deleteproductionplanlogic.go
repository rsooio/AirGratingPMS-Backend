package production_plan

import (
	"context"

	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/rpc/production_plan/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductionPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductionPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductionPlanLogic {
	return &DeleteProductionPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductionPlanLogic) DeleteProductionPlan(req *types.DeleteProductionPlanReq) (resp *types.DeleteProductionPlanReply, err error) {
	// TODO: delete only if void

	_, err = l.svcCtx.ProductionPlanRpc.Delete(l.ctx, &pb.DeleteReq{
		Id: req.Id,
	})

	return &types.DeleteProductionPlanReply{
		Message: "OK",
	}, err
}
