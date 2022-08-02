package logic

import (
	"context"

	"air-grating-pms-backend/model/product_set"
	"air-grating-pms-backend/rpc/product_set/internal/svc"
	"air-grating-pms-backend/rpc/product_set/pb"

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

func (l *InsertLogic) Insert(in *pb.ProductSetInfo) (*pb.InsertResp, error) {
	result, err := l.svcCtx.ProductSetModel.Insert(l.ctx, &product_set.ProductSet{
		OrderId: in.OrderId,
		Remark:  in.Remark,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return &pb.InsertResp{Id: id}, err
}
