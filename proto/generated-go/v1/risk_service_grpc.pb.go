// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/risk_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RiskService_ListRisks_FullMethodName  = "/bytebase.v1.RiskService/ListRisks"
	RiskService_CreateRisk_FullMethodName = "/bytebase.v1.RiskService/CreateRisk"
	RiskService_UpdateRisk_FullMethodName = "/bytebase.v1.RiskService/UpdateRisk"
	RiskService_DeleteRisk_FullMethodName = "/bytebase.v1.RiskService/DeleteRisk"
)

// RiskServiceClient is the client API for RiskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RiskServiceClient interface {
	ListRisks(ctx context.Context, in *ListRisksRequest, opts ...grpc.CallOption) (*ListRisksResponse, error)
	CreateRisk(ctx context.Context, in *CreateRiskRequest, opts ...grpc.CallOption) (*Risk, error)
	UpdateRisk(ctx context.Context, in *UpdateRiskRequest, opts ...grpc.CallOption) (*Risk, error)
	DeleteRisk(ctx context.Context, in *DeleteRiskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type riskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRiskServiceClient(cc grpc.ClientConnInterface) RiskServiceClient {
	return &riskServiceClient{cc}
}

func (c *riskServiceClient) ListRisks(ctx context.Context, in *ListRisksRequest, opts ...grpc.CallOption) (*ListRisksResponse, error) {
	out := new(ListRisksResponse)
	err := c.cc.Invoke(ctx, RiskService_ListRisks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riskServiceClient) CreateRisk(ctx context.Context, in *CreateRiskRequest, opts ...grpc.CallOption) (*Risk, error) {
	out := new(Risk)
	err := c.cc.Invoke(ctx, RiskService_CreateRisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riskServiceClient) UpdateRisk(ctx context.Context, in *UpdateRiskRequest, opts ...grpc.CallOption) (*Risk, error) {
	out := new(Risk)
	err := c.cc.Invoke(ctx, RiskService_UpdateRisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riskServiceClient) DeleteRisk(ctx context.Context, in *DeleteRiskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RiskService_DeleteRisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RiskServiceServer is the server API for RiskService service.
// All implementations must embed UnimplementedRiskServiceServer
// for forward compatibility
type RiskServiceServer interface {
	ListRisks(context.Context, *ListRisksRequest) (*ListRisksResponse, error)
	CreateRisk(context.Context, *CreateRiskRequest) (*Risk, error)
	UpdateRisk(context.Context, *UpdateRiskRequest) (*Risk, error)
	DeleteRisk(context.Context, *DeleteRiskRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRiskServiceServer()
}

// UnimplementedRiskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRiskServiceServer struct {
}

func (UnimplementedRiskServiceServer) ListRisks(context.Context, *ListRisksRequest) (*ListRisksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRisks not implemented")
}
func (UnimplementedRiskServiceServer) CreateRisk(context.Context, *CreateRiskRequest) (*Risk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRisk not implemented")
}
func (UnimplementedRiskServiceServer) UpdateRisk(context.Context, *UpdateRiskRequest) (*Risk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRisk not implemented")
}
func (UnimplementedRiskServiceServer) DeleteRisk(context.Context, *DeleteRiskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRisk not implemented")
}
func (UnimplementedRiskServiceServer) mustEmbedUnimplementedRiskServiceServer() {}

// UnsafeRiskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RiskServiceServer will
// result in compilation errors.
type UnsafeRiskServiceServer interface {
	mustEmbedUnimplementedRiskServiceServer()
}

func RegisterRiskServiceServer(s grpc.ServiceRegistrar, srv RiskServiceServer) {
	s.RegisterService(&RiskService_ServiceDesc, srv)
}

func _RiskService_ListRisks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRisksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskServiceServer).ListRisks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiskService_ListRisks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskServiceServer).ListRisks(ctx, req.(*ListRisksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiskService_CreateRisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskServiceServer).CreateRisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiskService_CreateRisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskServiceServer).CreateRisk(ctx, req.(*CreateRiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiskService_UpdateRisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskServiceServer).UpdateRisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiskService_UpdateRisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskServiceServer).UpdateRisk(ctx, req.(*UpdateRiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiskService_DeleteRisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskServiceServer).DeleteRisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiskService_DeleteRisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskServiceServer).DeleteRisk(ctx, req.(*DeleteRiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RiskService_ServiceDesc is the grpc.ServiceDesc for RiskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RiskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.RiskService",
	HandlerType: (*RiskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRisks",
			Handler:    _RiskService_ListRisks_Handler,
		},
		{
			MethodName: "CreateRisk",
			Handler:    _RiskService_CreateRisk_Handler,
		},
		{
			MethodName: "UpdateRisk",
			Handler:    _RiskService_UpdateRisk_Handler,
		},
		{
			MethodName: "DeleteRisk",
			Handler:    _RiskService_DeleteRisk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/risk_service.proto",
}
