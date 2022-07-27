package logic

import (
	"context"
	"fmt"

	"air-grating-pms-backend/rpc/authentication/internal/svc"
	"air-grating-pms-backend/rpc/authentication/types/authentication"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnforceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEnforceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnforceLogic {
	return &EnforceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EnforceLogic) Enforce(in *authentication.AuthReq) (*authentication.AuthResp, error) {
	fmt.Printf("Subject: %s\nObject: %s\nAction: %s\n", in.Subject, in.Object, in.Action)

	ok, err := l.svcCtx.Casbin.Enforce(in.Subject, in.Object, in.Action)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Permission: %v\n", ok)
	return &authentication.AuthResp{
		Permission: ok,
	}, nil
}
