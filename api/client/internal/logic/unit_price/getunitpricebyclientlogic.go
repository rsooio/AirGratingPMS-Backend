package unit_price

import (
	"context"

	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnitPriceByClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUnitPriceByClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnitPriceByClientLogic {
	return &GetUnitPriceByClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUnitPriceByClientLogic) GetUnitPriceByClient(req *types.GetUnitPriceByClientReq) (resp *types.GetUnitPriceByClientReply, err error) {
	// todo: add your logic here and delete this line

	return
}
