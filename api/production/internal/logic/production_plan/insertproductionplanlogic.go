package production_plan

import (
	"context"

	"air-grating-pms-backend/api/production/internal/svc"
	"air-grating-pms-backend/api/production/internal/types"
	"air-grating-pms-backend/rpc/production_plan/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertProductionPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertProductionPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertProductionPlanLogic {
	return &InsertProductionPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertProductionPlanLogic) InsertProductionPlan(req *types.InsertProductionPlanReq) (resp *types.InsertProductionPlanReply, err error) {
	info, err := l.svcCtx.ProductionPlanRpc.Insert(l.ctx, &pb.ProductionPlanInfo{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		WorkshopId:   utils.SelectWorkshopId(l.ctx, req.WorkshopId),
		Remark:       req.Remark,
	})

	return &types.InsertProductionPlanReply{
		Id:      info.Id,
		Message: "OK",
	}, err
}
