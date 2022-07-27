package logic

import (
	"context"
	"database/sql"

	"air-grating-pms-backend/rpc/enterprise/internal/svc"
	"air-grating-pms-backend/rpc/enterprise/pb"

	"github.com/dtm-labs/dtmgrpc"
	// _ "github.com/go-sql-driver/mysql"
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

func (l *InsertXaLogic) InsertXa(in *pb.EnterpriseInfo) (*pb.InsertResp, error) {
	var id int64
	err := dtmgrpc.XaLocalTransaction(l.ctx, l.svcCtx.Config.Mysql.DBConf, func(db *sql.DB, xa *dtmgrpc.XaGrpc) error {
		reault, err := db.Exec("INSERT INTO `enterprise` (name,address,remark,version) values (?, ?, ?, ?)", in.GetName(), in.GetAddress(), in.GetRemark(), in.GetVersion())
		if err != nil {
			return err
		}

		id, err = reault.LastInsertId()
		return err
	})
	if err != nil {
		return nil, err
	}

	return &pb.InsertResp{Id: id}, err
}
