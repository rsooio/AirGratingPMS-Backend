package product

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.UpdateProductReq) (resp *types.UpdateProductReply, err error) {
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &pb.ProductInfo{
		Id:           req.Id,
		ProductSetId: req.ProductSetId,
		TechnologyId: req.TechnologyId,
		Length:       req.Length,
		Width:        req.Width,
		UnitPrice:    req.UnitPrice,
		Quantity:     req.Quantity,
		Remark:       req.Remark,
	})

	return &types.UpdateProductReply{
		Message:       "OK",
		ChangedFields: []string{},
	}, err
}
