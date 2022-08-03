package logic

import (
	"context"

	"air-grating-pms-backend/rpc/technology/internal/svc"
	"air-grating-pms-backend/rpc/technology/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *pb.DeleteReq) (*pb.Empty, error) {
	return &pb.Empty{}, l.svcCtx.TechnologyModel.Delete(l.ctx, in.Id)
}
