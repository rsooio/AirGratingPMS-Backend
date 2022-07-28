package logic

import (
	"context"
	"database/sql"

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
		Id:             in.GetId(),
		EnterpriseId:   in.GetExpireTime(),
		WorkshopId:     in.GetWorkshopId(),
		Username:       in.GetUsername(),
		Role:           in.GetRole(),
		Name:           in.GetName(),
		HashedPassword: in.GetHashedPassword(),
		Gender:         sql.NullString{String: in.GetGender(), Valid: true},
		PhoneNumber:    sql.NullString{String: in.GetPhoneNumber(), Valid: true},
		Email:          sql.NullString{String: in.GetEmail(), Valid: true},
		Address:        sql.NullString{String: in.GetAddress(), Valid: true},
		ExpireTime:     in.GetExpireTime(),
		Remark:         sql.NullString{String: in.GetRemark(), Valid: true},
		Version:        in.GetVersion(),
	})
}
