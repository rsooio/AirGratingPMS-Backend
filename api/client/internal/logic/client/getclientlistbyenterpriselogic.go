package client

import (
	"context"

	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"
	"air-grating-pms-backend/rpc/client/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClientListByEnterpriseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClientListByEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClientListByEnterpriseLogic {
	return &GetClientListByEnterpriseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClientListByEnterpriseLogic) GetClientListByEnterprise(req *types.GetClientListByEnterpriseReq) (resp *types.GetClientListByEnterpriseReply, err error) {
	list, err := l.svcCtx.ClientRpc.FindListByEnterprise(l.ctx, &pb.FindListByEnterpriseReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetClientListByEnterpriseReply{
		Count:   list.Count,
		List:    *list.ApiList(),
		Message: "OK",
	}, err
}
