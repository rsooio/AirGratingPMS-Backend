package logic

import (
	"context"

	"air-grating-pms-backend/model/technology"
	"air-grating-pms-backend/rpc/technology/internal/svc"
	"air-grating-pms-backend/rpc/technology/pb"

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

func (l *UpdateLogic) Update(in *pb.TechnologyInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.TechnologyModel.Update(l.ctx, &technology.Technology{
		Id:           in.Id,
		EnterpriseId: in.EnterpriseId,
		WorkshopId:   in.WorkshopId,
		Name:         in.Name,
		Info:         in.Info,
		Remark:       in.Remark,
	})
}
