package logic

import (
	"context"

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
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return &pb.InsertResp{Id: id}, err
}
