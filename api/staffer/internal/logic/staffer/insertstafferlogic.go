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

type InsertStafferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertStafferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertStafferLogic {
	return &InsertStafferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertStafferLogic) InsertStaffer(req *types.InsertStafferReq) (resp *types.InsertStafferReply, err error) {
	workshopId := req.WorkshopId
	if utils.GetRole(l.ctx) != "boss" {
		workshopId = utils.GetWorkshopId(l.ctx)
	}

	hashedPassword, err := bcrypt.Encrypt(l.svcCtx.Config.DefaultPassword)
	if err != nil {
		return nil, err
	}

	info, err := l.svcCtx.StafferRPC.Insert(l.ctx, &pb.StafferInfo{
		EnterpriseId:   utils.GetEnterpriseId(l.ctx),
		WorkshopId:     workshopId,
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

	return &types.InsertStafferReply{
		Id:       info.Id,
		Message:  "OK",
		Password: l.svcCtx.Config.DefaultPassword,
	}, err
}
