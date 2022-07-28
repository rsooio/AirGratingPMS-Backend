package logic

import (
	"context"

	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/rpc/staffer/utils"

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
	list, err := l.svcCtx.StafferModel.FindListByWorkshop(l.ctx, in.GetEnterpriseId(), in.GetWorkshopId(), (in.GetPageNumber()-1)*in.GetPageSize(), in.GetPageSize())
	if err != nil {
		return nil, err
	}

	resp, err := utils.ListConvert(list)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
