package logic

import (
	"context"

	"air-grating-pms-backend/rpc/order/internal/svc"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindListByProductionPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindListByProductionPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindListByProductionPlanLogic {
	return &FindListByProductionPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindListByProductionPlanLogic) FindListByProductionPlan(in *pb.FindListByProductionPlanReq) (*pb.OrderList, error) {
	list, count, err := l.svcCtx.OrderModel.FindListByProductionPlan(l.ctx, in.ProductionPlanId, (in.PageNumber-1)*in.PageSize, in.PageSize)

	return &pb.OrderList{
		Count: count,
		List:  *list.RpcList(),
	}, err
}
