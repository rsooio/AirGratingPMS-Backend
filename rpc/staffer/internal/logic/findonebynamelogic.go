package logic

import (
	"context"

	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOneByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneByNameLogic {
	return &FindOneByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneByNameLogic) FindOneByName(in *pb.FindOneByNameReq) (*pb.StafferInfo, error) {
	info, err := l.svcCtx.StafferModel.FindOneByEnterpriseIdUsername(l.ctx, in.GetEnterpriseId(), in.GetUsername())

	return info.Rpc(), err
}
