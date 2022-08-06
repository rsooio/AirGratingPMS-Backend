package client

import (
	"context"

	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"
	"air-grating-pms-backend/rpc/client/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertClientLogic {
	return &InsertClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertClientLogic) InsertClient(req *types.InsertClientReq) (resp *types.InsertClientReply, err error) {
	info, err := l.svcCtx.ClientRpc.Insert(l.ctx, &pb.ClientInfo{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		WorkshopId:   utils.SelectWorkshopId(l.ctx, req.WorkshopId),
		Name:         req.Name,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		Address:      req.Address,
		Remark:       req.Remark,
	})

	return &types.InsertClientReply{
		Id:      info.Id,
		Message: "OK",
	}, err
}
