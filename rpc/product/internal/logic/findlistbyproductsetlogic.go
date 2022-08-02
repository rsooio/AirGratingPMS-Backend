package logic

import (
	"context"

	"air-grating-pms-backend/rpc/product/internal/svc"
	"air-grating-pms-backend/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindListByProductSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindListByProductSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindListByProductSetLogic {
	return &FindListByProductSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindListByProductSetLogic) FindListByProductSet(in *pb.FindListByProductSetReq) (*pb.ProductList, error) {
	list, count, err := l.svcCtx.ProductModel.FindListByProductSetId(l.ctx, in.GetProductSetId(), (in.GetPageNumber()-1)*in.GetPageSize(), in.GetPageNumber())

	return &pb.ProductList{
		Count: count,
		List:  *list.RpcList(),
	}, err
}
