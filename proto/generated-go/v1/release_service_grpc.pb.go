// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: v1/release_service.proto

package v1

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
	ReleaseService_GetRelease_FullMethodName    = "/bytebase.v1.ReleaseService/GetRelease"
	ReleaseService_ListReleases_FullMethodName  = "/bytebase.v1.ReleaseService/ListReleases"
	ReleaseService_CreateRelease_FullMethodName = "/bytebase.v1.ReleaseService/CreateRelease"
	ReleaseService_UpdateRelease_FullMethodName = "/bytebase.v1.ReleaseService/UpdateRelease"
)

// ReleaseServiceClient is the client API for ReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReleaseServiceClient interface {
	GetRelease(ctx context.Context, in *GetReleaseRequest, opts ...grpc.CallOption) (*Release, error)
	ListReleases(ctx context.Context, in *ListReleasesRequest, opts ...grpc.CallOption) (*ListReleasesResponse, error)
	CreateRelease(ctx context.Context, in *CreateReleaseRequest, opts ...grpc.CallOption) (*Release, error)
	UpdateRelease(ctx context.Context, in *UpdateReleaseRequest, opts ...grpc.CallOption) (*Release, error)
}

type releaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReleaseServiceClient(cc grpc.ClientConnInterface) ReleaseServiceClient {
	return &releaseServiceClient{cc}
}

func (c *releaseServiceClient) GetRelease(ctx context.Context, in *GetReleaseRequest, opts ...grpc.CallOption) (*Release, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Release)
	err := c.cc.Invoke(ctx, ReleaseService_GetRelease_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) ListReleases(ctx context.Context, in *ListReleasesRequest, opts ...grpc.CallOption) (*ListReleasesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListReleasesResponse)
	err := c.cc.Invoke(ctx, ReleaseService_ListReleases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) CreateRelease(ctx context.Context, in *CreateReleaseRequest, opts ...grpc.CallOption) (*Release, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Release)
	err := c.cc.Invoke(ctx, ReleaseService_CreateRelease_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) UpdateRelease(ctx context.Context, in *UpdateReleaseRequest, opts ...grpc.CallOption) (*Release, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Release)
	err := c.cc.Invoke(ctx, ReleaseService_UpdateRelease_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReleaseServiceServer is the server API for ReleaseService service.
// All implementations must embed UnimplementedReleaseServiceServer
// for forward compatibility.
type ReleaseServiceServer interface {
	GetRelease(context.Context, *GetReleaseRequest) (*Release, error)
	ListReleases(context.Context, *ListReleasesRequest) (*ListReleasesResponse, error)
	CreateRelease(context.Context, *CreateReleaseRequest) (*Release, error)
	UpdateRelease(context.Context, *UpdateReleaseRequest) (*Release, error)
	mustEmbedUnimplementedReleaseServiceServer()
}

// UnimplementedReleaseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReleaseServiceServer struct{}

func (UnimplementedReleaseServiceServer) GetRelease(context.Context, *GetReleaseRequest) (*Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelease not implemented")
}
func (UnimplementedReleaseServiceServer) ListReleases(context.Context, *ListReleasesRequest) (*ListReleasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReleases not implemented")
}
func (UnimplementedReleaseServiceServer) CreateRelease(context.Context, *CreateReleaseRequest) (*Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelease not implemented")
}
func (UnimplementedReleaseServiceServer) UpdateRelease(context.Context, *UpdateReleaseRequest) (*Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRelease not implemented")
}
func (UnimplementedReleaseServiceServer) mustEmbedUnimplementedReleaseServiceServer() {}
func (UnimplementedReleaseServiceServer) testEmbeddedByValue()                        {}

// UnsafeReleaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReleaseServiceServer will
// result in compilation errors.
type UnsafeReleaseServiceServer interface {
	mustEmbedUnimplementedReleaseServiceServer()
}

func RegisterReleaseServiceServer(s grpc.ServiceRegistrar, srv ReleaseServiceServer) {
	// If the following call pancis, it indicates UnimplementedReleaseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReleaseService_ServiceDesc, srv)
}

func _ReleaseService_GetRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).GetRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReleaseService_GetRelease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).GetRelease(ctx, req.(*GetReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_ListReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).ListReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReleaseService_ListReleases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).ListReleases(ctx, req.(*ListReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_CreateRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).CreateRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReleaseService_CreateRelease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).CreateRelease(ctx, req.(*CreateReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_UpdateRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).UpdateRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReleaseService_UpdateRelease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).UpdateRelease(ctx, req.(*UpdateReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReleaseService_ServiceDesc is the grpc.ServiceDesc for ReleaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReleaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.ReleaseService",
	HandlerType: (*ReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRelease",
			Handler:    _ReleaseService_GetRelease_Handler,
		},
		{
			MethodName: "ListReleases",
			Handler:    _ReleaseService_ListReleases_Handler,
		},
		{
			MethodName: "CreateRelease",
			Handler:    _ReleaseService_CreateRelease_Handler,
		},
		{
			MethodName: "UpdateRelease",
			Handler:    _ReleaseService_UpdateRelease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/release_service.proto",
}
