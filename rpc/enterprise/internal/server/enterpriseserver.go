// Code generated by goctl. DO NOT EDIT!
// Source: enterprise.proto

package server

import (
	"context"

	"air-grating-pms-backend/rpc/enterprise/internal/logic"
	"air-grating-pms-backend/rpc/enterprise/internal/svc"
	"air-grating-pms-backend/rpc/enterprise/pb"
)

type EnterpriseServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedEnterpriseServer
}

func NewEnterpriseServer(svcCtx *svc.ServiceContext) *EnterpriseServer {
	return &EnterpriseServer{
		svcCtx: svcCtx,
	}
}

func (s *EnterpriseServer) Insert(ctx context.Context, in *pb.EnterpriseInfo) (*pb.InsertResp, error) {
	l := logic.NewInsertLogic(ctx, s.svcCtx)
	return l.Insert(in)
}

func (s *EnterpriseServer) Delete(ctx context.Context, in *pb.DeleteReq) (*pb.Empty, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *EnterpriseServer) Change(ctx context.Context, in *pb.EnterpriseInfoWithId) (*pb.Empty, error) {
	l := logic.NewChangeLogic(ctx, s.svcCtx)
	return l.Change(in)
}

func (s *EnterpriseServer) PartialChange(ctx context.Context, in *pb.EnterpriseInfoWithId) (*pb.Empty, error) {
	l := logic.NewPartialChangeLogic(ctx, s.svcCtx)
	return l.PartialChange(in)
}

func (s *EnterpriseServer) FindOneByName(ctx context.Context, in *pb.FindOneByNameReq) (*pb.EnterpriseInfo, error) {
	l := logic.NewFindOneByNameLogic(ctx, s.svcCtx)
	return l.FindOneByName(in)
}

func (s *EnterpriseServer) InsertXa(ctx context.Context, in *pb.EnterpriseInfo) (*pb.InsertResp, error) {
	l := logic.NewInsertXaLogic(ctx, s.svcCtx)
	return l.InsertXa(in)
}