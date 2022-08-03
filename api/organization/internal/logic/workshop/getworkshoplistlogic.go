package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/workshop/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkshopListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWorkshopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkshopListLogic {
	return &GetWorkshopListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWorkshopListLogic) GetWorkshopList(req *types.GetWorkshopListReq) (resp *types.GetWorkshopListReply, err error) {
	list, err := l.svcCtx.WorkshopRpc.FindListByEnterprise(l.ctx, &pb.FindListByEnterpriseReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})

	return &types.GetWorkshopListReply{
		Message:      "OK",
		WorkshopList: *list.ApiList(),
	}, err
}
