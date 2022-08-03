package logic

import (
	"context"

	"air-grating-pms-backend/rpc/enterprise/internal/svc"
	"air-grating-pms-backend/rpc/enterprise/pb"

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

func (l *FindOneByNameLogic) FindOneByName(in *pb.FindOneByNameReq) (*pb.EnterpriseInfo, error) {
	info, err := l.svcCtx.EnterpriseModel.FindOneByName(l.ctx, in.GetUsername())
	if err != nil {
		return nil, err
	}

	return &pb.EnterpriseInfo{
		Id:      info.Id,
		Name:    info.Name,
		Address: info.Address,
		Remark:  info.Remark,
		Version: info.Version,
	}, nil
}
