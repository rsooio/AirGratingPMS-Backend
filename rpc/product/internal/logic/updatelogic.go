package logic

import (
	"context"

	"air-grating-pms-backend/model/product"
	"air-grating-pms-backend/rpc/product/internal/svc"
	"air-grating-pms-backend/rpc/product/pb"

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

func (l *UpdateLogic) Update(in *pb.ProductInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.ProductModel.Update(l.ctx, &product.Product{
		Id:           in.Id,
		ProductSetId: in.ProductSetId,
		TechnologyId: in.TechnologyId,
		Length:       in.Length,
		Width:        in.Width,
		UnitPrice:    in.UnitPrice,
		Quantity:     in.Quantity,
		Remark:       in.Remark,
	})
}
