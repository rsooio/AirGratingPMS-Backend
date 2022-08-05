package logic

import (
	"context"

	"air-grating-pms-backend/model/production_plan"
	"air-grating-pms-backend/rpc/production_plan/internal/svc"
	"air-grating-pms-backend/rpc/production_plan/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLogic) Insert(in *pb.ProductionPlanInfo) (*pb.InsertResp, error) {
	result, err := l.svcCtx.ProductionPlanModel.Insert(l.ctx, &production_plan.ProductionPlan{
		EnterpriseId: in.EnterpriseId,
		WorkshopId:   in.WorkshopId,
		State:        in.State,
		Remark:       in.Remark,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return &pb.InsertResp{
		Id: id,
	}, err
}
