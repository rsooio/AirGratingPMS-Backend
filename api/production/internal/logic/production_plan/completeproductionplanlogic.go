package production_plan

import (
	"context"
	"errors"

	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/model/order"
	"air-grating-pms-backend/model/production_plan"
	opb "air-grating-pms-backend/rpc/order/pb"
	ppb "air-grating-pms-backend/rpc/production_plan/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteProductionPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompleteProductionPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteProductionPlanLogic {
	return &CompleteProductionPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteProductionPlanLogic) CompleteProductionPlan(req *types.CompleteProductionPlanReq) (resp *types.CompleteProductionPlanReply, err error) {
	info, err := l.svcCtx.ProductionPlanRpc.FindOneById(l.ctx, &ppb.FindOneByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	if info.State != production_plan.InProduction {
		return nil, errors.New("production plan not in product")
	}

	_, err = l.svcCtx.ProductionPlanRpc.Update(l.ctx, &ppb.ProductionPlanInfo{
		Id:    req.Id,
		State: production_plan.ProductionDone,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.OrderRpc.UpdateStateByProductionId(l.ctx, &opb.UpdateStateByProductionIdReq{
		OldState:         order.InProduction,
		NewState:         order.ProductionDone,
		ProductionPlanId: req.Id,
	})
	if err != nil {
		l.svcCtx.ProductionPlanRpc.Update(l.ctx, &ppb.ProductionPlanInfo{
			Id:    req.Id,
			State: production_plan.InProduction,
		})
		return nil, err
	}

	return &types.CompleteProductionPlanReply{
		Message: "OK",
	}, err
}
