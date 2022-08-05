package logic

import (
	"context"

	"air-grating-pms-backend/rpc/production_plan/internal/svc"
	"air-grating-pms-backend/rpc/production_plan/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindListByEnterpriseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindListByEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindListByEnterpriseLogic {
	return &FindListByEnterpriseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindListByEnterpriseLogic) FindListByEnterprise(in *pb.FindListByEnterpriseReq) (*pb.ProductionPlanList, error) {
	list, count, err := l.svcCtx.ProductionPlanModel.FindListByEnterprise(l.ctx, in.EnterpriseId, (in.PageNumber-1)*in.PageSize, in.PageSize)

	return &pb.ProductionPlanList{
		Count: count,
		List:  *list.RpcList(),
	}, err
}
