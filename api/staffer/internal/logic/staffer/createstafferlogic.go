package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils"
	"air-grating-pms-backend/utils/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStafferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStafferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStafferLogic {
	return &CreateStafferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStafferLogic) CreateStaffer(req *types.CreateStafferReq) (resp *types.CreateReply, err error) {
	hashedPassword, err := bcrypt.Encrypt(l.svcCtx.Config.DefaultPassword)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.StafferRPC.Insert(l.ctx, &pb.StafferInfo{
		EnterpriseId:   utils.GetEnterpriseId(l.ctx),
		WorkshopId:     req.WorkshopId,
		Username:       req.Username,
		Role:           req.Role,
		Name:           req.Name,
		HashedPassword: hashedPassword,
		Gender:         req.Gender,
		PhoneNumber:    req.PhoneNumber,
		Email:          req.Email,
		Address:        req.Address,
		ExpireTime:     l.svcCtx.Config.Auth.AccessExpire,
		Remark:         req.Remark,
	})

	return &types.CreateReply{
		Message:  "OK",
		Password: l.svcCtx.Config.DefaultPassword,
	}, err
}
