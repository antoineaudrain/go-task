// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: workspace.proto

package workspace

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateWorkspaceRequest) Reset() {
	*x = CreateWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspaceRequest) ProtoMessage() {}

func (x *CreateWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWorkspaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceID string `protobuf:"bytes,1,opt,name=workspaceID,proto3" json:"workspaceID,omitempty"`
}

func (x *GetWorkspaceRequest) Reset() {
	*x = GetWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkspaceRequest) ProtoMessage() {}

func (x *GetWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*GetWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{1}
}

func (x *GetWorkspaceRequest) GetWorkspaceID() string {
	if x != nil {
		return x.WorkspaceID
	}
	return ""
}

type ListWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListWorkspaceRequest) Reset() {
	*x = ListWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkspaceRequest) ProtoMessage() {}

func (x *ListWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*ListWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{2}
}

type UpdateWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceID string `protobuf:"bytes,1,opt,name=workspaceID,proto3" json:"workspaceID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateWorkspaceRequest) Reset() {
	*x = UpdateWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkspaceRequest) ProtoMessage() {}

func (x *UpdateWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateWorkspaceRequest) GetWorkspaceID() string {
	if x != nil {
		return x.WorkspaceID
	}
	return ""
}

func (x *UpdateWorkspaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceID string `protobuf:"bytes,1,opt,name=workspaceID,proto3" json:"workspaceID,omitempty"`
}

func (x *DeleteWorkspaceRequest) Reset() {
	*x = DeleteWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceRequest) ProtoMessage() {}

func (x *DeleteWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteWorkspaceRequest) GetWorkspaceID() string {
	if x != nil {
		return x.WorkspaceID
	}
	return ""
}

type Workspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Workspace) Reset() {
	*x = Workspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace) ProtoMessage() {}

func (x *Workspace) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace.ProtoReflect.Descriptor instead.
func (*Workspace) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{5}
}

func (x *Workspace) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Workspace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workspace) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Workspace) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *CreateWorkspaceResponse) Reset() {
	*x = CreateWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspaceResponse) ProtoMessage() {}

func (x *CreateWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*CreateWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{6}
}

func (x *CreateWorkspaceResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

type GetWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *GetWorkspaceResponse) Reset() {
	*x = GetWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkspaceResponse) ProtoMessage() {}

func (x *GetWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*GetWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{7}
}

func (x *GetWorkspaceResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

type ListWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspaces []*Workspace `protobuf:"bytes,1,rep,name=workspaces,proto3" json:"workspaces,omitempty"`
}

func (x *ListWorkspaceResponse) Reset() {
	*x = ListWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkspaceResponse) ProtoMessage() {}

func (x *ListWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*ListWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{8}
}

func (x *ListWorkspaceResponse) GetWorkspaces() []*Workspace {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

type UpdateWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *UpdateWorkspaceResponse) Reset() {
	*x = UpdateWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkspaceResponse) ProtoMessage() {}

func (x *UpdateWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*UpdateWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateWorkspaceResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

type DeleteWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace *Workspace `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *DeleteWorkspaceResponse) Reset() {
	*x = DeleteWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceResponse) ProtoMessage() {}

func (x *DeleteWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteWorkspaceResponse) GetWorkspace() *Workspace {
	if x != nil {
		return x.Workspace
	}
	return nil
}

var File_workspace_proto protoreflect.FileDescriptor

var file_workspace_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x22, 0x6b, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x4d, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22,
	0x4d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x4d,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x4d, 0x0a,
	0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x32, 0xd0, 0x03, 0x0a,
	0x10, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x1f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x13, 0x5a, 0x11, 0x67, 0x6f, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workspace_proto_rawDescOnce sync.Once
	file_workspace_proto_rawDescData = file_workspace_proto_rawDesc
)

func file_workspace_proto_rawDescGZIP() []byte {
	file_workspace_proto_rawDescOnce.Do(func() {
		file_workspace_proto_rawDescData = protoimpl.X.CompressGZIP(file_workspace_proto_rawDescData)
	})
	return file_workspace_proto_rawDescData
}

var file_workspace_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_workspace_proto_goTypes = []interface{}{
	(*CreateWorkspaceRequest)(nil),  // 0: workspace.CreateWorkspaceRequest
	(*GetWorkspaceRequest)(nil),     // 1: workspace.GetWorkspaceRequest
	(*ListWorkspaceRequest)(nil),    // 2: workspace.ListWorkspaceRequest
	(*UpdateWorkspaceRequest)(nil),  // 3: workspace.UpdateWorkspaceRequest
	(*DeleteWorkspaceRequest)(nil),  // 4: workspace.DeleteWorkspaceRequest
	(*Workspace)(nil),               // 5: workspace.Workspace
	(*CreateWorkspaceResponse)(nil), // 6: workspace.CreateWorkspaceResponse
	(*GetWorkspaceResponse)(nil),    // 7: workspace.GetWorkspaceResponse
	(*ListWorkspaceResponse)(nil),   // 8: workspace.ListWorkspaceResponse
	(*UpdateWorkspaceResponse)(nil), // 9: workspace.UpdateWorkspaceResponse
	(*DeleteWorkspaceResponse)(nil), // 10: workspace.DeleteWorkspaceResponse
}
var file_workspace_proto_depIdxs = []int32{
	5,  // 0: workspace.CreateWorkspaceResponse.workspace:type_name -> workspace.Workspace
	5,  // 1: workspace.GetWorkspaceResponse.workspace:type_name -> workspace.Workspace
	5,  // 2: workspace.ListWorkspaceResponse.workspaces:type_name -> workspace.Workspace
	5,  // 3: workspace.UpdateWorkspaceResponse.workspace:type_name -> workspace.Workspace
	5,  // 4: workspace.DeleteWorkspaceResponse.workspace:type_name -> workspace.Workspace
	0,  // 5: workspace.WorkspaceService.CreateWorkspace:input_type -> workspace.CreateWorkspaceRequest
	1,  // 6: workspace.WorkspaceService.GetWorkspace:input_type -> workspace.GetWorkspaceRequest
	2,  // 7: workspace.WorkspaceService.ListWorkspaces:input_type -> workspace.ListWorkspaceRequest
	3,  // 8: workspace.WorkspaceService.UpdateWorkspace:input_type -> workspace.UpdateWorkspaceRequest
	4,  // 9: workspace.WorkspaceService.DeleteWorkspace:input_type -> workspace.DeleteWorkspaceRequest
	6,  // 10: workspace.WorkspaceService.CreateWorkspace:output_type -> workspace.CreateWorkspaceResponse
	7,  // 11: workspace.WorkspaceService.GetWorkspace:output_type -> workspace.GetWorkspaceResponse
	8,  // 12: workspace.WorkspaceService.ListWorkspaces:output_type -> workspace.ListWorkspaceResponse
	9,  // 13: workspace.WorkspaceService.UpdateWorkspace:output_type -> workspace.UpdateWorkspaceResponse
	10, // 14: workspace.WorkspaceService.DeleteWorkspace:output_type -> workspace.DeleteWorkspaceResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_workspace_proto_init() }
func file_workspace_proto_init() {
	if File_workspace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workspace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkspaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkspaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkspaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkspaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkspaceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkspaceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkspaceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkspaceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_workspace_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workspace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workspace_proto_goTypes,
		DependencyIndexes: file_workspace_proto_depIdxs,
		MessageInfos:      file_workspace_proto_msgTypes,
	}.Build()
	File_workspace_proto = out.File
	file_workspace_proto_rawDesc = nil
	file_workspace_proto_goTypes = nil
	file_workspace_proto_depIdxs = nil
}
