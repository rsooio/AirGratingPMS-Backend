package unit_price

import (
	"context"

	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUnitPriceByClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUnitPriceByClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUnitPriceByClientLogic {
	return &SetUnitPriceByClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUnitPriceByClientLogic) SetUnitPriceByClient(req *types.SetUnitPriceByClientReq) (resp *types.SetUnitPriceByClientReply, err error) {
	// todo: add your logic here and delete this line

	return
}
