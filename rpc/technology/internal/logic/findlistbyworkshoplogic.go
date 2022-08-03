package logic

import (
	"context"

	"air-grating-pms-backend/rpc/technology/internal/svc"
	"air-grating-pms-backend/rpc/technology/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindListByWorkshopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindListByWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindListByWorkshopLogic {
	return &FindListByWorkshopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindListByWorkshopLogic) FindListByWorkshop(in *pb.FindListByWorkshopReq) (*pb.TechnologyList, error) {
	list, count, err := l.svcCtx.TechnologyModel.FindListByWorkshop(l.ctx, in.EnterpriseId, in.WorkshopId, (in.PageNumber-1)*in.PageSize, in.PageSize)

	return &pb.TechnologyList{
		Count: count,
		List:  *list.RpcList(),
	}, err
}
