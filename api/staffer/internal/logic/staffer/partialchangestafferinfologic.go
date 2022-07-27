package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartialChangeStafferInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartialChangeStafferInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartialChangeStafferInfoLogic {
	return &PartialChangeStafferInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartialChangeStafferInfoLogic) PartialChangeStafferInfo(req *types.PartialChangeStafferInfoReq) (resp *types.PartialChangeStafferInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
