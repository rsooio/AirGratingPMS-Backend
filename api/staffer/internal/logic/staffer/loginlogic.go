package staffer

import (
	"context"
	"errors"
	"strings"
	"time"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/common/middleware"
	"air-grating-pms-backend/model/staffer"
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

	enterpriseID := l.ctx.Value(middleware.EnterpriseID).(int64)
	stafferInfo, err := l.svcCtx.StafferModel.FindOneByEnterpriseIdUsername(l.ctx, enterpriseID, req.Username)
	switch err {
	case nil:
	case staffer.ErrNotFound:
		return nil, errors.New("用户名不存在")
	default:
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
	claims["uid"] = stafferInfo.Id
	claims["rol"] = stafferInfo.Role
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwt, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}

	return &types.LoginReply{
		Id:           stafferInfo.Id,
		Name:         stafferInfo.Name,
		Gender:       stafferInfo.Gender.String,
		Role:         stafferInfo.Role,
		AccessToken:  jwt,
		AccessExpire: accessExpire,
	}, err
}
