package production_plan

import (
	"context"

	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/rpc/production_plan/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductionPlanListByWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductionPlanListByWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductionPlanListByWorkshopLogic {
	return &GetProductionPlanListByWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductionPlanListByWorkshopLogic) GetProductionPlanListByWorkshop(req *types.GetProductionPlanListByWorkshopReq) (resp *types.GetProductionPlanListByWorkshopReply, err error) {
	list, err := l.svcCtx.ProductionPlanRpc.FindListByWorkshop(l.ctx, &pb.FindListByWorkshopReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		WorkshopId:   utils.SelectWorkshopId(l.ctx, req.WorkshopId),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetProductionPlanListByWorkshopReply{
		Message: "OK",
		Count:   list.Count,
		List:    *list.ApiList(),
	}, err
}
