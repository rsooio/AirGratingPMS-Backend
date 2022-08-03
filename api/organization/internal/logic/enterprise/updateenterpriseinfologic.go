package enterprise

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/enterprise/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEnterpriseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEnterpriseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEnterpriseInfoLogic {
	return &UpdateEnterpriseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEnterpriseInfoLogic) UpdateEnterpriseInfo(req *types.UpdateEnterpriseInfoReq) (resp *types.UpdateEnterpriseInfoReply, err error) {
	_, err = l.svcCtx.EnterpriseRpc.Update(l.ctx, &pb.EnterpriseInfo{
		Id:      utils.GetEnterpriseId(l.ctx),
		Name:    req.Name,
		Address: req.Address,
		Remark:  req.Remark,
	})

	return &types.UpdateEnterpriseInfoReply{
		Message: "OK",
	}, err
}
