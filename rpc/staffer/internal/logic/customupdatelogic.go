package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/staffer"
	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomUpdateLogic {
	return &CustomUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomUpdateLogic) CustomUpdate(in *pb.StafferInfo) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.StafferModel.CustomUpdate(l.ctx, &staffer.Staffer{
		Id:          in.GetId(),
		WorkshopId:  in.GetWorkshopId(),
		Username:    in.GetUsername(),
		Role:        in.GetRole(),
		Name:        in.GetName(),
		Gender:      sql.NullString{String: in.GetGender(), Valid: true},
		PhoneNumber: sql.NullString{String: in.GetPhoneNumber(), Valid: true},
		Email:       sql.NullString{String: in.GetEmail(), Valid: true},
		Address:     sql.NullString{String: in.GetAddress(), Valid: true},
		ExpireTime:  in.GetExpireTime(),
		Remark:      sql.NullString{String: in.GetRemark(), Valid: true},
		Version:     in.GetVersion(),
	})
}
