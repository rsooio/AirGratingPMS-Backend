package logic

import (
	"context"

	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/rpc/staffer/utils"

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

func (l *FindListByEnterpriseLogic) FindListByEnterprise(in *pb.FindListByEnterpriseReq) (*pb.StafferList, error) {
	list, err := l.svcCtx.StafferModel.FindListByEnterprise(l.ctx, in.GetEnterpriseId(), (in.GetPageNumber()-1)*in.GetPageSize(), in.GetPageSize())
	if err != nil {
		return nil, err
	}

	resp, err := utils.Convert(list)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
