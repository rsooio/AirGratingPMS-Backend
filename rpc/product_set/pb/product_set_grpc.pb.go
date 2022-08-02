// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: product_set.proto

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

// ProductsetClient is the client API for Productset service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsetClient interface {
	Insert(ctx context.Context, in *ProductSetInfo, opts ...grpc.CallOption) (*InsertResp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *ProductSetInfo, opts ...grpc.CallOption) (*Empty, error)
	// rpc CustomUpdate (ProductSetInfo) returns (Empty);
	FindOneById(ctx context.Context, in *FindOneByIdReq, opts ...grpc.CallOption) (*ProductSetInfo, error)
	FindListByOrder(ctx context.Context, in *FindListByOrderReq, opts ...grpc.CallOption) (*ProductSetList, error)
}

type productsetClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsetClient(cc grpc.ClientConnInterface) ProductsetClient {
	return &productsetClient{cc}
}

func (c *productsetClient) Insert(ctx context.Context, in *ProductSetInfo, opts ...grpc.CallOption) (*InsertResp, error) {
	out := new(InsertResp)
	err := c.cc.Invoke(ctx, "/productset.productset/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsetClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/productset.productset/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsetClient) Update(ctx context.Context, in *ProductSetInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/productset.productset/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsetClient) FindOneById(ctx context.Context, in *FindOneByIdReq, opts ...grpc.CallOption) (*ProductSetInfo, error) {
	out := new(ProductSetInfo)
	err := c.cc.Invoke(ctx, "/productset.productset/FindOneById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsetClient) FindListByOrder(ctx context.Context, in *FindListByOrderReq, opts ...grpc.CallOption) (*ProductSetList, error) {
	out := new(ProductSetList)
	err := c.cc.Invoke(ctx, "/productset.productset/FindListByOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsetServer is the server API for Productset service.
// All implementations must embed UnimplementedProductsetServer
// for forward compatibility
type ProductsetServer interface {
	Insert(context.Context, *ProductSetInfo) (*InsertResp, error)
	Delete(context.Context, *DeleteReq) (*Empty, error)
	Update(context.Context, *ProductSetInfo) (*Empty, error)
	// rpc CustomUpdate (ProductSetInfo) returns (Empty);
	FindOneById(context.Context, *FindOneByIdReq) (*ProductSetInfo, error)
	FindListByOrder(context.Context, *FindListByOrderReq) (*ProductSetList, error)
	mustEmbedUnimplementedProductsetServer()
}

// UnimplementedProductsetServer must be embedded to have forward compatible implementations.
type UnimplementedProductsetServer struct {
}

func (UnimplementedProductsetServer) Insert(context.Context, *ProductSetInfo) (*InsertResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedProductsetServer) Delete(context.Context, *DeleteReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProductsetServer) Update(context.Context, *ProductSetInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductsetServer) FindOneById(context.Context, *FindOneByIdReq) (*ProductSetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneById not implemented")
}
func (UnimplementedProductsetServer) FindListByOrder(context.Context, *FindListByOrderReq) (*ProductSetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListByOrder not implemented")
}
func (UnimplementedProductsetServer) mustEmbedUnimplementedProductsetServer() {}

// UnsafeProductsetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsetServer will
// result in compilation errors.
type UnsafeProductsetServer interface {
	mustEmbedUnimplementedProductsetServer()
}

func RegisterProductsetServer(s grpc.ServiceRegistrar, srv ProductsetServer) {
	s.RegisterService(&Productset_ServiceDesc, srv)
}

func _Productset_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSetInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsetServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productset.productset/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsetServer).Insert(ctx, req.(*ProductSetInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Productset_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsetServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productset.productset/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsetServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Productset_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSetInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsetServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productset.productset/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsetServer).Update(ctx, req.(*ProductSetInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Productset_FindOneById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsetServer).FindOneById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productset.productset/FindOneById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsetServer).FindOneById(ctx, req.(*FindOneByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Productset_FindListByOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindListByOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsetServer).FindListByOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productset.productset/FindListByOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsetServer).FindListByOrder(ctx, req.(*FindListByOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Productset_ServiceDesc is the grpc.ServiceDesc for Productset service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Productset_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "productset.productset",
	HandlerType: (*ProductsetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Productset_Insert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Productset_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Productset_Update_Handler,
		},
		{
			MethodName: "FindOneById",
			Handler:    _Productset_FindOneById_Handler,
		},
		{
			MethodName: "FindListByOrder",
			Handler:    _Productset_FindListByOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_set.proto",
}
