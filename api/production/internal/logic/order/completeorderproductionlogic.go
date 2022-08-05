package order

import (
	"context"
	"errors"

	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/model/order"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteOrderProductionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompleteOrderProductionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteOrderProductionLogic {
	return &CompleteOrderProductionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteOrderProductionLogic) CompleteOrderProduction(req *types.CompleteOrderProductionReq) (resp *types.CompleteOrderProductionReply, err error) {
	// TODO: Try to complete Production Plan

	info, err := l.svcCtx.OrderRpc.FindOneById(l.ctx, &pb.FindOneByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	if info.State != order.InProduction {
		return nil, errors.New("order not in production")
	}

	_, err = l.svcCtx.OrderRpc.Update(l.ctx, &pb.OrderInfo{
		Id:    req.Id,
		State: order.ProductionDone,
	})

	return &types.CompleteOrderProductionReply{
		Message: "OK",
	}, err
}
