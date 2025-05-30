// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: pkg/proto/url.proto

package crawlerv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UrlService_CreateUrl_FullMethodName = "/crawler.v1.UrlService/CreateUrl"
	UrlService_GetUrls_FullMethodName   = "/crawler.v1.UrlService/GetUrls"
	UrlService_UpdateUrl_FullMethodName = "/crawler.v1.UrlService/UpdateUrl"
)

// UrlServiceClient is the client API for UrlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlServiceClient interface {
	CreateUrl(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlResponse, error)
	GetUrls(ctx context.Context, in *GetUrlsRequest, opts ...grpc.CallOption) (*GetUrlsResponse, error)
	UpdateUrl(ctx context.Context, in *UpdateUrlRequest, opts ...grpc.CallOption) (*UpdateUrlResponse, error)
}

type urlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlServiceClient(cc grpc.ClientConnInterface) UrlServiceClient {
	return &urlServiceClient{cc}
}

func (c *urlServiceClient) CreateUrl(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUrlResponse)
	err := c.cc.Invoke(ctx, UrlService_CreateUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) GetUrls(ctx context.Context, in *GetUrlsRequest, opts ...grpc.CallOption) (*GetUrlsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUrlsResponse)
	err := c.cc.Invoke(ctx, UrlService_GetUrls_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) UpdateUrl(ctx context.Context, in *UpdateUrlRequest, opts ...grpc.CallOption) (*UpdateUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUrlResponse)
	err := c.cc.Invoke(ctx, UrlService_UpdateUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlServiceServer is the server API for UrlService service.
// All implementations must embed UnimplementedUrlServiceServer
// for forward compatibility.
type UrlServiceServer interface {
	CreateUrl(context.Context, *CreateUrlRequest) (*CreateUrlResponse, error)
	GetUrls(context.Context, *GetUrlsRequest) (*GetUrlsResponse, error)
	UpdateUrl(context.Context, *UpdateUrlRequest) (*UpdateUrlResponse, error)
	mustEmbedUnimplementedUrlServiceServer()
}

// UnimplementedUrlServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUrlServiceServer struct{}

func (UnimplementedUrlServiceServer) CreateUrl(context.Context, *CreateUrlRequest) (*CreateUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUrl not implemented")
}
func (UnimplementedUrlServiceServer) GetUrls(context.Context, *GetUrlsRequest) (*GetUrlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrls not implemented")
}
func (UnimplementedUrlServiceServer) UpdateUrl(context.Context, *UpdateUrlRequest) (*UpdateUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUrl not implemented")
}
func (UnimplementedUrlServiceServer) mustEmbedUnimplementedUrlServiceServer() {}
func (UnimplementedUrlServiceServer) testEmbeddedByValue()                    {}

// UnsafeUrlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlServiceServer will
// result in compilation errors.
type UnsafeUrlServiceServer interface {
	mustEmbedUnimplementedUrlServiceServer()
}

func RegisterUrlServiceServer(s grpc.ServiceRegistrar, srv UrlServiceServer) {
	// If the following call pancis, it indicates UnimplementedUrlServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UrlService_ServiceDesc, srv)
}

func _UrlService_CreateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).CreateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlService_CreateUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).CreateUrl(ctx, req.(*CreateUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_GetUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).GetUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlService_GetUrls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).GetUrls(ctx, req.(*GetUrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_UpdateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).UpdateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlService_UpdateUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).UpdateUrl(ctx, req.(*UpdateUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlService_ServiceDesc is the grpc.ServiceDesc for UrlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crawler.v1.UrlService",
	HandlerType: (*UrlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUrl",
			Handler:    _UrlService_CreateUrl_Handler,
		},
		{
			MethodName: "GetUrls",
			Handler:    _UrlService_GetUrls_Handler,
		},
		{
			MethodName: "UpdateUrl",
			Handler:    _UrlService_UpdateUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/url.proto",
}
