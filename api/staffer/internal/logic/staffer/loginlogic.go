package staffer

import (
	"context"
	"errors"
	"strings"
	"time"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils/bcrypt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	stafferInfo, err := l.svcCtx.StafferRPC.FindOneByName(l.ctx, &pb.FindOneByNameReq{
		EnterpriseId: req.EnterpriseId,
		Username:     req.Username,
	})
	if err != nil {
		return nil, err
	}

	if !bcrypt.Verify(stafferInfo.HashedPassword, req.Password) {
		return nil, errors.New("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := stafferInfo.ExpireTime
	if accessExpire < 1800 || accessExpire > 28800 {
		accessExpire = l.svcCtx.Config.Auth.AccessExpire
	}

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["rol"] = stafferInfo.Role
	claims["uid"] = stafferInfo.Id
	claims["wid"] = stafferInfo.WorkshopId
	claims["eid"] = stafferInfo.EnterpriseId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwt, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}

	return &types.LoginReply{
		Id:           stafferInfo.Id,
		Name:         stafferInfo.Name,
		Gender:       stafferInfo.Gender,
		Role:         stafferInfo.Role,
		AccessToken:  jwt,
		AccessExpire: accessExpire,
	}, err
}
