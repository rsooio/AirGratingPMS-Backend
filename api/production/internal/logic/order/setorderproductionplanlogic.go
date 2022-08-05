package order

import (
	"context"
	"errors"

	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetOrderProductionPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetOrderProductionPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetOrderProductionPlanLogic {
	return &SetOrderProductionPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetOrderProductionPlanLogic) SetOrderProductionPlan(req *types.SetOrderProductionPlanReq) (resp *types.SetOrderProductionPlanReply, err error) {
	info, err := l.svcCtx.OrderRpc.FindOneById(l.ctx, &pb.FindOneByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	if info.ProductionPlanId != 0 {
		return nil, errors.New("production plan already exist")
	}

	_, err = l.svcCtx.OrderRpc.Update(l.ctx, &pb.OrderInfo{
		Id:               req.Id,
		ProductionPlanId: req.ProductionPlanId,
	})

	return &types.SetOrderProductionPlanReply{
		Message: "OK",
	}, err
}
