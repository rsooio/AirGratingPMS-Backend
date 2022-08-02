package product

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertProductLogic {
	return &InsertProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertProductLogic) InsertProduct(req *types.InsertProductReq) (resp *types.InsertProductReply, err error) {
	info, err := l.svcCtx.ProductRpc.Insert(l.ctx, &pb.ProductInfo{
		ProductSetId: req.ProductSetId,
		TechnologyId: req.TechnologyId,
		Length:       req.Length,
		Width:        req.Width,
		UnitPrice:    req.UnitPrice,
		Quantity:     req.Quantity,
		Remark:       req.Remark,
	})

	return &types.InsertProductReply{
		Id:      info.Id,
		Message: "OK",
	}, err
}
