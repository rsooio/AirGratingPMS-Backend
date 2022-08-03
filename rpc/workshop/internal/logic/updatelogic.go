package logic

import (
	"context"

	"air-grating-pms-backend/model/workshop"
	"air-grating-pms-backend/rpc/workshop/internal/svc"
	"air-grating-pms-backend/rpc/workshop/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *pb.WorkshopInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.WorkshopModel.Update(l.ctx, &workshop.Workshop{
		Id:           in.Id,
		EnterpriseId: in.EnterpriseId,
		Name:         in.Name,
		Address:      in.Address,
		PhoneNumber:  in.PhoneNumber,
		ManagerId:    in.ManagerId,
		Remark:       in.Remark,
		Version:      in.Version,
	})
}
