package order

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderInfoLogic {
	return &UpdateOrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderInfoLogic) UpdateOrderInfo(req *types.UpdateOrderInfoReq) (resp *types.UpdateOrderInfoReply, err error) {
	_, err = l.svcCtx.OrderRpc.Update(l.ctx, &pb.OrderInfo{
		Id:                req.Id,
		ClientId:          req.ClientId,
		Address:           req.Address,
		Linkman:           req.Linkman,
		PhoneNumber:       req.PhoneNumber,
		Email:             req.Email,
		CorrespondingCode: req.CorrespondingCode,
		Remark:            req.Remark,
	})

	return &types.UpdateOrderInfoReply{
		Message:       "OK",
		ChangedFields: []string{},
	}, err
}
