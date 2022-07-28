package enterprise

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/enterprise/pb"

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
	l.svcCtx.EnterpriseRPC.PartialUpdate(l.ctx, &pb.EnterpriseInfo{
		Name:    req.Name,
		Address: req.Address,
		Remark:  req.Remark,
	})

	return
}
