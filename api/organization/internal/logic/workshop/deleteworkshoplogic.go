package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/workshop/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWorkshopLogic {
	return &DeleteWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteWorkshopLogic) DeleteWorkshop(req *types.DeleteWorkshopReq) (resp *types.DeleteWorkshopReply, err error) {
	_, err = l.svcCtx.WorkshopRPC.Delete(l.ctx, &pb.DeleteReq{Id: req.Id})

	return &types.DeleteWorkshopReply{
		Message: "OK",
	}, err
}
