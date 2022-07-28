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
	// ep := &enterprise.Enterprise{
	// 	Name:    req.Name,
	// 	Address: sql.NullString{String: req.Address, Valid: req.Address != ""},
	// 	Remark:  sql.NullString{String: req.Remark, Valid: req.Remark != ""},
	// }

	// result, err := l.svcCtx.EnterpriseModel.Insert(l.ctx, ep)
	// if err != nil {
	// 	return nil, err
	// }

	// eid, err := result.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// }

	// hashedPassword, err := bcrypt.Encrypt(req.Password)
	// if err != nil {
	// 	err2 := l.svcCtx.EnterpriseModel.Delete(l.ctx, eid)
	// 	if err2 != nil {
	// 		return nil, fmt.Errorf("%s||%s", err.Error(), err2.Error())
	// 	}

	// 	return nil, err
	// }

	// result, err = l.svcCtx.StafferModel.Insert(l.ctx, &staffer.Staffer{
	// 	Username:       req.BossName,
	// 	HashedPassword: hashedPassword,
	// })
	// if err != nil {
	// 	err2 := l.svcCtx.EnterpriseModel.Delete(l.ctx, eid)
	// 	if err2 != nil {
	// 		return nil, fmt.Errorf("%s||%s", err.Error(), err2.Error())
	// 	}

	// 	return nil, err
	// }

	// uid, err := result.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// }

	// ep.Id = eid
	// ep.BossId = sql.NullInt64{Int64: uid, Valid: true}
	// err = l.svcCtx.EnterpriseModel.Update(l.ctx, ep)
	// if err != nil {
	// 	l.svcCtx.EnterpriseModel.Delete(l.ctx, eid)
	// 	l.svcCtx.StafferModel.Delete(l.ctx, uid)

	// 	return nil, err
	// }

	// return &types.CreateEnterpriseReply{
	// 	Message: "OK",
	// }, nil

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
