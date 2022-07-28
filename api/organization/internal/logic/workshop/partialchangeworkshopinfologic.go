package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/workshop/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartialChangeWorkshopInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartialChangeWorkshopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartialChangeWorkshopInfoLogic {
	return &PartialChangeWorkshopInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartialChangeWorkshopInfoLogic) PartialChangeWorkshopInfo(req *types.PartialChangeWorkshopInfoReq) (resp *types.PartialChangeWorkshopInfoReply, err error) {
	_, err = l.svcCtx.WorkshopRPC.PartialUpdate(l.ctx, &pb.WorkshopInfo{
		Id:           req.Id,
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		ManagerId:    req.ManagerId,
		Name:         req.Name,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		Remark:       req.Remark,
	})

	return &types.PartialChangeWorkshopInfoReply{
		Message: "OK",
		// todo
		ChangedFields: []string{},
	}, err
}
