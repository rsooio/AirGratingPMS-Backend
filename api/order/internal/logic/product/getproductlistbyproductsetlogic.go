package product

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListByProductSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductListByProductSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListByProductSetLogic {
	return &GetProductListByProductSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductListByProductSetLogic) GetProductListByProductSet(req *types.GetProductListByProductSetReq) (resp *types.GetProductListByProductSetReply, err error) {
	list, err := l.svcCtx.ProductRpc.FindListByProductSet(l.ctx, &pb.FindListByProductSetReq{
		ProductSetId: req.ProductSetId,
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetProductListByProductSetReply{
		Message:     "OK",
		Count:       list.Count,
		ProductList: *list.ApiList(),
	}, err
}
