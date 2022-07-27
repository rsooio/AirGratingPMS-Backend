package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartialChangeWorkshopInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartialChangeWorkshopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartialChangeWorkshopInfoLogic {
	return &PartialChangeWorkshopInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartialChangeWorkshopInfoLogic) PartialChangeWorkshopInfo(req *types.PartialChangeWorkshopInfoReq) (resp *types.PartialChangeWorkshopInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
