package logic

import (
	"context"

	"air-grating-pms-backend/model/order"
	"air-grating-pms-backend/rpc/order/internal/svc"
	"air-grating-pms-backend/rpc/order/pb"

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

func (l *InsertLogic) Insert(in *pb.OrderInfo) (*pb.InsertResp, error) {
	result, err := l.svcCtx.OrderModel.Insert(l.ctx, &order.Order{
		EnterpriseId:      in.EnterpriseId,
		WorkshopId:        in.WorkshopId,
		ClientId:          in.ClientId,
		Address:           in.Address,
		Linkman:           in.Linkman,
		PhoneNumber:       in.PhoneNumber,
		Email:             in.Email,
		CorrespondingCode: in.CorrespondingCode,
		Remark:            in.Remark,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return &pb.InsertResp{Id: id}, err
}
