// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: node/node.proto

package node

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
	Dispersal_StoreChunks_FullMethodName = "/node.Dispersal/StoreChunks"
	Dispersal_NodeInfo_FullMethodName    = "/node.Dispersal/NodeInfo"
)

// DispersalClient is the client API for Dispersal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispersalClient interface {
	// StoreChunks validates that the chunks match what the Node is supposed to receive (
	// different Nodes are responsible for different chunks, as EigenDA is horizontally
	// sharded) and is correctly coded (e.g. each chunk must be a valid KZG multiproof)
	// according to the EigenDA protocol. It also stores the chunks along with metadata
	// for the protocol-defined length of custody. It will return a signature at the
	// end to attest to the data in this request it has processed.
	StoreChunks(ctx context.Context, in *StoreChunksRequest, opts ...grpc.CallOption) (*StoreChunksReply, error)
	// Retrieve node info metadata
	NodeInfo(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoReply, error)
}

type dispersalClient struct {
	cc grpc.ClientConnInterface
}

func NewDispersalClient(cc grpc.ClientConnInterface) DispersalClient {
	return &dispersalClient{cc}
}

func (c *dispersalClient) StoreChunks(ctx context.Context, in *StoreChunksRequest, opts ...grpc.CallOption) (*StoreChunksReply, error) {
	out := new(StoreChunksReply)
	err := c.cc.Invoke(ctx, Dispersal_StoreChunks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispersalClient) NodeInfo(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoReply, error) {
	out := new(NodeInfoReply)
	err := c.cc.Invoke(ctx, Dispersal_NodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispersalServer is the server API for Dispersal service.
// All implementations must embed UnimplementedDispersalServer
// for forward compatibility
type DispersalServer interface {
	// StoreChunks validates that the chunks match what the Node is supposed to receive (
	// different Nodes are responsible for different chunks, as EigenDA is horizontally
	// sharded) and is correctly coded (e.g. each chunk must be a valid KZG multiproof)
	// according to the EigenDA protocol. It also stores the chunks along with metadata
	// for the protocol-defined length of custody. It will return a signature at the
	// end to attest to the data in this request it has processed.
	StoreChunks(context.Context, *StoreChunksRequest) (*StoreChunksReply, error)
	// Retrieve node info metadata
	NodeInfo(context.Context, *NodeInfoRequest) (*NodeInfoReply, error)
	mustEmbedUnimplementedDispersalServer()
}

// UnimplementedDispersalServer must be embedded to have forward compatible implementations.
type UnimplementedDispersalServer struct {
}

func (UnimplementedDispersalServer) StoreChunks(context.Context, *StoreChunksRequest) (*StoreChunksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreChunks not implemented")
}
func (UnimplementedDispersalServer) NodeInfo(context.Context, *NodeInfoRequest) (*NodeInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeInfo not implemented")
}
func (UnimplementedDispersalServer) mustEmbedUnimplementedDispersalServer() {}

// UnsafeDispersalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispersalServer will
// result in compilation errors.
type UnsafeDispersalServer interface {
	mustEmbedUnimplementedDispersalServer()
}

func RegisterDispersalServer(s grpc.ServiceRegistrar, srv DispersalServer) {
	s.RegisterService(&Dispersal_ServiceDesc, srv)
}

func _Dispersal_StoreChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreChunksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispersalServer).StoreChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dispersal_StoreChunks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispersalServer).StoreChunks(ctx, req.(*StoreChunksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispersal_NodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispersalServer).NodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dispersal_NodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispersalServer).NodeInfo(ctx, req.(*NodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dispersal_ServiceDesc is the grpc.ServiceDesc for Dispersal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dispersal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.Dispersal",
	HandlerType: (*DispersalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreChunks",
			Handler:    _Dispersal_StoreChunks_Handler,
		},
		{
			MethodName: "NodeInfo",
			Handler:    _Dispersal_NodeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/node.proto",
}

const (
	Retrieval_RetrieveChunks_FullMethodName = "/node.Retrieval/RetrieveChunks"
	Retrieval_GetBlobHeader_FullMethodName  = "/node.Retrieval/GetBlobHeader"
	Retrieval_NodeInfo_FullMethodName       = "/node.Retrieval/NodeInfo"
)

// RetrievalClient is the client API for Retrieval service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RetrievalClient interface {
	// RetrieveChunks retrieves the chunks for a blob custodied at the Node.
	RetrieveChunks(ctx context.Context, in *RetrieveChunksRequest, opts ...grpc.CallOption) (*RetrieveChunksReply, error)
	// GetBlobHeader is similar to RetrieveChunks, this just returns the header of the blob.
	GetBlobHeader(ctx context.Context, in *GetBlobHeaderRequest, opts ...grpc.CallOption) (*GetBlobHeaderReply, error)
	// Retrieve node info metadata
	NodeInfo(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoReply, error)
}

type retrievalClient struct {
	cc grpc.ClientConnInterface
}

func NewRetrievalClient(cc grpc.ClientConnInterface) RetrievalClient {
	return &retrievalClient{cc}
}

func (c *retrievalClient) RetrieveChunks(ctx context.Context, in *RetrieveChunksRequest, opts ...grpc.CallOption) (*RetrieveChunksReply, error) {
	out := new(RetrieveChunksReply)
	err := c.cc.Invoke(ctx, Retrieval_RetrieveChunks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalClient) GetBlobHeader(ctx context.Context, in *GetBlobHeaderRequest, opts ...grpc.CallOption) (*GetBlobHeaderReply, error) {
	out := new(GetBlobHeaderReply)
	err := c.cc.Invoke(ctx, Retrieval_GetBlobHeader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalClient) NodeInfo(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoReply, error) {
	out := new(NodeInfoReply)
	err := c.cc.Invoke(ctx, Retrieval_NodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetrievalServer is the server API for Retrieval service.
// All implementations must embed UnimplementedRetrievalServer
// for forward compatibility
type RetrievalServer interface {
	// RetrieveChunks retrieves the chunks for a blob custodied at the Node.
	RetrieveChunks(context.Context, *RetrieveChunksRequest) (*RetrieveChunksReply, error)
	// GetBlobHeader is similar to RetrieveChunks, this just returns the header of the blob.
	GetBlobHeader(context.Context, *GetBlobHeaderRequest) (*GetBlobHeaderReply, error)
	// Retrieve node info metadata
	NodeInfo(context.Context, *NodeInfoRequest) (*NodeInfoReply, error)
	mustEmbedUnimplementedRetrievalServer()
}

// UnimplementedRetrievalServer must be embedded to have forward compatible implementations.
type UnimplementedRetrievalServer struct {
}

func (UnimplementedRetrievalServer) RetrieveChunks(context.Context, *RetrieveChunksRequest) (*RetrieveChunksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveChunks not implemented")
}
func (UnimplementedRetrievalServer) GetBlobHeader(context.Context, *GetBlobHeaderRequest) (*GetBlobHeaderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlobHeader not implemented")
}
func (UnimplementedRetrievalServer) NodeInfo(context.Context, *NodeInfoRequest) (*NodeInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeInfo not implemented")
}
func (UnimplementedRetrievalServer) mustEmbedUnimplementedRetrievalServer() {}

// UnsafeRetrievalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RetrievalServer will
// result in compilation errors.
type UnsafeRetrievalServer interface {
	mustEmbedUnimplementedRetrievalServer()
}

func RegisterRetrievalServer(s grpc.ServiceRegistrar, srv RetrievalServer) {
	s.RegisterService(&Retrieval_ServiceDesc, srv)
}

func _Retrieval_RetrieveChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveChunksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServer).RetrieveChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Retrieval_RetrieveChunks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServer).RetrieveChunks(ctx, req.(*RetrieveChunksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieval_GetBlobHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlobHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServer).GetBlobHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Retrieval_GetBlobHeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServer).GetBlobHeader(ctx, req.(*GetBlobHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieval_NodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServer).NodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Retrieval_NodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServer).NodeInfo(ctx, req.(*NodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Retrieval_ServiceDesc is the grpc.ServiceDesc for Retrieval service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Retrieval_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.Retrieval",
	HandlerType: (*RetrievalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveChunks",
			Handler:    _Retrieval_RetrieveChunks_Handler,
		},
		{
			MethodName: "GetBlobHeader",
			Handler:    _Retrieval_GetBlobHeader_Handler,
		},
		{
			MethodName: "NodeInfo",
			Handler:    _Retrieval_NodeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/node.proto",
}
