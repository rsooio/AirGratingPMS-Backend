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

type UpdateStafferPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStafferPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStafferPasswordLogic {
	return &UpdateStafferPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStafferPasswordLogic) UpdateStafferPassword(req *types.UpdateStafferPasswordReq) (resp *types.UpdateStafferPasswordReply, err error) {
	id := utils.GetStafferId(l.ctx)
	stafferInfo, err := l.svcCtx.StafferRPC.FindOneById(l.ctx, &pb.FindOneByIdReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	if !bcrypt.Verify(stafferInfo.HashedPassword, req.OldPassword) {
		return nil, errors.New("用户密码不正确")
	}

	hashedPassword, err := bcrypt.Encrypt(req.NewPassword)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.StafferRPC.Update(l.ctx, &pb.StafferInfo{
		Id:             id,
		HashedPassword: hashedPassword,
	})
	return &types.UpdateStafferPasswordReply{
		Message: "OK",
	}, err
}
