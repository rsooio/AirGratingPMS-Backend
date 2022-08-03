package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/workshop/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertWorkshopLogic {
	return &InsertWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertWorkshopLogic) InsertWorkshop(req *types.InsertWorkshopReq) (resp *types.InsertWorkshopReply, err error) {
	info, err := l.svcCtx.WorkshopRpc.Insert(l.ctx, &pb.WorkshopInfo{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		ManagerId:    req.ManagerId,
		Name:         req.Name,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		Remark:       req.Remark,
	})

	return &types.InsertWorkshopReply{
		Id:      info.Id,
		Message: "OK",
	}, err
}
