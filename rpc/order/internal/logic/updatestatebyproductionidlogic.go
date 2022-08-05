package logic

import (
	"context"

	"air-grating-pms-backend/rpc/order/internal/svc"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStateByProductionIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStateByProductionIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStateByProductionIdLogic {
	return &UpdateStateByProductionIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStateByProductionIdLogic) UpdateStateByProductionId(in *pb.UpdateStateByProductionIdReq) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.OrderModel.UpdateStateByProductionId(l.ctx, in.ProductionPlanId, in.OldState, in.NewState)
}
