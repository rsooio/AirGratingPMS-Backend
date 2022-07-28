package enterprise

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/enterprise/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEnterpriseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEnterpriseLogic {
	return &DeleteEnterpriseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEnterpriseLogic) DeleteEnterprise() (resp *types.DeleteEnterpriseReply, err error) {
	_, err = l.svcCtx.EnterpriseRPC.Delete(l.ctx, &pb.DeleteReq{
		Id: utils.GetEnterpriseId(l.ctx),
	})

	return &types.DeleteEnterpriseReply{
		Message: "OK",
	}, err
}
