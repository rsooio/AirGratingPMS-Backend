package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	convert "air-grating-pms-backend/api/organization/utils"
	"air-grating-pms-backend/rpc/workshop/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveWorkshopLogic {
	return &RetrieveWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveWorkshopLogic) RetrieveWorkshop(req *types.RetrieveWorkshopReq) (resp *types.RetrieveWorkshopReply, err error) {
	list, err := l.svcCtx.WorkshopRPC.FindListByEnterprise(l.ctx, &pb.FindListByEnterpriseReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		PageSize:     int32(req.PageSize),
		PageNumber:   int32(req.PageNumber),
	})
	if err != nil {
		return nil, err
	}

	listresp, err := convert.Convert(list)

	return &types.RetrieveWorkshopReply{
		Message:      "",
		WorkshopList: *listresp,
	}, err
}
