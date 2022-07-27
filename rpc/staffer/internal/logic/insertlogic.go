package logic

import (
	"context"
	"database/sql"

	staffermodel "air-grating-pms-backend/model/staffer"
	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLogic) Insert(in *pb.StafferInfo) (*pb.InsertResp, error) {
	result, err := l.svcCtx.StafferModel.Insert(l.ctx, &staffermodel.Staffer{
		EnterpriseId:   in.GetEnterpriseId(),
		WorkshopId:     in.GetWorkshopId(),
		Username:       in.GetUsername(),
		Role:           in.GetRole(),
		Name:           in.GetName(),
		HashedPassword: in.GetHashedPassword(),
		Gender:         sql.NullString{String: in.GetGender(), Valid: in.GetGender() != ""},
		PhoneNumber:    sql.NullString{String: in.GetPhoneNumber(), Valid: in.GetPhoneNumber() != ""},
		Email:          sql.NullString{String: in.GetEmail(), Valid: in.GetEmail() != ""},
		Address:        sql.NullString{String: in.GetAddress(), Valid: in.GetAddress() != ""},
		ExpireTime:     in.ExpireTime,
		Remark:         sql.NullString{String: in.GetRemark(), Valid: in.GetRemark() != ""},
		Version:        in.Version,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return &pb.InsertResp{Id: id}, err
}
