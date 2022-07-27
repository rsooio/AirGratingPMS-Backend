package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
