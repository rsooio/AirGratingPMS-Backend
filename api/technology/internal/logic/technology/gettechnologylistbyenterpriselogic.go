package technology

import (
	"context"

	"air-grating-pms-backend/api/technology/internal/svc"
	"air-grating-pms-backend/api/technology/internal/types"
	"air-grating-pms-backend/rpc/technology/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTechnologyListByEnterpriseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTechnologyListByEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTechnologyListByEnterpriseLogic {
	return &GetTechnologyListByEnterpriseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTechnologyListByEnterpriseLogic) GetTechnologyListByEnterprise(req *types.GetTechnologyListByEnterpriseReq) (resp *types.GetTechnologyListByEnterpriseReply, err error) {
	list, err := l.svcCtx.TechnologyRpc.FindListByEnterprise(l.ctx, &pb.FindListByEnterpriseReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetTechnologyListByEnterpriseReply{
		Message:        "OK",
		Count:          list.Count,
		TechnologyList: *list.ApiList(),
	}, err
}
