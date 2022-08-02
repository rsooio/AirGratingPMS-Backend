package logic

import (
	"context"

	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"

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

func (l *FindListByWorkshopLogic) FindListByWorkshop(in *pb.FindListByWorkshopReq) (*pb.StafferList, error) {
	list, count, err := l.svcCtx.StafferModel.FindListByWorkshop(l.ctx, in.GetEnterpriseId(), in.GetWorkshopId(), (in.GetPageNumber()-1)*in.GetPageSize(), in.GetPageSize())
	if err != nil {
		return nil, err
	}

	return &pb.StafferList{
		Count: count,
		List:  *list.RpcList(),
	}, nil
}
