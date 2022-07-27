package enterprise

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeEnterpriseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeEnterpriseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeEnterpriseInfoLogic {
	return &ChangeEnterpriseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeEnterpriseInfoLogic) ChangeEnterpriseInfo(req *types.ChangeEnterpriseInfoReq) (resp *types.ChangeEnterpriseInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
