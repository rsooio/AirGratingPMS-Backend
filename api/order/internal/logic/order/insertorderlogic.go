package order

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/order/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertOrderLogic {
	return &InsertOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertOrderLogic) InsertOrder(req *types.InsertOrderReq) (resp *types.InsertOrderReply, err error) {
	workshopId := req.WorkshopId
	if utils.GetRole(l.ctx) != "boss" {
		workshopId = utils.GetWorkshopId(l.ctx)
	}

	info, err := l.svcCtx.OrderRpc.Insert(l.ctx, &pb.OrderInfo{
		EnterpriseId:      utils.GetEnterpriseId(l.ctx),
		WorkshopId:        workshopId,
		ClientId:          req.ClientId,
		Address:           req.Address,
		Linkman:           req.Linkman,
		PhoneNumber:       req.PhoneNumber,
		Email:             req.Email,
		CorrespondingCode: req.CorrespondingCode,
		Remark:            req.Remark,
	})

	return &types.InsertOrderReply{
		Id:      info.Id,
		Message: "OK",
	}, err
}
