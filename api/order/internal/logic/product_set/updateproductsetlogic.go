package product_set

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/product_set/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSetLogic {
	return &UpdateProductSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductSetLogic) UpdateProductSet(req *types.UpdateProductSetReq) (resp *types.UpdateProductSetReply, err error) {
	_, err = l.svcCtx.ProductSetRpc.Update(l.ctx, &pb.ProductSetInfo{
		Id:     req.Id,
		Remark: req.Remark,
	})

	return &types.UpdateProductSetReply{
		Message:       "OK",
		ChangedFields: []string{},
	}, err
}
