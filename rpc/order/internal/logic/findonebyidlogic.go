package logic

import (
	"context"

	"air-grating-pms-backend/rpc/order/internal/svc"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOneByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneByIdLogic {
	return &FindOneByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneByIdLogic) FindOneById(in *pb.FindOneByIdReq) (*pb.OrderInfo, error) {
	info, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.GetId())

	return info.Rpc(), err
}
