package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetStafferPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetStafferPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetStafferPasswordLogic {
	return &ResetStafferPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetStafferPasswordLogic) ResetStafferPassword(req *types.ResetStafferPasswordReq) (resp *types.ResetStafferPasswordReply, err error) {
	// todo: add your logic here and delete this line

	return
}
