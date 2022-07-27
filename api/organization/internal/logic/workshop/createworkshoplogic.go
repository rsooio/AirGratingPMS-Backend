package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWorkshopLogic {
	return &CreateWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWorkshopLogic) CreateWorkshop(req *types.CreateWorkshopReq) (resp *types.CreateWorkshopReply, err error) {
	// todo: add your logic here and delete this line

	return
}
