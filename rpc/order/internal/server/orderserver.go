// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package server

import (
	"context"

	"air-grating-pms-backend/rpc/order/internal/logic"
	"air-grating-pms-backend/rpc/order/internal/svc"
	"air-grating-pms-backend/rpc/order/pb"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) Insert(ctx context.Context, in *pb.OrderInfo) (*pb.InsertResp, error) {
	l := logic.NewInsertLogic(ctx, s.svcCtx)
	return l.Insert(in)
}

func (s *OrderServer) Delete(ctx context.Context, in *pb.DeleteReq) (*pb.Empty, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *OrderServer) Update(ctx context.Context, in *pb.OrderInfo) (*pb.Empty, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

//  rpc CustomUpdate (OrderInfo) returns (Empty);
func (s *OrderServer) FindOneById(ctx context.Context, in *pb.FindOneByIdReq) (*pb.OrderInfo, error) {
	l := logic.NewFindOneByIdLogic(ctx, s.svcCtx)
	return l.FindOneById(in)
}

func (s *OrderServer) FindListByWorkshop(ctx context.Context, in *pb.FindListByWorkshopReq) (*pb.OrderList, error) {
	l := logic.NewFindListByWorkshopLogic(ctx, s.svcCtx)
	return l.FindListByWorkshop(in)
}

func (s *OrderServer) FindListByEnterprise(ctx context.Context, in *pb.FindListByEnterpriseReq) (*pb.OrderList, error) {
	l := logic.NewFindListByEnterpriseLogic(ctx, s.svcCtx)
	return l.FindListByEnterprise(in)
}