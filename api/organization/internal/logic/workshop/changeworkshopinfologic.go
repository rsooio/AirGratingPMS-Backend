package workshop

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/workshop/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeWorkshopInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeWorkshopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeWorkshopInfoLogic {
	return &ChangeWorkshopInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeWorkshopInfoLogic) ChangeWorkshopInfo(req *types.ChangeWorkshopInfoReq) (resp *types.ChangeWorkshopInfoReply, err error) {
	_, err = l.svcCtx.WorkshopRPC.Update(l.ctx, &pb.WorkshopInfo{
		Id:           req.Id,
		EnterpriseId: l.ctx.Value("ent").(int64),
		ManagerId:    req.ManagerId,
		Name:         req.Name,
		Address:      req.Address,
		PhoneNumber:  req.PhoneNumber,
		Remark:       req.Remark,
	})

	return &types.ChangeWorkshopInfoReply{
		Message: "OK",
	}, err
}
