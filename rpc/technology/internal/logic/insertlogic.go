package logic

import (
	"context"

	"air-grating-pms-backend/model/technology"
	"air-grating-pms-backend/rpc/technology/internal/svc"
	"air-grating-pms-backend/rpc/technology/pb"

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

func (l *InsertLogic) Insert(in *pb.TechnologyInfo) (*pb.InsertResp, error) {
	result, err := l.svcCtx.TechnologyModel.Insert(l.ctx, &technology.Technology{
		EnterpriseId: in.EnterpriseId,
		WorkshopId:   in.WorkshopId,
		Name:         in.Name,
		Info:         in.Info,
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
