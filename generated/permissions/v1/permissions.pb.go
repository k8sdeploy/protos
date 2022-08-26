// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: v1/permissions.proto

package v1

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

type CreateUserRequest_Level int32

const (
	CreateUserRequest_USER   CreateUserRequest_Level = 0
	CreateUserRequest_LEADER CreateUserRequest_Level = 1
	CreateUserRequest_OWNER  CreateUserRequest_Level = 3
)

// Enum value maps for CreateUserRequest_Level.
var (
	CreateUserRequest_Level_name = map[int32]string{
		0: "USER",
		1: "LEADER",
		3: "OWNER",
	}
	CreateUserRequest_Level_value = map[string]int32{
		"USER":   0,
		"LEADER": 1,
		"OWNER":  3,
	}
)

func (x CreateUserRequest_Level) Enum() *CreateUserRequest_Level {
	p := new(CreateUserRequest_Level)
	*p = x
	return p
}

func (x CreateUserRequest_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateUserRequest_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_permissions_proto_enumTypes[0].Descriptor()
}

func (CreateUserRequest_Level) Type() protoreflect.EnumType {
	return &file_v1_permissions_proto_enumTypes[0]
}

func (x CreateUserRequest_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateUserRequest_Level.Descriptor instead.
func (CreateUserRequest_Level) EnumDescriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{1, 0}
}

type PermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	APIKey   string `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	Action   string `protobuf:"bytes,3,opt,name=Action,proto3" json:"Action,omitempty"`
	Resource string `protobuf:"bytes,4,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Filter   string `protobuf:"bytes,5,opt,name=Filter,proto3" json:"Filter,omitempty"`
}

func (x *PermissionRequest) Reset() {
	*x = PermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionRequest) ProtoMessage() {}

func (x *PermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionRequest.ProtoReflect.Descriptor instead.
func (*PermissionRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *PermissionRequest) GetAPIKey() string {
	if x != nil {
		return x.APIKey
	}
	return ""
}

func (x *PermissionRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *PermissionRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *PermissionRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string                   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	APIKey    string                   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	Level     *CreateUserRequest_Level `protobuf:"varint,3,opt,name=level,proto3,enum=permissions.v1.CreateUserRequest_Level,oneof" json:"level,omitempty"`
	CompanyID *string                  `protobuf:"bytes,4,opt,name=CompanyID,proto3,oneof" json:"CompanyID,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_permissions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateUserRequest) GetAPIKey() string {
	if x != nil {
		return x.APIKey
	}
	return ""
}

func (x *CreateUserRequest) GetLevel() CreateUserRequest_Level {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return CreateUserRequest_USER
}

func (x *CreateUserRequest) GetCompanyID() string {
	if x != nil && x.CompanyID != nil {
		return *x.CompanyID
	}
	return ""
}

type MultiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*PermissionRequest `protobuf:"bytes,1,rep,name=Requests,proto3" json:"Requests,omitempty"`
}

func (x *MultiRequest) Reset() {
	*x = MultiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiRequest) ProtoMessage() {}

func (x *MultiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_permissions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiRequest.ProtoReflect.Descriptor instead.
func (*MultiRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *MultiRequest) GetRequests() []*PermissionRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

type PermissionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action   string `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"`
	Resource string `protobuf:"bytes,2,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Filter   string `protobuf:"bytes,3,opt,name=Filter,proto3" json:"Filter,omitempty"`
}

func (x *PermissionItem) Reset() {
	*x = PermissionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionItem) ProtoMessage() {}

func (x *PermissionItem) ProtoReflect() protoreflect.Message {
	mi := &file_v1_permissions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionItem.ProtoReflect.Descriptor instead.
func (*PermissionItem) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionItem) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *PermissionItem) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *PermissionItem) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	APIKey string `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_permissions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetRequest) GetAPIKey() string {
	if x != nil {
		return x.APIKey
	}
	return ""
}

type AllowedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	APIKey   string `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	Action   string `protobuf:"bytes,3,opt,name=Action,proto3" json:"Action,omitempty"`
	Resource string `protobuf:"bytes,4,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Filter   string `protobuf:"bytes,5,opt,name=Filter,proto3" json:"Filter,omitempty"`
}

func (x *AllowedRequest) Reset() {
	*x = AllowedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedRequest) ProtoMessage() {}

