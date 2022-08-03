package logic

import (
	"context"

	"air-grating-pms-backend/rpc/technology/internal/svc"
	"air-grating-pms-backend/rpc/technology/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOneByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneByIdLogic {
	return &FindOneByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneByIdLogic) FindOneById(in *pb.FindOneByIdReq) (*pb.TechnologyInfo, error) {
	info, err := l.svcCtx.TechnologyModel.FindOne(l.ctx, in.Id)

	return info.Rpc(), err
}
