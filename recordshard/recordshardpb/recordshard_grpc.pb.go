// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package recordshardpb

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

// RecordShardClient is the client API for RecordShard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordShardClient interface {
	Append(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*CommittedRecord, error)
	Subscribe(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (RecordShard_SubscribeClient, error)
}

type recordShardClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordShardClient(cc grpc.ClientConnInterface) RecordShardClient {
	return &recordShardClient{cc}
}

func (c *recordShardClient) Append(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*CommittedRecord, error) {
	out := new(CommittedRecord)
	err := c.cc.Invoke(ctx, "/recordshardpb.RecordShard/Append", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordShardClient) Subscribe(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (RecordShard_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &RecordShard_ServiceDesc.Streams[0], "/recordshardpb.RecordShard/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordShardSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecordShard_SubscribeClient interface {
	Recv() (*CommittedRecord, error)
	grpc.ClientStream
}

type recordShardSubscribeClient struct {
	grpc.ClientStream
}

func (x *recordShardSubscribeClient) Recv() (*CommittedRecord, error) {
	m := new(CommittedRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecordShardServer is the server API for RecordShard service.
// All implementations must embed UnimplementedRecordShardServer
// for forward compatibility
type RecordShardServer interface {
	Append(context.Context, *AppendRequest) (*CommittedRecord, error)
	Subscribe(*ReadRequest, RecordShard_SubscribeServer) error
	mustEmbedUnimplementedRecordShardServer()
}

// UnimplementedRecordShardServer must be embedded to have forward compatible implementations.
type UnimplementedRecordShardServer struct {
}

func (UnimplementedRecordShardServer) Append(context.Context, *AppendRequest) (*CommittedRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Append not implemented")
}
func (UnimplementedRecordShardServer) Subscribe(*ReadRequest, RecordShard_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedRecordShardServer) mustEmbedUnimplementedRecordShardServer() {}

// UnsafeRecordShardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordShardServer will
// result in compilation errors.
type UnsafeRecordShardServer interface {
	mustEmbedUnimplementedRecordShardServer()
}

func RegisterRecordShardServer(s grpc.ServiceRegistrar, srv RecordShardServer) {
	s.RegisterService(&RecordShard_ServiceDesc, srv)
}

func _RecordShard_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordShardServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recordshardpb.RecordShard/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordShardServer).Append(ctx, req.(*AppendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordShard_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordShardServer).Subscribe(m, &recordShardSubscribeServer{stream})
}

type RecordShard_SubscribeServer interface {
	Send(*CommittedRecord) error
	grpc.ServerStream
}

type recordShardSubscribeServer struct {
	grpc.ServerStream
}

func (x *recordShardSubscribeServer) Send(m *CommittedRecord) error {
	return x.ServerStream.SendMsg(m)
}

// RecordShard_ServiceDesc is the grpc.ServiceDesc for RecordShard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordShard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recordshardpb.RecordShard",
	HandlerType: (*RecordShardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Append",
			Handler:    _RecordShard_Append_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _RecordShard_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "recordshard.proto",
}
