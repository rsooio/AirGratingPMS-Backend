package logic

import (
	"context"

	"air-grating-pms-backend/model/staffer"
	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"

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

func (l *UpdateLogic) Update(in *pb.StafferInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.StafferModel.Update(l.ctx, &staffer.Staffer{
		Id:             in.Id,
		EnterpriseId:   in.EnterpriseId,
		WorkshopId:     in.WorkshopId,
		Username:       in.Username,
		Role:           in.Role,
		Name:           in.Name,
		HashedPassword: in.HashedPassword,
		Gender:         in.Gender,
		PhoneNumber:    in.PhoneNumber,
		Email:          in.Email,
		Address:        in.Address,
		ExpireTime:     in.ExpireTime,
		Remark:         in.Remark,
		Version:        in.Version,
	})
}
