package logic

import (
	"context"

	"air-grating-pms-backend/model/product_set"
	"air-grating-pms-backend/rpc/product_set/internal/svc"
	"air-grating-pms-backend/rpc/product_set/pb"

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

func (l *UpdateLogic) Update(in *pb.ProductSetInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.ProductSetModel.Update(l.ctx, &product_set.ProductSet{
		Id:      in.Id,
		OrderId: in.OrderId,
		Remark:  in.Remark,
	})
}
