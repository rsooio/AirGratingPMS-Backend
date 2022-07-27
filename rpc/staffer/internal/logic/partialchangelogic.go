package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/staffer"
	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartialChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPartialChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartialChangeLogic {
	return &PartialChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PartialChangeLogic) PartialChange(in *pb.StafferInfoWithId) (*pb.Empty, error) {
	now, err := l.svcCtx.StafferModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	out := staffer.Staffer{
		Id: in.GetId(),
	}

	if in.GetEnterpriseId() == out.EnterpriseId {
		out.EnterpriseId = now.EnterpriseId
	} else {
		out.EnterpriseId = in.GetEnterpriseId()
	}
	if in.GetWorkshopId() == out.WorkshopId {
		out.WorkshopId = now.WorkshopId
	} else {
		out.WorkshopId = in.GetWorkshopId()
	}
	if in.GetUsername() == out.Username {
		out.Username = now.Username
	} else {
		out.Username = in.GetUsername()
	}
	if in.GetRole() == out.Role {
		out.Role = now.Role
	} else {
		out.Role = in.GetRole()
	}
	if in.GetName() == out.Name {
		out.Name = now.Name
	} else {
		out.Name = in.GetName()
	}
	if in.GetHashedPassword() == out.HashedPassword {
		out.HashedPassword = now.HashedPassword
	} else {
		out.HashedPassword = in.GetHashedPassword()
	}
	if in.GetGender() == out.Gender.String {
		out.Gender = now.Gender
	} else {
		out.Gender = sql.NullString{String: in.GetGender(), Valid: in.GetGender() != ""}
	}
	if in.GetPhoneNumber() == out.PhoneNumber.String {
		out.PhoneNumber = now.PhoneNumber
	} else {
		out.PhoneNumber = sql.NullString{String: in.GetPhoneNumber(), Valid: in.GetPhoneNumber() != ""}
	}
	if in.GetEmail() == out.Email.String {
		out.Email = now.Email
	} else {
		out.PhoneNumber = sql.NullString{String: in.GetEmail(), Valid: in.GetEmail() != ""}
	}
	if in.GetAddress() == out.Address.String {
		out.Address = now.Address
	} else {
		out.Address = sql.NullString{String: in.GetAddress(), Valid: in.GetAddress() != ""}
	}
	if in.GetExpireTime() == out.ExpireTime {
		out.ExpireTime = now.ExpireTime
	} else {
		out.ExpireTime = in.GetExpireTime()
	}
	if in.GetRemark() == out.Remark.String {
		out.Remark = now.Remark
	} else {
		out.Remark = sql.NullString{String: in.GetRemark(), Valid: in.GetRemark() != ""}
	}
	if in.GetVersion() == out.Version {
		out.Version = now.Version
	} else {
		out.Version = in.GetVersion()
	}

	return &pb.Empty{}, l.svcCtx.StafferModel.Update(l.ctx, &out)
}
