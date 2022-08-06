package client

import (
	"context"

	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"
	"air-grating-pms-backend/rpc/client/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClientListByWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClientListByWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClientListByWorkshopLogic {
	return &GetClientListByWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClientListByWorkshopLogic) GetClientListByWorkshop(req *types.GetClientListByWorkshopReq) (resp *types.GetClientListByWorkshopReply, err error) {
	list, err := l.svcCtx.ClientRpc.FindListByWorkshop(l.ctx, &pb.FindListByWorkshopReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		WorkshopId:   utils.SelectWorkshopId(l.ctx, req.WorkshopId),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetClientListByWorkshopReply{
		Count:   list.Count,
		List:    *list.ApiList(),
		Message: "OK",
	}, err
}
