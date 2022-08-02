package logic

import (
	"context"

	"air-grating-pms-backend/model/product"
	"air-grating-pms-backend/rpc/product/internal/svc"
	"air-grating-pms-backend/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLogic) Insert(in *pb.ProductInfo) (*pb.InsertResp, error) {
	result, err := l.svcCtx.ProductModel.Insert(l.ctx, &product.Product{
		ProductSetId: in.ProductSetId,
		TechnologyId: in.TechnologyId,
		Length:       in.Length,
		Width:        in.Width,
		UnitPrice:    in.UnitPrice,
		Quantity:     in.Quantity,
		Remark:       in.Remark,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return &pb.InsertResp{Id: id}, err
}
