package product_set

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/product_set/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSetLogic {
	return &DeleteProductSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductSetLogic) DeleteProductSet(req *types.DeleteProductSetReq) (resp *types.DeleteProductSetReply, err error) {
	_, err = l.svcCtx.ProductSetRpc.Delete(l.ctx, &pb.DeleteReq{
		Id: req.Id,
	})

	return &types.DeleteProductSetReply{
		Message: "OK",
	}, err
}
