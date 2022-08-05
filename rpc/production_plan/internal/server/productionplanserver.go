// Code generated by goctl. DO NOT EDIT!
// Source: production_plan.proto

package server

import (
	"context"

	"air-grating-pms-backend/rpc/production_plan/internal/logic"
	"air-grating-pms-backend/rpc/production_plan/internal/svc"
	"air-grating-pms-backend/rpc/production_plan/pb"
)

type ProductionPlanServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductionPlanServer
}

func NewProductionPlanServer(svcCtx *svc.ServiceContext) *ProductionPlanServer {
	return &ProductionPlanServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductionPlanServer) Insert(ctx context.Context, in *pb.ProductionPlanInfo) (*pb.InsertResp, error) {
	l := logic.NewInsertLogic(ctx, s.svcCtx)
	return l.Insert(in)
}

func (s *ProductionPlanServer) Delete(ctx context.Context, in *pb.DeleteReq) (*pb.Empty, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *ProductionPlanServer) Update(ctx context.Context, in *pb.ProductionPlanInfo) (*pb.Empty, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *ProductionPlanServer) FindOneById(ctx context.Context, in *pb.FindOneByIdReq) (*pb.ProductionPlanInfo, error) {
	l := logic.NewFindOneByIdLogic(ctx, s.svcCtx)
	return l.FindOneById(in)
}

func (s *ProductionPlanServer) FindListByWorkshop(ctx context.Context, in *pb.FindListByWorkshopReq) (*pb.ProductionPlanList, error) {
	l := logic.NewFindListByWorkshopLogic(ctx, s.svcCtx)
	return l.FindListByWorkshop(in)
}

func (s *ProductionPlanServer) FindListByEnterprise(ctx context.Context, in *pb.FindListByEnterpriseReq) (*pb.ProductionPlanList, error) {
	l := logic.NewFindListByEnterpriseLogic(ctx, s.svcCtx)
	return l.FindListByEnterprise(in)
}