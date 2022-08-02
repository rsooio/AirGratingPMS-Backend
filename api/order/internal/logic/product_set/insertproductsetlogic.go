package product_set

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/product_set/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertProductSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertProductSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertProductSetLogic {
	return &InsertProductSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertProductSetLogic) InsertProductSet(req *types.InsertProductSetReq) (resp *types.InsertProductSetReply, err error) {
	info, err := l.svcCtx.ProductSetRpc.Insert(l.ctx, &pb.ProductSetInfo{
		OrderId: req.OrderId,
		Remark:  req.Remark,
	})

	return &types.InsertProductSetReply{
		Id:      info.Id,
		Message: "OK",
	}, err
}
