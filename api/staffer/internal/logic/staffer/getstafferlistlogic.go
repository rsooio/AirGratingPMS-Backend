package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStafferListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStafferListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStafferListLogic {
	return &GetStafferListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStafferListLogic) GetStafferList(req *types.GetStafferListReq) (resp *types.GetStafferListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
