package enterprise

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/enterprise/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeEnterpriseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeEnterpriseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeEnterpriseInfoLogic {
	return &ChangeEnterpriseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeEnterpriseInfoLogic) ChangeEnterpriseInfo(req *types.ChangeEnterpriseInfoReq) (resp *types.ChangeEnterpriseInfoReply, err error) {
	_, err = l.svcCtx.EnterpriseRPC.Update(l.ctx, &pb.EnterpriseInfo{
		Id:      l.ctx.Value("ent").(int64),
		Name:    req.Name,
		Address: req.Address,
		Remark:  req.Remark,
	})

	return &types.ChangeEnterpriseInfoReply{
		Message: "OK",
	}, err
}
