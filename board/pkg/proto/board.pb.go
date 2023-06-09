// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: board.proto

package board

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

type CreateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateBoardRequest) Reset() {
	*x = CreateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardRequest) ProtoMessage() {}

func (x *CreateBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardRequest.ProtoReflect.Descriptor instead.
func (*CreateBoardRequest) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBoardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardID string `protobuf:"bytes,1,opt,name=boardID,proto3" json:"boardID,omitempty"`
}

func (x *GetBoardRequest) Reset() {
	*x = GetBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoardRequest) ProtoMessage() {}

func (x *GetBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoardRequest.ProtoReflect.Descriptor instead.
func (*GetBoardRequest) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{1}
}

func (x *GetBoardRequest) GetBoardID() string {
	if x != nil {
		return x.BoardID
	}
	return ""
}

type ListBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBoardRequest) Reset() {
	*x = ListBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBoardRequest) ProtoMessage() {}

func (x *ListBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBoardRequest.ProtoReflect.Descriptor instead.
func (*ListBoardRequest) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{2}
}

type UpdateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardID string `protobuf:"bytes,1,opt,name=boardID,proto3" json:"boardID,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateBoardRequest) Reset() {
	*x = UpdateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoardRequest) ProtoMessage() {}

func (x *UpdateBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoardRequest.ProtoReflect.Descriptor instead.
func (*UpdateBoardRequest) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBoardRequest) GetBoardID() string {
	if x != nil {
		return x.BoardID
	}
	return ""
}

func (x *UpdateBoardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardID string `protobuf:"bytes,1,opt,name=boardID,proto3" json:"boardID,omitempty"`
}

func (x *DeleteBoardRequest) Reset() {
	*x = DeleteBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBoardRequest) ProtoMessage() {}

func (x *DeleteBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBoardRequest.ProtoReflect.Descriptor instead.
func (*DeleteBoardRequest) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBoardRequest) GetBoardID() string {
	if x != nil {
		return x.BoardID
	}
	return ""
}

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{5}
}

func (x *Board) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Board) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Board) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Board) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Board *Board `protobuf:"bytes,1,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *CreateBoardResponse) Reset() {
	*x = CreateBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardResponse) ProtoMessage() {}

func (x *CreateBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardResponse.ProtoReflect.Descriptor instead.
func (*CreateBoardResponse) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{6}
}

func (x *CreateBoardResponse) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

type GetBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Board *Board `protobuf:"bytes,1,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *GetBoardResponse) Reset() {
	*x = GetBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoardResponse) ProtoMessage() {}

func (x *GetBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoardResponse.ProtoReflect.Descriptor instead.
func (*GetBoardResponse) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{7}
}

func (x *GetBoardResponse) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

type ListBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boards []*Board `protobuf:"bytes,1,rep,name=boards,proto3" json:"boards,omitempty"`
}

func (x *ListBoardResponse) Reset() {
	*x = ListBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBoardResponse) ProtoMessage() {}

func (x *ListBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBoardResponse.ProtoReflect.Descriptor instead.
func (*ListBoardResponse) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{8}
}

func (x *ListBoardResponse) GetBoards() []*Board {
	if x != nil {
		return x.Boards
	}
	return nil
}

type UpdateBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Board *Board `protobuf:"bytes,1,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *UpdateBoardResponse) Reset() {
	*x = UpdateBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoardResponse) ProtoMessage() {}

func (x *UpdateBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoardResponse.ProtoReflect.Descriptor instead.
func (*UpdateBoardResponse) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateBoardResponse) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

type DeleteBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Board *Board `protobuf:"bytes,1,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *DeleteBoardResponse) Reset() {
	*x = DeleteBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBoardResponse) ProtoMessage() {}

func (x *DeleteBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_board_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBoardResponse.ProtoReflect.Descriptor instead.
func (*DeleteBoardResponse) Descriptor() ([]byte, []int) {
	return file_board_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteBoardResponse) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

var File_board_proto protoreflect.FileDescriptor

var file_board_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44, 0x22, 0x12, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x42, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x49, 0x44, 0x22, 0x67, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x39, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22,
	0x39, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x22, 0x39, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x32, 0xe8, 0x02, 0x0a, 0x0c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x12, 0x19, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x67,
	0x6f, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_board_proto_rawDescOnce sync.Once
	file_board_proto_rawDescData = file_board_proto_rawDesc
)

func file_board_proto_rawDescGZIP() []byte {
	file_board_proto_rawDescOnce.Do(func() {
		file_board_proto_rawDescData = protoimpl.X.CompressGZIP(file_board_proto_rawDescData)
	})
	return file_board_proto_rawDescData
}

var file_board_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_board_proto_goTypes = []interface{}{
	(*CreateBoardRequest)(nil),  // 0: board.CreateBoardRequest
	(*GetBoardRequest)(nil),     // 1: board.GetBoardRequest
	(*ListBoardRequest)(nil),    // 2: board.ListBoardRequest
	(*UpdateBoardRequest)(nil),  // 3: board.UpdateBoardRequest
	(*DeleteBoardRequest)(nil),  // 4: board.DeleteBoardRequest
	(*Board)(nil),               // 5: board.Board
	(*CreateBoardResponse)(nil), // 6: board.CreateBoardResponse
	(*GetBoardResponse)(nil),    // 7: board.GetBoardResponse
	(*ListBoardResponse)(nil),   // 8: board.ListBoardResponse
	(*UpdateBoardResponse)(nil), // 9: board.UpdateBoardResponse
	(*DeleteBoardResponse)(nil), // 10: board.DeleteBoardResponse
}
var file_board_proto_depIdxs = []int32{
	5,  // 0: board.CreateBoardResponse.board:type_name -> board.Board
	5,  // 1: board.GetBoardResponse.board:type_name -> board.Board
	5,  // 2: board.ListBoardResponse.boards:type_name -> board.Board
	5,  // 3: board.UpdateBoardResponse.board:type_name -> board.Board
	5,  // 4: board.DeleteBoardResponse.board:type_name -> board.Board
	0,  // 5: board.BoardService.CreateBoard:input_type -> board.CreateBoardRequest
	1,  // 6: board.BoardService.GetBoard:input_type -> board.GetBoardRequest
	2,  // 7: board.BoardService.ListBoards:input_type -> board.ListBoardRequest
	3,  // 8: board.BoardService.UpdateBoard:input_type -> board.UpdateBoardRequest
	4,  // 9: board.BoardService.DeleteBoard:input_type -> board.DeleteBoardRequest
	6,  // 10: board.BoardService.CreateBoard:output_type -> board.CreateBoardResponse
	7,  // 11: board.BoardService.GetBoard:output_type -> board.GetBoardResponse
	8,  // 12: board.BoardService.ListBoards:output_type -> board.ListBoardResponse
	9,  // 13: board.BoardService.UpdateBoard:output_type -> board.UpdateBoardResponse
	10, // 14: board.BoardService.DeleteBoard:output_type -> board.DeleteBoardResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_board_proto_init() }
func file_board_proto_init() {
	if File_board_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_board_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardRequest); i {
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
		file_board_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBoardRequest); i {
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
		file_board_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBoardRequest); i {
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
		file_board_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoardRequest); i {
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
		file_board_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBoardRequest); i {
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
		file_board_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
		file_board_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardResponse); i {
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
		file_board_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBoardResponse); i {
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
		file_board_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBoardResponse); i {
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
		file_board_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoardResponse); i {
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
		file_board_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBoardResponse); i {
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
			RawDescriptor: file_board_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_board_proto_goTypes,
		DependencyIndexes: file_board_proto_depIdxs,
		MessageInfos:      file_board_proto_msgTypes,
	}.Build()
	File_board_proto = out.File
	file_board_proto_rawDesc = nil
	file_board_proto_goTypes = nil
	file_board_proto_depIdxs = nil
}
