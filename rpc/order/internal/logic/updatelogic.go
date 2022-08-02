package logic

import (
	"context"

	"air-grating-pms-backend/model/order"
	"air-grating-pms-backend/rpc/order/internal/svc"
	"air-grating-pms-backend/rpc/order/pb"

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

func (l *UpdateLogic) Update(in *pb.OrderInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.OrderModel.Update(l.ctx, &order.Order{
		Id:                in.Id,
		EnterpriseId:      in.EnterpriseId,
		WorkshopId:        in.WorkshopId,
		ClientId:          in.ClientId,
		ProductionPlanId:  in.ProductionPlanId,
		State:             in.State,
		Address:           in.Address,
		Linkman:           in.Linkman,
		PhoneNumber:       in.PhoneNumber,
		Email:             in.Email,
		CorrespondingCode: in.CorrespondingCode,
		Remark:            in.Remark,
	})
}
