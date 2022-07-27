package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/model/enterprise"
	"air-grating-pms-backend/rpc/enterprise/internal/svc"
	"air-grating-pms-backend/rpc/enterprise/pb"

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

func (l *InsertLogic) Insert(in *pb.EnterpriseInfo) (*pb.InsertResp, error) {
	result, err := l.svcCtx.EnterpriseModel.Insert(l.ctx, &enterprise.Enterprise{
		Name:    in.GetName(),
		Address: sql.NullString{String: in.GetAddress(), Valid: in.GetAddress() != ""},
		Remark:  sql.NullString{String: in.GetRemark(), Valid: in.GetRemark() != ""},
		Version: in.GetVersion(),
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return &pb.InsertResp{Id: id}, err
}
