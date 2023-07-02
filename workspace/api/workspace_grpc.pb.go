// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package workspace

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

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error)
	GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error)
	UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*UpdateWorkspaceResponse, error)
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error)
	ListWorkspaceMembers(ctx context.Context, in *ListWorkspaceMembersRequest, opts ...grpc.CallOption) (*ListWorkspaceMembersResponse, error)
	InviteWorkspaceMember(ctx context.Context, in *InviteWorkspaceMemberRequest, opts ...grpc.CallOption) (*InviteWorkspaceMemberResponse, error)
	RemoveWorkspaceMember(ctx context.Context, in *RemoveWorkspaceMemberRequest, opts ...grpc.CallOption) (*RemoveWorkspaceMemberResponse, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error) {
	out := new(CreateWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/CreateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error) {
	out := new(GetWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/GetWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*UpdateWorkspaceResponse, error) {
	out := new(UpdateWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/UpdateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error) {
	out := new(DeleteWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/DeleteWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ListWorkspaceMembers(ctx context.Context, in *ListWorkspaceMembersRequest, opts ...grpc.CallOption) (*ListWorkspaceMembersResponse, error) {
	out := new(ListWorkspaceMembersResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/ListWorkspaceMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) InviteWorkspaceMember(ctx context.Context, in *InviteWorkspaceMemberRequest, opts ...grpc.CallOption) (*InviteWorkspaceMemberResponse, error) {
	out := new(InviteWorkspaceMemberResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/InviteWorkspaceMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) RemoveWorkspaceMember(ctx context.Context, in *RemoveWorkspaceMemberRequest, opts ...grpc.CallOption) (*RemoveWorkspaceMemberResponse, error) {
	out := new(RemoveWorkspaceMemberResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceService/RemoveWorkspaceMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
// All implementations must embed UnimplementedWorkspaceServiceServer
// for forward compatibility
type WorkspaceServiceServer interface {
	CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error)
	GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error)
	UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*UpdateWorkspaceResponse, error)
	DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error)
	ListWorkspaceMembers(context.Context, *ListWorkspaceMembersRequest) (*ListWorkspaceMembersResponse, error)
	InviteWorkspaceMember(context.Context, *InviteWorkspaceMemberRequest) (*InviteWorkspaceMemberResponse, error)
	RemoveWorkspaceMember(context.Context, *RemoveWorkspaceMemberRequest) (*RemoveWorkspaceMemberResponse, error)
	mustEmbedUnimplementedWorkspaceServiceServer()
}

// UnimplementedWorkspaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (UnimplementedWorkspaceServiceServer) CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*UpdateWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListWorkspaceMembers(context.Context, *ListWorkspaceMembersRequest) (*ListWorkspaceMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaceMembers not implemented")
}
func (UnimplementedWorkspaceServiceServer) InviteWorkspaceMember(context.Context, *InviteWorkspaceMemberRequest) (*InviteWorkspaceMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteWorkspaceMember not implemented")
}
func (UnimplementedWorkspaceServiceServer) RemoveWorkspaceMember(context.Context, *RemoveWorkspaceMemberRequest) (*RemoveWorkspaceMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWorkspaceMember not implemented")
}
func (UnimplementedWorkspaceServiceServer) mustEmbedUnimplementedWorkspaceServiceServer() {}

// UnsafeWorkspaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServiceServer will
// result in compilation errors.
type UnsafeWorkspaceServiceServer interface {
	mustEmbedUnimplementedWorkspaceServiceServer()
}

func RegisterWorkspaceServiceServer(s grpc.ServiceRegistrar, srv WorkspaceServiceServer) {
	s.RegisterService(&WorkspaceService_ServiceDesc, srv)
}

func _WorkspaceService_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/CreateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, req.(*CreateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/GetWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspace(ctx, req.(*GetWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/UpdateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, req.(*UpdateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/DeleteWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, req.(*DeleteWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ListWorkspaceMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListWorkspaceMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/ListWorkspaceMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListWorkspaceMembers(ctx, req.(*ListWorkspaceMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_InviteWorkspaceMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteWorkspaceMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).InviteWorkspaceMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/InviteWorkspaceMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).InviteWorkspaceMember(ctx, req.(*InviteWorkspaceMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_RemoveWorkspaceMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveWorkspaceMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).RemoveWorkspaceMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceService/RemoveWorkspaceMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).RemoveWorkspaceMember(ctx, req.(*RemoveWorkspaceMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspaceService_ServiceDesc is the grpc.ServiceDesc for WorkspaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkspace",
			Handler:    _WorkspaceService_CreateWorkspace_Handler,
		},
		{
			MethodName: "GetWorkspace",
			Handler:    _WorkspaceService_GetWorkspace_Handler,
		},
		{
			MethodName: "UpdateWorkspace",
			Handler:    _WorkspaceService_UpdateWorkspace_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _WorkspaceService_DeleteWorkspace_Handler,
		},
		{
			MethodName: "ListWorkspaceMembers",
			Handler:    _WorkspaceService_ListWorkspaceMembers_Handler,
		},
		{
			MethodName: "InviteWorkspaceMember",
			Handler:    _WorkspaceService_InviteWorkspaceMember_Handler,
		},
		{
			MethodName: "RemoveWorkspaceMember",
			Handler:    _WorkspaceService_RemoveWorkspaceMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace.proto",
}