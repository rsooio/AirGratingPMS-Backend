package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/workshop/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWorkshopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateWorkshopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWorkshopLogic {
	return &CreateWorkshopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWorkshopLogic) CreateWorkshop(req *types.CreateWorkshopReq) (resp *types.CreateWorkshopReply, err error) {
	_, err = l.svcCtx.WorkshopRPC.Insert(l.ctx, &pb.WorkshopInfo{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		ManagerId:    req.ManagerId,
		Name:         req.Name,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		Remark:       req.Remark,
	})

	return &types.CreateWorkshopReply{
		Message: "OK",
	}, err
}
