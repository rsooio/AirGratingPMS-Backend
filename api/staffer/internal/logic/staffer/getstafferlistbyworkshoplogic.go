package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStafferListByWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStafferListByWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStafferListByWorkshopLogic {
	return &GetStafferListByWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStafferListByWorkshopLogic) GetStafferListByWorkshop(req *types.GetStafferListByWorkshopReq) (resp *types.GetStafferListByWorkshopReply, err error) {
	list, err := l.svcCtx.StafferRPC.FindListByWorkshop(l.ctx, &pb.FindListByWorkshopReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		WorkshopId:   utils.GetWorkshopId(l.ctx),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetStafferListByWorkshopReply{
		Message:     "OK",
		Count:       list.Count,
		StafferList: *list.ApiList(),
	}, err
}
