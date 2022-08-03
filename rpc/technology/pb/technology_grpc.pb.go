// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: technology.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TechnologyClient is the client API for Technology service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TechnologyClient interface {
	Insert(ctx context.Context, in *TechnologyInfo, opts ...grpc.CallOption) (*InsertResp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *TechnologyInfo, opts ...grpc.CallOption) (*Empty, error)
	FindOneById(ctx context.Context, in *FindOneByIdReq, opts ...grpc.CallOption) (*TechnologyInfo, error)
	FindListByWorkshop(ctx context.Context, in *FindListByWorkshopReq, opts ...grpc.CallOption) (*TechnologyList, error)
	FindListByEnterprise(ctx context.Context, in *FindListByEnterpriseReq, opts ...grpc.CallOption) (*TechnologyList, error)
}

type technologyClient struct {
	cc grpc.ClientConnInterface
}

func NewTechnologyClient(cc grpc.ClientConnInterface) TechnologyClient {
	return &technologyClient{cc}
}

func (c *technologyClient) Insert(ctx context.Context, in *TechnologyInfo, opts ...grpc.CallOption) (*InsertResp, error) {
	out := new(InsertResp)
	err := c.cc.Invoke(ctx, "/technology.technology/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *technologyClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/technology.technology/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *technologyClient) Update(ctx context.Context, in *TechnologyInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/technology.technology/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *technologyClient) FindOneById(ctx context.Context, in *FindOneByIdReq, opts ...grpc.CallOption) (*TechnologyInfo, error) {
	out := new(TechnologyInfo)
	err := c.cc.Invoke(ctx, "/technology.technology/FindOneById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *technologyClient) FindListByWorkshop(ctx context.Context, in *FindListByWorkshopReq, opts ...grpc.CallOption) (*TechnologyList, error) {
	out := new(TechnologyList)
	err := c.cc.Invoke(ctx, "/technology.technology/FindListByWorkshop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *technologyClient) FindListByEnterprise(ctx context.Context, in *FindListByEnterpriseReq, opts ...grpc.CallOption) (*TechnologyList, error) {
	out := new(TechnologyList)
	err := c.cc.Invoke(ctx, "/technology.technology/FindListByEnterprise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TechnologyServer is the server API for Technology service.
// All implementations must embed UnimplementedTechnologyServer
// for forward compatibility
type TechnologyServer interface {
	Insert(context.Context, *TechnologyInfo) (*InsertResp, error)
	Delete(context.Context, *DeleteReq) (*Empty, error)
	Update(context.Context, *TechnologyInfo) (*Empty, error)
	FindOneById(context.Context, *FindOneByIdReq) (*TechnologyInfo, error)
	FindListByWorkshop(context.Context, *FindListByWorkshopReq) (*TechnologyList, error)
	FindListByEnterprise(context.Context, *FindListByEnterpriseReq) (*TechnologyList, error)
	mustEmbedUnimplementedTechnologyServer()
}

// UnimplementedTechnologyServer must be embedded to have forward compatible implementations.
type UnimplementedTechnologyServer struct {
}

func (UnimplementedTechnologyServer) Insert(context.Context, *TechnologyInfo) (*InsertResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedTechnologyServer) Delete(context.Context, *DeleteReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTechnologyServer) Update(context.Context, *TechnologyInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTechnologyServer) FindOneById(context.Context, *FindOneByIdReq) (*TechnologyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneById not implemented")
}
func (UnimplementedTechnologyServer) FindListByWorkshop(context.Context, *FindListByWorkshopReq) (*TechnologyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListByWorkshop not implemented")
}
func (UnimplementedTechnologyServer) FindListByEnterprise(context.Context, *FindListByEnterpriseReq) (*TechnologyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListByEnterprise not implemented")
}
func (UnimplementedTechnologyServer) mustEmbedUnimplementedTechnologyServer() {}

// UnsafeTechnologyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TechnologyServer will
// result in compilation errors.
type UnsafeTechnologyServer interface {
	mustEmbedUnimplementedTechnologyServer()
}

func RegisterTechnologyServer(s grpc.ServiceRegistrar, srv TechnologyServer) {
	s.RegisterService(&Technology_ServiceDesc, srv)
}

func _Technology_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TechnologyInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnologyServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/technology.technology/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnologyServer).Insert(ctx, req.(*TechnologyInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Technology_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnologyServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/technology.technology/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnologyServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Technology_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TechnologyInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnologyServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/technology.technology/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnologyServer).Update(ctx, req.(*TechnologyInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Technology_FindOneById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnologyServer).FindOneById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/technology.technology/FindOneById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnologyServer).FindOneById(ctx, req.(*FindOneByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Technology_FindListByWorkshop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindListByWorkshopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnologyServer).FindListByWorkshop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/technology.technology/FindListByWorkshop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnologyServer).FindListByWorkshop(ctx, req.(*FindListByWorkshopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Technology_FindListByEnterprise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindListByEnterpriseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnologyServer).FindListByEnterprise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/technology.technology/FindListByEnterprise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnologyServer).FindListByEnterprise(ctx, req.(*FindListByEnterpriseReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Technology_ServiceDesc is the grpc.ServiceDesc for Technology service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Technology_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "technology.technology",
	HandlerType: (*TechnologyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Technology_Insert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Technology_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Technology_Update_Handler,
		},
		{
			MethodName: "FindOneById",
			Handler:    _Technology_FindOneById_Handler,
		},
		{
			MethodName: "FindListByWorkshop",
			Handler:    _Technology_FindListByWorkshop_Handler,
		},
		{
			MethodName: "FindListByEnterprise",
			Handler:    _Technology_FindListByEnterprise_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "technology.proto",
}
