// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/product.proto

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductMsg, error)
	FetchOneProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*ProductReq, error)
	FetchAllProduct(ctx context.Context, in *ProductEmpty, opts ...grpc.CallOption) (ProductService_FetchAllProductClient, error)
	UpdateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductMsg, error)
	DeleteProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*ProductMsg, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductMsg, error) {
	out := new(ProductMsg)
	err := c.cc.Invoke(ctx, "/grpc.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FetchOneProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*ProductReq, error) {
	out := new(ProductReq)
	err := c.cc.Invoke(ctx, "/grpc.ProductService/FetchOneProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FetchAllProduct(ctx context.Context, in *ProductEmpty, opts ...grpc.CallOption) (ProductService_FetchAllProductClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], "/grpc.ProductService/FetchAllProduct", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceFetchAllProductClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_FetchAllProductClient interface {
	Recv() (*ProductReq, error)
	grpc.ClientStream
}

type productServiceFetchAllProductClient struct {
	grpc.ClientStream
}

func (x *productServiceFetchAllProductClient) Recv() (*ProductReq, error) {
	m := new(ProductReq)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *ProductReq, opts ...grpc.CallOption) (*ProductMsg, error) {
	out := new(ProductMsg)
	err := c.cc.Invoke(ctx, "/grpc.ProductService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*ProductMsg, error) {
	out := new(ProductMsg)
	err := c.cc.Invoke(ctx, "/grpc.ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	CreateProduct(context.Context, *ProductReq) (*ProductMsg, error)
	FetchOneProduct(context.Context, *ProductId) (*ProductReq, error)
	FetchAllProduct(*ProductEmpty, ProductService_FetchAllProductServer) error
	UpdateProduct(context.Context, *ProductReq) (*ProductMsg, error)
	DeleteProduct(context.Context, *ProductId) (*ProductMsg, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) CreateProduct(context.Context, *ProductReq) (*ProductMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) FetchOneProduct(context.Context, *ProductId) (*ProductReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchOneProduct not implemented")
}
func (UnimplementedProductServiceServer) FetchAllProduct(*ProductEmpty, ProductService_FetchAllProductServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchAllProduct not implemented")
}
func (UnimplementedProductServiceServer) UpdateProduct(context.Context, *ProductReq) (*ProductMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *ProductId) (*ProductMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FetchOneProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FetchOneProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ProductService/FetchOneProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FetchOneProduct(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FetchAllProduct_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).FetchAllProduct(m, &productServiceFetchAllProductServer{stream})
}

type ProductService_FetchAllProductServer interface {
	Send(*ProductReq) error
	grpc.ServerStream
}

type productServiceFetchAllProductServer struct {
	grpc.ServerStream
}

func (x *productServiceFetchAllProductServer) Send(m *ProductReq) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ProductService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*ProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "FetchOneProduct",
			Handler:    _ProductService_FetchOneProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchAllProduct",
			Handler:       _ProductService_FetchAllProduct_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/product.proto",
}
