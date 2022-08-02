package product_set

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/product_set/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSetListByOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductSetListByOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSetListByOrderLogic {
	return &GetProductSetListByOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductSetListByOrderLogic) GetProductSetListByOrder(req *types.GetProductSetListByOrderReq) (resp *types.GetProductSetListByOrderReply, err error) {
	list, err := l.svcCtx.ProductSetRpc.FindListByOrder(l.ctx, &pb.FindListByOrderReq{
		OrderId:    req.OrderId,
		PageSize:   req.PageSize,
		PageNumber: req.PageNumber,
	})

	return &types.GetProductSetListByOrderReply{
		Message:        "OK",
		Count:          list.Count,
		ProductSetList: *list.ApiList(),
	}, err
}
