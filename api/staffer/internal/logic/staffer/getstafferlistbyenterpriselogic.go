package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStafferListByEnterpriseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStafferListByEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStafferListByEnterpriseLogic {
	return &GetStafferListByEnterpriseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStafferListByEnterpriseLogic) GetStafferListByEnterprise(req *types.GetStafferListByEnterpriseReq) (resp *types.GetStafferListByEnterpriseReply, err error) {
	list, err := l.svcCtx.StafferRPC.FindListByEnterprise(l.ctx, &pb.FindListByEnterpriseReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetStafferListByEnterpriseReply{
		Message:     "OK",
		Count:       list.Count,
		StafferList: *list.ApiList(),
	}, err
}
