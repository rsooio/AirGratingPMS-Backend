package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetStafferPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetStafferPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetStafferPasswordLogic {
	return &ResetStafferPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetStafferPasswordLogic) ResetStafferPassword(req *types.ResetStafferPasswordReq) (resp *types.ResetStafferPasswordReply, err error) {
	hashedPassword, err := bcrypt.Encrypt(l.svcCtx.Config.DefaultPassword)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.StafferRPC.PartialUpdate(l.ctx, &pb.StafferInfo{
		Id:             req.Id,
		HashedPassword: hashedPassword,
	})

	return &types.ResetStafferPasswordReply{
		Message:  "OK",
		Password: l.svcCtx.Config.DefaultPassword,
	}, err
}
