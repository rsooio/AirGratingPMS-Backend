package logic

import (
	"context"

	"air-grating-pms-backend/model/client"
	"air-grating-pms-backend/rpc/client/internal/svc"
	"air-grating-pms-backend/rpc/client/pb"

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

func (l *UpdateLogic) Update(in *pb.ClientInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.ClientModel.Update(l.ctx, &client.Client{
		Id:           in.Id,
		EnterpriseId: in.EnterpriseId,
		WorkshopId:   in.WorkshopId,
		Name:         in.Name,
		PhoneNumber:  in.PhoneNumber,
		Email:        in.Email,
		Address:      in.Address,
		Remark:       in.Remark,
	})
}
