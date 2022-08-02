// Code generated by goctl. DO NOT EDIT!
// Source: product_set.proto

package server

import (
	"context"

	"air-grating-pms-backend/rpc/product_set/internal/logic"
	"air-grating-pms-backend/rpc/product_set/internal/svc"
	"air-grating-pms-backend/rpc/product_set/pb"
)

type ProductsetServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductsetServer
}

func NewProductsetServer(svcCtx *svc.ServiceContext) *ProductsetServer {
	return &ProductsetServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductsetServer) Insert(ctx context.Context, in *pb.ProductSetInfo) (*pb.InsertResp, error) {
	l := logic.NewInsertLogic(ctx, s.svcCtx)
	return l.Insert(in)
}

func (s *ProductsetServer) Delete(ctx context.Context, in *pb.DeleteReq) (*pb.Empty, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *ProductsetServer) Update(ctx context.Context, in *pb.ProductSetInfo) (*pb.Empty, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

//  rpc CustomUpdate (ProductSetInfo) returns (Empty);
func (s *ProductsetServer) FindOneById(ctx context.Context, in *pb.FindOneByIdReq) (*pb.ProductSetInfo, error) {
	l := logic.NewFindOneByIdLogic(ctx, s.svcCtx)
	return l.FindOneById(in)
}

func (s *ProductsetServer) FindListByOrder(ctx context.Context, in *pb.FindListByOrderReq) (*pb.ProductSetList, error) {
	l := logic.NewFindListByOrderLogic(ctx, s.svcCtx)
	return l.FindListByOrder(in)
}