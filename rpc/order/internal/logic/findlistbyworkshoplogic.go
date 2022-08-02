package logic

import (
	"context"

	"air-grating-pms-backend/rpc/order/internal/svc"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindListByWorkshopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindListByWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindListByWorkshopLogic {
	return &FindListByWorkshopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindListByWorkshopLogic) FindListByWorkshop(in *pb.FindListByWorkshopReq) (*pb.OrderList, error) {
	list, count, err := l.svcCtx.OrderModel.FindListByWorkshop(l.ctx, in.GetEnterpriseId(), in.GetWorkshopId(), (in.GetPageNumber()-1)*in.GetPageSize(), in.GetPageSize())

	return &pb.OrderList{
		Count: count,
		List:  *list.RpcList(),
	}, err
}
