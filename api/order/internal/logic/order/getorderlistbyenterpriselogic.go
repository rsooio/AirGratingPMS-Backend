package order

import (
	"context"

	"air-grating-pms-backend/api/order/internal/svc"
	"air-grating-pms-backend/api/order/internal/types"
	"air-grating-pms-backend/rpc/order/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListByEnterpriseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderListByEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListByEnterpriseLogic {
	return &GetOrderListByEnterpriseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListByEnterpriseLogic) GetOrderListByEnterprise(req *types.GetOrderListByEnterpriseReq) (resp *types.GetOrderListByEnterpriseReply, err error) {
	list, err := l.svcCtx.OrderRpc.FindListByEnterprise(l.ctx, &pb.FindListByEnterpriseReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetOrderListByEnterpriseReply{
		Message:   "OK",
		Count:     list.Count,
		OrderList: *list.ApiList(),
	}, err
}
