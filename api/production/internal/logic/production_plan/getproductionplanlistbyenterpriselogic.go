package production_plan

import (
	"context"

	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/rpc/production_plan/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductionPlanListByEnterpriseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductionPlanListByEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductionPlanListByEnterpriseLogic {
	return &GetProductionPlanListByEnterpriseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductionPlanListByEnterpriseLogic) GetProductionPlanListByEnterprise(req *types.GetProductionPlanListByEnterpriseReq) (resp *types.GetProductionPlanListByEnterpriseReply, err error) {
	list, err := l.svcCtx.ProductionPlanRpc.FindListByEnterprise(l.ctx, &pb.FindListByEnterpriseReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetProductionPlanListByEnterpriseReply{
		Message: "OK",
		Count:   list.Count,
		List:    *list.ApiList(),
	}, err
}
