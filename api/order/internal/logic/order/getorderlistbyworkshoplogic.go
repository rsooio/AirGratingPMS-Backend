package order

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/order/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListByWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderListByWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListByWorkshopLogic {
	return &GetOrderListByWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListByWorkshopLogic) GetOrderListByWorkshop(req *types.GetOrderListByWorkshopReq) (resp *types.GetOrderListByWorkshopReply, err error) {
	workshopId := req.WorkshopId
	if utils.GetRole(l.ctx) != "boss" {
		workshopId = utils.GetWorkshopId(l.ctx)
	}

	list, err := l.svcCtx.OrderRpc.FindListByWorkshop(l.ctx, &pb.FindListByWorkshopReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		WorkshopId:   workshopId,
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetOrderListByWorkshopReply{
		Message:   "OK",
		Count:     list.Count,
		OrderList: *list.ApiList(),
	}, err
}
