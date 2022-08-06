package logic

import (
	"context"

	"air-grating-pms-backend/model/client"
	"air-grating-pms-backend/rpc/client/internal/svc"
	"air-grating-pms-backend/rpc/client/pb"

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

func (l *InsertLogic) Insert(in *pb.ClientInfo) (*pb.InsertResp, error) {
	result, err := l.svcCtx.ClientModel.Insert(l.ctx, &client.Client{
		EnterpriseId: in.EnterpriseId,
		WorkshopId:   in.WorkshopId,
		Name:         in.Name,
		PhoneNumber:  in.PhoneNumber,
		Email:        in.Email,
		Address:      in.Address,
		Remark:       in.Remark,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return &pb.InsertResp{
		Id: id,
	}, err
}
