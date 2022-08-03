package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/workshop/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWorkshopInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateWorkshopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWorkshopInfoLogic {
	return &UpdateWorkshopInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWorkshopInfoLogic) UpdateWorkshopInfo(req *types.UpdateWorkshopInfoReq) (resp *types.UpdateWorkshopInfoReply, err error) {
	_, err = l.svcCtx.WorkshopRpc.Update(l.ctx, &pb.WorkshopInfo{
		Id:           req.Id,
		EnterpriseId: l.ctx.Value("ent").(int64),
		ManagerId:    req.ManagerId,
		Name:         req.Name,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		Remark:       req.Remark,
	})

	return &types.UpdateWorkshopInfoReply{
		Message:       "OK",
		ChangedFields: []string{},
	}, err
}
