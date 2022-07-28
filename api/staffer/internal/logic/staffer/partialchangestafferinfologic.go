package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartialChangeStafferInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartialChangeStafferInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartialChangeStafferInfoLogic {
	return &PartialChangeStafferInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartialChangeStafferInfoLogic) PartialChangeStafferInfo(req *types.PartialChangeStafferInfoReq) (resp *types.PartialChangeStafferInfoReply, err error) {
	workshopId := req.WorkshopId
	if utils.GetRole(l.ctx) != "boss" {
		workshopId = utils.GetWorkshopId(l.ctx)
	}

	_, err = l.svcCtx.StafferRPC.PartialUpdate(l.ctx, &pb.StafferInfo{
		Id:          req.Id,
		WorkshopId:  workshopId,
		Username:    req.Username,
		Role:        req.Role,
		Name:        req.Email,
		Gender:      req.Gender,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Address:     req.Address,
		Remark:      req.Remark,
	})

	return &types.PartialChangeStafferInfoReply{
		Message: "OK",
	}, err
}
