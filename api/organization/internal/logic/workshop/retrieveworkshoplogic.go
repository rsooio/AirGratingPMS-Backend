package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveWorkshopLogic {
	return &RetrieveWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveWorkshopLogic) RetrieveWorkshop(req *types.RetrieveWorkshopReq) (resp *types.RetrieveWorkshopReply, err error) {
	// todo: add your logic here and delete this line

	return
}
