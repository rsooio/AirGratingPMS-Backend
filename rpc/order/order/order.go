// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package order

import (
	"context"

	"air-grating-pms-backend/rpc/order/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeleteReq               = pb.DeleteReq
	Empty                   = pb.Empty
	FindListByEnterpriseReq = pb.FindListByEnterpriseReq
	FindListByWorkshopReq   = pb.FindListByWorkshopReq
	FindOneByIdReq          = pb.FindOneByIdReq
	FindOneByNameReq        = pb.FindOneByNameReq
	InsertResp              = pb.InsertResp
	OrderInfo               = pb.OrderInfo
	OrderList               = pb.OrderList

	Order interface {
		Insert(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*InsertResp, error)
		Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error)
		Update(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*Empty, error)
		//  rpc CustomUpdate (OrderInfo) returns (Empty);
		FindOneById(ctx context.Context, in *FindOneByIdReq, opts ...grpc.CallOption) (*OrderInfo, error)
		FindListByWorkshop(ctx context.Context, in *FindListByWorkshopReq, opts ...grpc.CallOption) (*OrderList, error)
		FindListByEnterprise(ctx context.Context, in *FindListByEnterpriseReq, opts ...grpc.CallOption) (*OrderList, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Insert(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*InsertResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.Insert(ctx, in, opts...)
}

func (m *defaultOrder) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.Delete(ctx, in, opts...)
}

func (m *defaultOrder) Update(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

//  rpc CustomUpdate (OrderInfo) returns (Empty);
func (m *defaultOrder) FindOneById(ctx context.Context, in *FindOneByIdReq, opts ...grpc.CallOption) (*OrderInfo, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.FindOneById(ctx, in, opts...)
}

func (m *defaultOrder) FindListByWorkshop(ctx context.Context, in *FindListByWorkshopReq, opts ...grpc.CallOption) (*OrderList, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.FindListByWorkshop(ctx, in, opts...)
}

func (m *defaultOrder) FindListByEnterprise(ctx context.Context, in *FindListByEnterpriseReq, opts ...grpc.CallOption) (*OrderList, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.FindListByEnterprise(ctx, in, opts...)
}
