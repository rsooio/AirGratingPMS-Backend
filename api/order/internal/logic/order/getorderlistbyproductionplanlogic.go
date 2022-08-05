package order

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListByProductionPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderListByProductionPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListByProductionPlanLogic {
	return &GetOrderListByProductionPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListByProductionPlanLogic) GetOrderListByProductionPlan(req *types.GetOrderListByProductionPlanReq) (resp *types.GetOrderListByProductionPlanReply, err error) {
	list, err := l.svcCtx.OrderRpc.FindListByProductionPlan(l.ctx, &pb.FindListByProductionPlanReq{
		ProductionPlanId: req.ProductionPlanId,
		PageSize:         req.PageSize,
		PageNumber:       req.PageNumber,
	})

	return &types.GetOrderListByProductionPlanReply{
		Message:     "OK",
		Count:       list.Count,
		ProductList: *list.ApiList(),
	}, err
}
