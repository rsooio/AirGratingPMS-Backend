// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package product

import (
	"context"

	"air-grating-pms-backend/rpc/product/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeleteReq               = pb.DeleteReq
	Empty                   = pb.Empty
	FindListByProductSetReq = pb.FindListByProductSetReq
	FindOneByIdReq          = pb.FindOneByIdReq
	InsertResp              = pb.InsertResp
	ProductInfo             = pb.ProductInfo
	ProductList             = pb.ProductList

	Product interface {
		Insert(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*InsertResp, error)
		Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error)
		Update(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error)
		//  rpc CustomUpdate (ProductInfo) returns (Empty);
		FindOneById(ctx context.Context, in *FindOneByIdReq, opts ...grpc.CallOption) (*ProductInfo, error)
		FindListByProductSet(ctx context.Context, in *FindListByProductSetReq, opts ...grpc.CallOption) (*ProductList, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Insert(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*InsertResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Insert(ctx, in, opts...)
}

func (m *defaultProduct) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Delete(ctx, in, opts...)
}

func (m *defaultProduct) Update(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

//  rpc CustomUpdate (ProductInfo) returns (Empty);
func (m *defaultProduct) FindOneById(ctx context.Context, in *FindOneByIdReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.FindOneById(ctx, in, opts...)
}

func (m *defaultProduct) FindListByProductSet(ctx context.Context, in *FindListByProductSetReq, opts ...grpc.CallOption) (*ProductList, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.FindListByProductSet(ctx, in, opts...)
}
