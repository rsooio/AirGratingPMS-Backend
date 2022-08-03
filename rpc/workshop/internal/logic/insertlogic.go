package logic

import (
	"context"

	"air-grating-pms-backend/model/workshop"
	"air-grating-pms-backend/rpc/workshop/internal/svc"
	"air-grating-pms-backend/rpc/workshop/pb"

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

func (l *InsertLogic) Insert(in *pb.WorkshopInfo) (*pb.InsertResp, error) {
	result, err := l.svcCtx.WorkshopModel.Insert(l.ctx, &workshop.Workshop{
		Id:           in.Id,
		EnterpriseId: in.EnterpriseId,
		Name:         in.Name,
		Address:      in.Address,
		PhoneNumber:  in.PhoneNumber,
		ManagerId:    in.ManagerId,
		Remark:       in.Remark,
		Version:      in.Version,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return &pb.InsertResp{Id: id}, err
}
