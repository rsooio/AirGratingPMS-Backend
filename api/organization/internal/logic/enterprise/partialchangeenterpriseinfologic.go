package enterprise

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartialChangeEnterpriseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartialChangeEnterpriseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartialChangeEnterpriseInfoLogic {
	return &PartialChangeEnterpriseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartialChangeEnterpriseInfoLogic) PartialChangeEnterpriseInfo(req *types.PartialChangeEnterpriseInfoReq) (resp *types.PartialChangeEnterpriseInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
