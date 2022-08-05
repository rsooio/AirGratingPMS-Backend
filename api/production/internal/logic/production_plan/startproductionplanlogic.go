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

type StartProductionPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartProductionPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartProductionPlanLogic {
	return &StartProductionPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartProductionPlanLogic) StartProductionPlan(req *types.StartProductionPlanReq) (resp *types.StartProductionPlanReply, err error) {
	info, err := l.svcCtx.ProductionPlanRpc.FindOneById(l.ctx, &ppb.FindOneByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	if info.State != production_plan.WaitProduction {
		return nil, errors.New("production plan not wait production")
	}

	_, err = l.svcCtx.ProductionPlanRpc.Update(l.ctx, &ppb.ProductionPlanInfo{
		Id:    req.Id,
		State: production_plan.InProduction,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.OrderRpc.UpdateStateByProductionId(l.ctx, &opb.UpdateStateByProductionIdReq{
		OldState:         order.WaitProduction,
		NewState:         order.InProduction,
		ProductionPlanId: req.Id,
	})
	if err != nil {
		l.svcCtx.ProductionPlanRpc.Update(l.ctx, &ppb.ProductionPlanInfo{
			Id:    req.Id,
			State: production_plan.WaitProduction,
		})
		return nil, err
	}

	return &types.StartProductionPlanReply{
		Message: "OK",
	}, err
}
