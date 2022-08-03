package technology

import (
	"context"

	"air-grating-pms-backend/api/technology/internal/svc"
	"air-grating-pms-backend/api/technology/internal/types"
	"air-grating-pms-backend/rpc/technology/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTechnologyListByWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTechnologyListByWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTechnologyListByWorkshopLogic {
	return &GetTechnologyListByWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTechnologyListByWorkshopLogic) GetTechnologyListByWorkshop(req *types.GetTechnologyListByWorkshopReq) (resp *types.GetTechnologyListByWorkshopReply, err error) {
	list, err := l.svcCtx.TechnologyRpc.FindListByWorkshop(l.ctx, &pb.FindListByWorkshopReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		WorkshopId:   utils.SelectWorkshopId(l.ctx, req.WorkshopId),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetTechnologyListByWorkshopReply{
		Message:        "OK",
		Count:          list.Count,
		TechnologyList: *list.ApiList(),
	}, err
}
