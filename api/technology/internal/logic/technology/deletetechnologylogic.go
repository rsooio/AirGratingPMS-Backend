package technology

import (
	"context"

	"air-grating-pms-backend/api/technology/internal/svc"
	"air-grating-pms-backend/api/technology/internal/types"
	"air-grating-pms-backend/rpc/technology/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTechnologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTechnologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTechnologyLogic {
	return &DeleteTechnologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTechnologyLogic) DeleteTechnology(req *types.DeleteTechnologyReq) (resp *types.DeleteTechnologyReply, err error) {
	_, err = l.svcCtx.TechnologyRpc.Delete(l.ctx, &pb.DeleteReq{
		Id: req.Id,
	})

	return &types.DeleteTechnologyReply{
		Message: "OK",
	}, err
}