func (x *AllowedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_permissions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowedRequest.ProtoReflect.Descriptor instead.
func (*AllowedRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{5}
}

func (x *AllowedRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AllowedRequest) GetAPIKey() string {
	if x != nil {
		return x.APIKey
	}
	return ""
}

func (x *AllowedRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *AllowedRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *AllowedRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type PermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []*PermissionItem `protobuf:"bytes,1,rep,name=Permissions,proto3" json:"Permissions,omitempty"`
	Status      *string           `protobuf:"bytes,99,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *PermissionResponse) Reset() {
	*x = PermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionResponse) ProtoMessage() {}

func (x *PermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_permissions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionResponse.ProtoReflect.Descriptor instead.
func (*PermissionResponse) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{6}
}

func (x *PermissionResponse) GetPermissions() []*PermissionItem {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *PermissionResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []*PermissionItem `protobuf:"bytes,1,rep,name=Permissions,proto3" json:"Permissions,omitempty"`
	Status      *string           `protobuf:"bytes,99,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_permissions_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{7}
}

func (x *GetResponse) GetPermissions() []*PermissionItem {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *GetResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

type AllowedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allowed bool    `protobuf:"varint,1,opt,name=Allowed,proto3" json:"Allowed,omitempty"`
	Status  *string `protobuf:"bytes,99,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *AllowedResponse) Reset() {
	*x = AllowedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedResponse) ProtoMessage() {}

func (x *AllowedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_permissions_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowedResponse.ProtoReflect.Descriptor instead.
func (*AllowedResponse) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{8}
}

func (x *AllowedResponse) GetAllowed() bool {
	if x != nil {
		return x.Allowed
	}
	return false
}

func (x *AllowedResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

var File_v1_permissions_proto protoreflect.FileDescriptor

var file_v1_permissions_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xec, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x42,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x22, 0x28, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x08,
	0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x45, 0x41, 0x44,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22, 0x4d, 0x0a, 0x0c, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x50,
	0x49, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x50, 0x49, 0x4b,
	0x65, 0x79, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0x7e, 0x0a, 0x12, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x77, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x40, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x63, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x53, 0x0a, 0x0f, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0xa7, 0x03, 0x0a, 0x12, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59,
	0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x48, 0x0a, 0x05, 0x43, 0x61, 0x6e, 0x44, 0x6f, 0x12, 0x1e, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x38, 0x73, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_permissions_proto_rawDescOnce sync.Once
	file_v1_permissions_proto_rawDescData = file_v1_permissions_proto_rawDesc
)

func file_v1_permissions_proto_rawDescGZIP() []byte {
	file_v1_permissions_proto_rawDescOnce.Do(func() {
		file_v1_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_permissions_proto_rawDescData)
	})
	return file_v1_permissions_proto_rawDescData
}

var file_v1_permissions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_permissions_proto_goTypes = []interface{}{
	(CreateUserRequest_Level)(0), // 0: permissions.v1.CreateUserRequest.Level
	(*PermissionRequest)(nil),    // 1: permissions.v1.PermissionRequest
	(*CreateUserRequest)(nil),    // 2: permissions.v1.CreateUserRequest
	(*MultiRequest)(nil),         // 3: permissions.v1.MultiRequest
	(*PermissionItem)(nil),       // 4: permissions.v1.PermissionItem
	(*GetRequest)(nil),           // 5: permissions.v1.GetRequest
	(*AllowedRequest)(nil),       // 6: permissions.v1.AllowedRequest
	(*PermissionResponse)(nil),   // 7: permissions.v1.PermissionResponse
	(*GetResponse)(nil),          // 8: permissions.v1.GetResponse
	(*AllowedResponse)(nil),      // 9: permissions.v1.AllowedResponse
}
var file_v1_permissions_proto_depIdxs = []int32{
	0, // 0: permissions.v1.CreateUserRequest.level:type_name -> permissions.v1.CreateUserRequest.Level
	1, // 1: permissions.v1.MultiRequest.Requests:type_name -> permissions.v1.PermissionRequest
	4, // 2: permissions.v1.PermissionResponse.Permissions:type_name -> permissions.v1.PermissionItem
	4, // 3: permissions.v1.GetResponse.Permissions:type_name -> permissions.v1.PermissionItem
	1, // 4: permissions.v1.PermissionsService.AddPermission:input_type -> permissions.v1.PermissionRequest
	1, // 5: permissions.v1.PermissionsService.RemovePermission:input_type -> permissions.v1.PermissionRequest
	3, // 6: permissions.v1.PermissionsService.MultiPermissions:input_type -> permissions.v1.MultiRequest
	5, // 7: permissions.v1.PermissionsService.Get:input_type -> permissions.v1.GetRequest
	6, // 8: permissions.v1.PermissionsService.CanDo:input_type -> permissions.v1.AllowedRequest
	7, // 9: permissions.v1.PermissionsService.AddPermission:output_type -> permissions.v1.PermissionResponse
	7, // 10: permissions.v1.PermissionsService.RemovePermission:output_type -> permissions.v1.PermissionResponse
	7, // 11: permissions.v1.PermissionsService.MultiPermissions:output_type -> permissions.v1.PermissionResponse
	8, // 12: permissions.v1.PermissionsService.Get:output_type -> permissions.v1.GetResponse
	9, // 13: permissions.v1.PermissionsService.CanDo:output_type -> permissions.v1.AllowedResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_permissions_proto_init() }
func file_v1_permissions_proto_init() {
	if File_v1_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionRequest); i {
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
		file_v1_permissions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_v1_permissions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiRequest); i {
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
		file_v1_permissions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionItem); i {
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
		file_v1_permissions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_v1_permissions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedRequest); i {
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
		file_v1_permissions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionResponse); i {
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
		file_v1_permissions_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_v1_permissions_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedResponse); i {
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
	file_v1_permissions_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_v1_permissions_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_v1_permissions_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_v1_permissions_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_permissions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_permissions_proto_goTypes,
		DependencyIndexes: file_v1_permissions_proto_depIdxs,
		EnumInfos:         file_v1_permissions_proto_enumTypes,
		MessageInfos:      file_v1_permissions_proto_msgTypes,
	}.Build()
	File_v1_permissions_proto = out.File
	file_v1_permissions_proto_rawDesc = nil
	file_v1_permissions_proto_goTypes = nil
	file_v1_permissions_proto_depIdxs = nil
}