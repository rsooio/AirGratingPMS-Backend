package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeStafferInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeStafferInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeStafferInfoLogic {
	return &ChangeStafferInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeStafferInfoLogic) ChangeStafferInfo(req *types.ChangeStafferInfoReq) (resp *types.ChangeStafferInfoReply, err error) {
	workshopId := req.WorkshopId
	if utils.GetRole(l.ctx) != "boss" {
		workshopId = utils.GetWorkshopId(l.ctx)
	}

	_, err = l.svcCtx.StafferRPC.CustomUpdate(l.ctx, &pb.StafferInfo{
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

	return &types.ChangeStafferInfoReply{
		Message: "OK",
	}, err
}
