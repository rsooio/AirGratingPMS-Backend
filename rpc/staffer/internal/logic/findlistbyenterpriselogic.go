package logic

import (
	"context"

	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"

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
	list, count, err := l.svcCtx.StafferModel.FindListByEnterprise(l.ctx, in.GetEnterpriseId(), (in.GetPageNumber()-1)*in.GetPageSize(), in.GetPageSize())
	if err != nil {
		return nil, err
	}

	return &pb.StafferList{
		Count: count,
		List:  *list.RpcList(),
	}, nil
}
