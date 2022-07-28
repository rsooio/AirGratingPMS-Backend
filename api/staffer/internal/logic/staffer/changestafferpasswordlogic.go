package staffer

import (
	"context"
	"errors"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils"
	"air-grating-pms-backend/utils/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeStafferPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeStafferPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeStafferPasswordLogic {
	return &ChangeStafferPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeStafferPasswordLogic) ChangeStafferPassword(req *types.ChangeStafferPasswordReq) (resp *types.ChangeStafferPasswordReply, err error) {
	uid := utils.GetStafferId(l.ctx)

	info, err := l.svcCtx.StafferRPC.FindOneById(l.ctx, &pb.FindOneByIdReq{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	if !bcrypt.Verify(info.HashedPassword, req.OldPassword) {
		return nil, errors.New("incorrect password")
	}

	hashedPassword, err := bcrypt.Encrypt(req.NewPassword)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.StafferRPC.PartialUpdate(l.ctx, &pb.StafferInfo{
		Id:             uid,
		HashedPassword: hashedPassword,
	})

	return &types.ChangeStafferPasswordReply{
		Message: "OK",
	}, err
}
