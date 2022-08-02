package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStafferInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStafferInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStafferInfoLogic {
	return &UpdateStafferInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStafferInfoLogic) UpdateStafferInfo(req *types.UpdateStafferInfoReq) (resp *types.UpdateStafferInfoReply, err error) {
	_, err = l.svcCtx.StafferRPC.Update(l.ctx, &pb.StafferInfo{
		Id:          req.Id,
		WorkshopId:  utils.SelectWorkshopId(l.ctx, req.WorkshopId),
		Username:    req.Username,
		Role:        req.Role,
		Name:        req.Name,
		Gender:      req.Gender,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Address:     req.Address,
		Remark:      req.Remark,
	})

	return &types.UpdateStafferInfoReply{
		Message:       "OK",
		ChangedFields: []string{},
	}, err
}
