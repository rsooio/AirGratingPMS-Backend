package logic

import (
	"context"

	"air-grating-pms-backend/rpc/order/internal/svc"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindListByEnterpriseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindListByEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindListByEnterpriseLogic {
	return &FindListByEnterpriseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindListByEnterpriseLogic) FindListByEnterprise(in *pb.FindListByEnterpriseReq) (*pb.OrderList, error) {
	list, count, err := l.svcCtx.OrderModel.FindListByEnterprise(l.ctx, in.GetEnterpriseId(), (in.GetPageNumber()-1)*in.GetPageSize(), in.GetPageSize())

	return &pb.OrderList{
		Count: count,
		List:  *list.RpcList(),
	}, err
}
