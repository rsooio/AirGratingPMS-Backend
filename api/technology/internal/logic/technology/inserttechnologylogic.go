package technology

import (
	"context"

	"air-grating-pms-backend/api/technology/internal/svc"
	"air-grating-pms-backend/api/technology/internal/types"
	"air-grating-pms-backend/rpc/technology/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertTechnologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertTechnologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertTechnologyLogic {
	return &InsertTechnologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertTechnologyLogic) InsertTechnology(req *types.InsertTechnologyReq) (resp *types.InsertTechnologyReply, err error) {
	info, err := l.svcCtx.TechnologyRpc.Insert(l.ctx, &pb.TechnologyInfo{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		WorkshopId:   utils.SelectWorkshopId(l.ctx, req.WorkshopId),
		Name:         req.Name,
		Info:         req.Info,
		Remark:       req.Remark,
	})

	return &types.InsertTechnologyReply{
		Id:      info.Id,
		Message: "OK",
	}, err
}
