// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: finance.proto

package finance

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
	FinanceService_GetQuoteData_FullMethodName = "/finance.FinanceService/getQuoteData"
)

// FinanceServiceClient is the client API for FinanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinanceServiceClient interface {
	GetQuoteData(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type financeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinanceServiceClient(cc grpc.ClientConnInterface) FinanceServiceClient {
	return &financeServiceClient{cc}
}

func (c *financeServiceClient) GetQuoteData(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, FinanceService_GetQuoteData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceServiceServer is the server API for FinanceService service.
// All implementations must embed UnimplementedFinanceServiceServer
// for forward compatibility
type FinanceServiceServer interface {
	GetQuoteData(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedFinanceServiceServer()
}

// UnimplementedFinanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFinanceServiceServer struct {
}

func (UnimplementedFinanceServiceServer) GetQuoteData(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuoteData not implemented")
}
func (UnimplementedFinanceServiceServer) mustEmbedUnimplementedFinanceServiceServer() {}

// UnsafeFinanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinanceServiceServer will
// result in compilation errors.
type UnsafeFinanceServiceServer interface {
	mustEmbedUnimplementedFinanceServiceServer()
}

func RegisterFinanceServiceServer(s grpc.ServiceRegistrar, srv FinanceServiceServer) {
	s.RegisterService(&FinanceService_ServiceDesc, srv)
}

func _FinanceService_GetQuoteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServiceServer).GetQuoteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinanceService_GetQuoteData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServiceServer).GetQuoteData(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// FinanceService_ServiceDesc is the grpc.ServiceDesc for FinanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finance.FinanceService",
	HandlerType: (*FinanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getQuoteData",
			Handler:    _FinanceService_GetQuoteData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance.proto",
}
