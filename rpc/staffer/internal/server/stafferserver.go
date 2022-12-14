// Code generated by goctl. DO NOT EDIT!
// Source: staffer.proto

package server

import (
	"context"

	"air-grating-pms-backend/rpc/staffer/internal/logic"
	"air-grating-pms-backend/rpc/staffer/internal/svc"
	"air-grating-pms-backend/rpc/staffer/pb"
)

type StafferServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedStafferServer
}

func NewStafferServer(svcCtx *svc.ServiceContext) *StafferServer {
	return &StafferServer{
		svcCtx: svcCtx,
	}
}

func (s *StafferServer) Insert(ctx context.Context, in *pb.StafferInfo) (*pb.InsertResp, error) {
	l := logic.NewInsertLogic(ctx, s.svcCtx)
	return l.Insert(in)
}

func (s *StafferServer) Delete(ctx context.Context, in *pb.DeleteReq) (*pb.Empty, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *StafferServer) Update(ctx context.Context, in *pb.StafferInfo) (*pb.Empty, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *StafferServer) FindOneById(ctx context.Context, in *pb.FindOneByIdReq) (*pb.StafferInfo, error) {
	l := logic.NewFindOneByIdLogic(ctx, s.svcCtx)
	return l.FindOneById(in)
}

func (s *StafferServer) FindOneByName(ctx context.Context, in *pb.FindOneByNameReq) (*pb.StafferInfo, error) {
	l := logic.NewFindOneByNameLogic(ctx, s.svcCtx)
	return l.FindOneByName(in)
}

func (s *StafferServer) FindListByWorkshop(ctx context.Context, in *pb.FindListByWorkshopReq) (*pb.StafferList, error) {
	l := logic.NewFindListByWorkshopLogic(ctx, s.svcCtx)
	return l.FindListByWorkshop(in)
}

func (s *StafferServer) FindListByEnterprise(ctx context.Context, in *pb.FindListByEnterpriseReq) (*pb.StafferList, error) {
	l := logic.NewFindListByEnterpriseLogic(ctx, s.svcCtx)
	return l.FindListByEnterprise(in)
}

func (s *StafferServer) InsertXa(ctx context.Context, in *pb.StafferInfo) (*pb.InsertResp, error) {
	l := logic.NewInsertXaLogic(ctx, s.svcCtx)
	return l.InsertXa(in)
}
