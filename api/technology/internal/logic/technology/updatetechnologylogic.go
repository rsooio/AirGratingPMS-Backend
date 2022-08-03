package technology

import (
	"context"

	"air-grating-pms-backend/api/technology/internal/svc"
	"air-grating-pms-backend/api/technology/internal/types"
	"air-grating-pms-backend/rpc/technology/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTechnologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTechnologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTechnologyLogic {
	return &UpdateTechnologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTechnologyLogic) UpdateTechnology(req *types.UpdateTechnologyReq) (resp *types.UpdateTechnologyReply, err error) {
	_, err = l.svcCtx.TechnologyRpc.Update(l.ctx, &pb.TechnologyInfo{
		Id:           req.Id,
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		WorkshopId:   utils.SelectWorkshopId(l.ctx, req.WorkshopId),
		Name:         req.Name,
		Info:         req.Info,
		Remark:       req.Remark,
	})

	return &types.UpdateTechnologyReply{
		Message:       "OK",
		ChangedFields: []string{},
	}, err
}
