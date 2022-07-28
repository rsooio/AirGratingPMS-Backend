package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/api/staffer/utils/convert"
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
	if err != nil {
		return nil, err
	}

	listResp, err := convert.ListConvert(list)

	return &types.GetStafferListByWorkshopReply{
		Message:     "OK",
		StafferList: *listResp,
	}, err
}
