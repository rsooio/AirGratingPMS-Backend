package client

import (
	"context"

	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"
	"air-grating-pms-backend/rpc/client/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClientLogic {
	return &UpdateClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateClientLogic) UpdateClient(req *types.UpdateClientReq) (resp *types.UpdateClientReply, err error) {
	_, err = l.svcCtx.ClientRpc.Update(l.ctx, &pb.ClientInfo{
		Id:          req.Id,
		WorkshopId:  utils.ZeroSelectWorkshopId(l.ctx, req.WorkshopId),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Address:     req.Address,
		Remark:      req.Remark,
	})

	return &types.UpdateClientReply{
		Message: "OK",
	}, err
}
