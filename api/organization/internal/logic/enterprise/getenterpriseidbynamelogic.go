package enterprise

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEnterpriseIdByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEnterpriseIdByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEnterpriseIdByNameLogic {
	return &GetEnterpriseIdByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEnterpriseIdByNameLogic) GetEnterpriseIdByName(req *types.GetEnterpriseIdByNameReq) (resp *types.GetEnterpriseIdByNameReply, err error) {
	// todo: add your logic here and delete this line

	return
}
