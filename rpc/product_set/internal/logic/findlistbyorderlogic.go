package logic

import (
	"context"

	"air-grating-pms-backend/rpc/product_set/internal/svc"
	"air-grating-pms-backend/rpc/product_set/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindListByOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindListByOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindListByOrderLogic {
	return &FindListByOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindListByOrderLogic) FindListByOrder(in *pb.FindListByOrderReq) (*pb.ProductSetList, error) {
	list, count, err := l.svcCtx.ProductSetModel.FindListByOrderId(l.ctx, in.GetOrderId(), (in.GetPageNumber()-1)*in.GetPageSize(), in.GetPageSize())

	return &pb.ProductSetList{
		Count: count,
		List:  *list.RpcList(),
	}, err
}
