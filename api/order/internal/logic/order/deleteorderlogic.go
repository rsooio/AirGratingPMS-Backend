package order

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOrderLogic) DeleteOrder(req *types.DeleteOrderReq) (resp *types.DeleteOrderReply, err error) {
	_, err = l.svcCtx.OrderRpc.Delete(l.ctx, &pb.DeleteReq{
		Id: req.Id,
	})

	return &types.DeleteOrderReply{
		Message: "OK",
	}, err
}
