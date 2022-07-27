package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeStafferInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeStafferInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeStafferInfoLogic {
	return &ChangeStafferInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeStafferInfoLogic) ChangeStafferInfo(req *types.ChangeStafferInfoReq) (resp *types.ChangeStafferInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
