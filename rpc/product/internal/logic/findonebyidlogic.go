package logic

import (
	"context"

	"air-grating-pms-backend/rpc/product/internal/svc"
	"air-grating-pms-backend/rpc/product/pb"

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

func (l *FindOneByIdLogic) FindOneById(in *pb.FindOneByIdReq) (*pb.ProductInfo, error) {
	info, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.GetId())

	return info.Rpc(), err
}
