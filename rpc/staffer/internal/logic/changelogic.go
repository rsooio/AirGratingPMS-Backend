package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/staffer"
	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeLogic {
	return &ChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeLogic) Change(in *pb.StafferInfoWithId) (*pb.Empty, error) {
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
