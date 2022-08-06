package logic

import (
	"context"

	"air-grating-pms-backend/rpc/client/internal/svc"
	"air-grating-pms-backend/rpc/client/pb"

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

func (l *FindListByEnterpriseLogic) FindListByEnterprise(in *pb.FindListByEnterpriseReq) (*pb.ClientList, error) {
	list, count, err := l.svcCtx.ClientModel.FindListByEnterprise(l.ctx, in.EnterpriseId, (in.PageNumber-1)*in.PageSize, in.PageSize)

	return &pb.ClientList{
		Count: count,
		List:  *list.RpcList(),
	}, err
}
