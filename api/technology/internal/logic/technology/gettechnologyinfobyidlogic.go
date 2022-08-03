package technology

import (
	"context"

	"air-grating-pms-backend/api/technology/internal/svc"
	"air-grating-pms-backend/api/technology/internal/types"
	"air-grating-pms-backend/rpc/technology/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTechnologyInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTechnologyInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTechnologyInfoByIdLogic {
	return &GetTechnologyInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTechnologyInfoByIdLogic) GetTechnologyInfoById(req *types.GetTechnologyInfoByIdReq) (resp *types.TechnologyInfo, err error) {
	info, err := l.svcCtx.TechnologyRpc.FindOneById(l.ctx, &pb.FindOneByIdReq{
		Id: req.Id,
	})

	return info.Api(), err
}
