package enterprise

import (
	"context"

	"air-grating-pms-backend/api/organization/internal/svc"
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/enterprise/enterprise"
	epb "air-grating-pms-backend/rpc/enterprise/pb"
	spb "air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/rpc/staffer/staffer"
	"air-grating-pms-backend/utils/bcrypt"

	_ "github.com/dtm-labs/driver-gozero"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEnterpriseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEnterpriseLogic {
	return &CreateEnterpriseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEnterpriseLogic) CreateEnterprise(req *types.CreateEnterpriseReq) (resp *types.CreateEnterpriseReply, err error) {
	hashedPassword, err := bcrypt.Encrypt(req.Password)
	if err != nil {
		return nil, err
	}

	gid := dtmgrpc.MustGenGid(l.svcCtx.Config.DTM.Server)
	err = dtmgrpc.XaGlobalTransaction("localhost:36790", gid, func(xa *dtmgrpc.XaGrpc) error {
		r := &epb.InsertResp{}
		err := xa.CallBranch(&epb.EnterpriseInfo{
			Name:    req.Name,
			Address: req.Address,
			Remark:  req.Remark,
		}, l.svcCtx.EnterpriseTarget+enterprise.InsertXaPath, r)
		if err != nil {
			return err
		}

		r2 := &spb.InsertResp{}
		err = xa.CallBranch(&spb.StafferInfo{
			EnterpriseId:   r.Id,
			Username:       req.BossName,
			HashedPassword: hashedPassword,
		}, l.svcCtx.StafferTarget+staffer.InsertXaPath, r2)
		return err
	})

	return &types.CreateEnterpriseReply{Message: "OK"}, err
}
