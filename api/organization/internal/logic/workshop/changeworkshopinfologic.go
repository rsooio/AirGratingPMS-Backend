package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeWorkshopInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeWorkshopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeWorkshopInfoLogic {
	return &ChangeWorkshopInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeWorkshopInfoLogic) ChangeWorkshopInfo(req *types.ChangeWorkshopInfoReq) (resp *types.ChangeWorkshopInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
