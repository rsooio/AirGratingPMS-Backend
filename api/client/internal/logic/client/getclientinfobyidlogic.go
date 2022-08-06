package client

import (
	"context"

	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"
	"air-grating-pms-backend/rpc/client/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClientInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClientInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClientInfoByIdLogic {
	return &GetClientInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClientInfoByIdLogic) GetClientInfoById(req *types.GetClientInfoByIdReq) (resp *types.ClientInfo, err error) {
	info, err := l.svcCtx.ClientRpc.FindOneById(l.ctx, &pb.FindOneByIdReq{
		Id: req.Id,
	})

	return info.Api(), err
}
