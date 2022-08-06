package client

import (
	"context"

	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"
	"air-grating-pms-backend/rpc/client/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCientLogic {
	return &DeleteCientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCientLogic) DeleteCient(req *types.DeleteCientReq) (resp *types.DeleteCientReply, err error) {
	_, err = l.svcCtx.ClientRpc.Delete(l.ctx, &pb.DeleteReq{
		Id: req.Id,
	})

	return &types.DeleteCientReply{
		Message: "OK",
	}, err
}
