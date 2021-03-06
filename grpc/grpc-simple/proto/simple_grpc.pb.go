// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/simple.proto

package proto

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

// ParserClient is the client API for Parser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParserClient interface {
	ParseAudit(ctx context.Context, in *ParseMessageRequest, opts ...grpc.CallOption) (*ParseMessageResponse, error)
}

type parserClient struct {
	cc grpc.ClientConnInterface
}

func NewParserClient(cc grpc.ClientConnInterface) ParserClient {
	return &parserClient{cc}
}

func (c *parserClient) ParseAudit(ctx context.Context, in *ParseMessageRequest, opts ...grpc.CallOption) (*ParseMessageResponse, error) {
	out := new(ParseMessageResponse)
	err := c.cc.Invoke(ctx, "/grpc_simple.Parser/ParseAudit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParserServer is the server API for Parser service.
// All implementations must embed UnimplementedParserServer
// for forward compatibility
type ParserServer interface {
	ParseAudit(context.Context, *ParseMessageRequest) (*ParseMessageResponse, error)
	mustEmbedUnimplementedParserServer()
}

// UnimplementedParserServer must be embedded to have forward compatible implementations.
type UnimplementedParserServer struct {
}

func (UnimplementedParserServer) ParseAudit(context.Context, *ParseMessageRequest) (*ParseMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseAudit not implemented")
}
func (UnimplementedParserServer) mustEmbedUnimplementedParserServer() {}

// UnsafeParserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParserServer will
// result in compilation errors.
type UnsafeParserServer interface {
	mustEmbedUnimplementedParserServer()
}

func RegisterParserServer(s grpc.ServiceRegistrar, srv ParserServer) {
	s.RegisterService(&Parser_ServiceDesc, srv)
}

func _Parser_ParseAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserServer).ParseAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_simple.Parser/ParseAudit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserServer).ParseAudit(ctx, req.(*ParseMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Parser_ServiceDesc is the grpc.ServiceDesc for Parser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Parser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_simple.Parser",
	HandlerType: (*ParserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseAudit",
			Handler:    _Parser_ParseAudit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/simple.proto",
}
