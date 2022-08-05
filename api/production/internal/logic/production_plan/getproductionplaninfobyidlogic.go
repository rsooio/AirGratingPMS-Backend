package production_plan

import (
	"context"

	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/rpc/production_plan/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductionPlanInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductionPlanInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductionPlanInfoByIdLogic {
	return &GetProductionPlanInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductionPlanInfoByIdLogic) GetProductionPlanInfoById(req *types.GetProductionPlanInfoByIdReq) (resp *types.ProductionPlanInfo, err error) {
	info, err := l.svcCtx.ProductionPlanRpc.FindOneById(l.ctx, &pb.FindOneByIdReq{
		Id: req.Id,
	})

	return info.Api(), err
}
