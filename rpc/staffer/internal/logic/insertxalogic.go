package logic

import (
	"context"
	"database/sql"
	"fmt"

	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type InsertXaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertXaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertXaLogic {
	return &InsertXaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertXaLogic) InsertXa(in *pb.StafferInfo) (*pb.InsertResp, error) {
	var id int64
	rows := "enterprise_id,workshop_id,username,role,name,hashed_password,gender,phone_number,email,address,expire_time,remark,version"
	query := fmt.Sprintf("insert into `staffer` (%s) values (?,?,?,?,?,?,?,?,?,?,?,?,?)", rows)
	err := dtmgrpc.XaLocalTransaction(l.ctx, l.svcCtx.Config.Mysql.DBConf, func(db *sql.DB, xa *dtmgrpc.XaGrpc) error {
		result, err := db.Exec(query, in.GetEnterpriseId(), in.GetWorkshopId(), in.GetUsername(), in.GetRole(), in.GetName(), in.GetHashedPassword(), in.GetGender(), in.GetPhoneNumber(), in.GetEmail(), in.GetAddress(), in.GetExpireTime(), in.GetRemark(), in.GetVersion())
		if err != nil {
			return err
		}

		id, err = result.LastInsertId()
		return err
	})

	return &pb.InsertResp{Id: id}, err
}
