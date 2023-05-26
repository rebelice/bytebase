// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/sql_service.proto

package v1

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

const (
	SQLService_Pretty_FullMethodName = "/bytebase.v1.SQLService/Pretty"
	SQLService_Query_FullMethodName  = "/bytebase.v1.SQLService/Query"
)

// SQLServiceClient is the client API for SQLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SQLServiceClient interface {
	Pretty(ctx context.Context, in *PrettyRequest, opts ...grpc.CallOption) (*PrettyResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type sQLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSQLServiceClient(cc grpc.ClientConnInterface) SQLServiceClient {
	return &sQLServiceClient{cc}
}

func (c *sQLServiceClient) Pretty(ctx context.Context, in *PrettyRequest, opts ...grpc.CallOption) (*PrettyResponse, error) {
	out := new(PrettyResponse)
	err := c.cc.Invoke(ctx, SQLService_Pretty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, SQLService_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SQLServiceServer is the server API for SQLService service.
// All implementations must embed UnimplementedSQLServiceServer
// for forward compatibility
type SQLServiceServer interface {
	Pretty(context.Context, *PrettyRequest) (*PrettyResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	mustEmbedUnimplementedSQLServiceServer()
}

// UnimplementedSQLServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSQLServiceServer struct {
}

func (UnimplementedSQLServiceServer) Pretty(context.Context, *PrettyRequest) (*PrettyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pretty not implemented")
}
func (UnimplementedSQLServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedSQLServiceServer) mustEmbedUnimplementedSQLServiceServer() {}

// UnsafeSQLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SQLServiceServer will
// result in compilation errors.
type UnsafeSQLServiceServer interface {
	mustEmbedUnimplementedSQLServiceServer()
}

func RegisterSQLServiceServer(s grpc.ServiceRegistrar, srv SQLServiceServer) {
	s.RegisterService(&SQLService_ServiceDesc, srv)
}

func _SQLService_Pretty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrettyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLServiceServer).Pretty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLService_Pretty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLServiceServer).Pretty(ctx, req.(*PrettyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SQLService_ServiceDesc is the grpc.ServiceDesc for SQLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SQLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.SQLService",
	HandlerType: (*SQLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pretty",
			Handler:    _SQLService_Pretty_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _SQLService_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/sql_service.proto",
}
