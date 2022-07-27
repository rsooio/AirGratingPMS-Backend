package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeStafferPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeStafferPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeStafferPasswordLogic {
	return &ChangeStafferPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeStafferPasswordLogic) ChangeStafferPassword(req *types.ChangeStafferPasswordReq) (resp *types.ChangeStafferPasswordReply, err error) {
	// todo: add your logic here and delete this line

	return
}
