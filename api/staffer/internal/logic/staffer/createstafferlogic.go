package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStafferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStafferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStafferLogic {
	return &CreateStafferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStafferLogic) CreateStaffer(req *types.CreateStafferReq) (resp *types.CreateReply, err error) {
	// todo: add your logic here and delete this line

	return
}
