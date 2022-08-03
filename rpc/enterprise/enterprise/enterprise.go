// Code generated by goctl. DO NOT EDIT!
// Source: enterprise.proto

package enterprise

import (
	"context"

	"air-grating-pms-backend/rpc/enterprise/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeleteReq        = pb.DeleteReq
	Empty            = pb.Empty
	EnterpriseInfo   = pb.EnterpriseInfo
	FindOneByNameReq = pb.FindOneByNameReq
	InsertResp       = pb.InsertResp

	Enterprise interface {
		Insert(ctx context.Context, in *EnterpriseInfo, opts ...grpc.CallOption) (*InsertResp, error)
		Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error)
		Update(ctx context.Context, in *EnterpriseInfo, opts ...grpc.CallOption) (*Empty, error)
		FindOneByName(ctx context.Context, in *FindOneByNameReq, opts ...grpc.CallOption) (*EnterpriseInfo, error)
		InsertXa(ctx context.Context, in *EnterpriseInfo, opts ...grpc.CallOption) (*InsertResp, error)
	}

	defaultEnterprise struct {
		cli zrpc.Client
	}
)

func NewEnterprise(cli zrpc.Client) Enterprise {
	return &defaultEnterprise{
		cli: cli,
	}
}

func (m *defaultEnterprise) Insert(ctx context.Context, in *EnterpriseInfo, opts ...grpc.CallOption) (*InsertResp, error) {
	client := pb.NewEnterpriseClient(m.cli.Conn())
	return client.Insert(ctx, in, opts...)
}

func (m *defaultEnterprise) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewEnterpriseClient(m.cli.Conn())
	return client.Delete(ctx, in, opts...)
}

func (m *defaultEnterprise) Update(ctx context.Context, in *EnterpriseInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewEnterpriseClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultEnterprise) FindOneByName(ctx context.Context, in *FindOneByNameReq, opts ...grpc.CallOption) (*EnterpriseInfo, error) {
	client := pb.NewEnterpriseClient(m.cli.Conn())
	return client.FindOneByName(ctx, in, opts...)
}

func (m *defaultEnterprise) InsertXa(ctx context.Context, in *EnterpriseInfo, opts ...grpc.CallOption) (*InsertResp, error) {
	client := pb.NewEnterpriseClient(m.cli.Conn())
	return client.InsertXa(ctx, in, opts...)
}
