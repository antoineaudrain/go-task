// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: board.proto

package board

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

// BoardServiceClient is the client API for BoardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardServiceClient interface {
	CreateBoard(ctx context.Context, in *CreateBoardRequest, opts ...grpc.CallOption) (*CreateBoardResponse, error)
	GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error)
	ListBoards(ctx context.Context, in *ListBoardRequest, opts ...grpc.CallOption) (*ListBoardResponse, error)
	UpdateBoard(ctx context.Context, in *UpdateBoardRequest, opts ...grpc.CallOption) (*UpdateBoardResponse, error)
	DeleteBoard(ctx context.Context, in *DeleteBoardRequest, opts ...grpc.CallOption) (*DeleteBoardResponse, error)
}

type boardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardServiceClient(cc grpc.ClientConnInterface) BoardServiceClient {
	return &boardServiceClient{cc}
}

func (c *boardServiceClient) CreateBoard(ctx context.Context, in *CreateBoardRequest, opts ...grpc.CallOption) (*CreateBoardResponse, error) {
	out := new(CreateBoardResponse)
	err := c.cc.Invoke(ctx, "/board.BoardService/CreateBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error) {
	out := new(GetBoardResponse)
	err := c.cc.Invoke(ctx, "/board.BoardService/GetBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) ListBoards(ctx context.Context, in *ListBoardRequest, opts ...grpc.CallOption) (*ListBoardResponse, error) {
	out := new(ListBoardResponse)
	err := c.cc.Invoke(ctx, "/board.BoardService/ListBoards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) UpdateBoard(ctx context.Context, in *UpdateBoardRequest, opts ...grpc.CallOption) (*UpdateBoardResponse, error) {
	out := new(UpdateBoardResponse)
	err := c.cc.Invoke(ctx, "/board.BoardService/UpdateBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) DeleteBoard(ctx context.Context, in *DeleteBoardRequest, opts ...grpc.CallOption) (*DeleteBoardResponse, error) {
	out := new(DeleteBoardResponse)
	err := c.cc.Invoke(ctx, "/board.BoardService/DeleteBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServiceServer is the server API for BoardService service.
// All implementations must embed UnimplementedBoardServiceServer
// for forward compatibility
type BoardServiceServer interface {
	CreateBoard(context.Context, *CreateBoardRequest) (*CreateBoardResponse, error)
	GetBoard(context.Context, *GetBoardRequest) (*GetBoardResponse, error)
	ListBoards(context.Context, *ListBoardRequest) (*ListBoardResponse, error)
	UpdateBoard(context.Context, *UpdateBoardRequest) (*UpdateBoardResponse, error)
	DeleteBoard(context.Context, *DeleteBoardRequest) (*DeleteBoardResponse, error)
	mustEmbedUnimplementedBoardServiceServer()
}

// UnimplementedBoardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBoardServiceServer struct {
}

func (UnimplementedBoardServiceServer) CreateBoard(context.Context, *CreateBoardRequest) (*CreateBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBoard not implemented")
}
func (UnimplementedBoardServiceServer) GetBoard(context.Context, *GetBoardRequest) (*GetBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoard not implemented")
}
func (UnimplementedBoardServiceServer) ListBoards(context.Context, *ListBoardRequest) (*ListBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBoards not implemented")
}
func (UnimplementedBoardServiceServer) UpdateBoard(context.Context, *UpdateBoardRequest) (*UpdateBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBoard not implemented")
}
func (UnimplementedBoardServiceServer) DeleteBoard(context.Context, *DeleteBoardRequest) (*DeleteBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBoard not implemented")
}
func (UnimplementedBoardServiceServer) mustEmbedUnimplementedBoardServiceServer() {}

// UnsafeBoardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServiceServer will
// result in compilation errors.
type UnsafeBoardServiceServer interface {
	mustEmbedUnimplementedBoardServiceServer()
}

func RegisterBoardServiceServer(s grpc.ServiceRegistrar, srv BoardServiceServer) {
	s.RegisterService(&BoardService_ServiceDesc, srv)
}

func _BoardService_CreateBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).CreateBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.BoardService/CreateBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).CreateBoard(ctx, req.(*CreateBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_GetBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).GetBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.BoardService/GetBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).GetBoard(ctx, req.(*GetBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_ListBoards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).ListBoards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.BoardService/ListBoards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).ListBoards(ctx, req.(*ListBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_UpdateBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).UpdateBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.BoardService/UpdateBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).UpdateBoard(ctx, req.(*UpdateBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_DeleteBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).DeleteBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.BoardService/DeleteBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).DeleteBoard(ctx, req.(*DeleteBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoardService_ServiceDesc is the grpc.ServiceDesc for BoardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "board.BoardService",
	HandlerType: (*BoardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBoard",
			Handler:    _BoardService_CreateBoard_Handler,
		},
		{
			MethodName: "GetBoard",
			Handler:    _BoardService_GetBoard_Handler,
		},
		{
			MethodName: "ListBoards",
			Handler:    _BoardService_ListBoards_Handler,
		},
		{
			MethodName: "UpdateBoard",
			Handler:    _BoardService_UpdateBoard_Handler,
		},
		{
			MethodName: "DeleteBoard",
			Handler:    _BoardService_DeleteBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "board.proto",
}