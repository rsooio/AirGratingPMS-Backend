package enterprise

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/enterprise/pb"

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
	info, err := l.svcCtx.EnterpriseRpc.FindOneByName(l.ctx, &pb.FindOneByNameReq{
		Username: req.Name,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetEnterpriseIdByNameReply{
		Id: info.Id,
	}, nil
}
