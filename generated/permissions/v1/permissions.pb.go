// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
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

type PermissionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Allowed string `protobuf:"bytes,2,opt,name=allowed,proto3" json:"allowed,omitempty"`
}

func (x *PermissionItem) Reset() {
	*x = PermissionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionItem) ProtoMessage() {}

func (x *PermissionItem) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PermissionItem.ProtoReflect.Descriptor instead.
func (*PermissionItem) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionItem) GetAllowed() string {
	if x != nil {
		return x.Allowed
	}
	return ""
}

type AllowedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid     string  `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Companyid  *string `protobuf:"bytes,2,opt,name=companyid,proto3,oneof" json:"companyid,omitempty"`
	Resourceid *string `protobuf:"bytes,3,opt,name=resourceid,proto3,oneof" json:"resourceid,omitempty"`
}

func (x *AllowedRequest) Reset() {
	*x = AllowedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedRequest) ProtoMessage() {}

func (x *AllowedRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AllowedRequest.ProtoReflect.Descriptor instead.
func (*AllowedRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *AllowedRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *AllowedRequest) GetCompanyid() string {
	if x != nil && x.Companyid != nil {
		return *x.Companyid
	}
	return ""
}

func (x *AllowedRequest) GetResourceid() string {
	if x != nil && x.Resourceid != nil {
		return *x.Resourceid
	}
	return ""
}

type CanDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid     string  `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Action     string  `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Companyid  *string `protobuf:"bytes,3,opt,name=companyid,proto3,oneof" json:"companyid,omitempty"`
	Resourceid *string `protobuf:"bytes,4,opt,name=resourceid,proto3,oneof" json:"resourceid,omitempty"`
}

func (x *CanDoRequest) Reset() {
	*x = CanDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanDoRequest) ProtoMessage() {}

func (x *CanDoRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CanDoRequest.ProtoReflect.Descriptor instead.
func (*CanDoRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *CanDoRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *CanDoRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *CanDoRequest) GetCompanyid() string {
	if x != nil && x.Companyid != nil {
		return *x.Companyid
	}
	return ""
}

func (x *CanDoRequest) GetResourceid() string {
	if x != nil && x.Resourceid != nil {
		return *x.Resourceid
	}
	return ""
}

type AddPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid     string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Companyid  string `protobuf:"bytes,2,opt,name=companyid,proto3" json:"companyid,omitempty"`
	Resourceid string `protobuf:"bytes,3,opt,name=resourceid,proto3" json:"resourceid,omitempty"`
	Permission string `protobuf:"bytes,4,opt,name=permission,proto3" json:"permission,omitempty"`
	Allowed    bool   `protobuf:"varint,5,opt,name=allowed,proto3" json:"allowed,omitempty"`
}

func (x *AddPermissionRequest) Reset() {
	*x = AddPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermissionRequest) ProtoMessage() {}

func (x *AddPermissionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddPermissionRequest.ProtoReflect.Descriptor instead.
func (*AddPermissionRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{3}
}

func (x *AddPermissionRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *AddPermissionRequest) GetCompanyid() string {
	if x != nil {
		return x.Companyid
	}
	return ""
}

func (x *AddPermissionRequest) GetResourceid() string {
	if x != nil {
		return x.Resourceid
	}
	return ""
}

func (x *AddPermissionRequest) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *AddPermissionRequest) GetAllowed() bool {
	if x != nil {
		return x.Allowed
	}
	return false
}

type RemovePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid     string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Companyid  string `protobuf:"bytes,2,opt,name=companyid,proto3" json:"companyid,omitempty"`
	Resourceid string `protobuf:"bytes,3,opt,name=resourceid,proto3" json:"resourceid,omitempty"`
	Permission string `protobuf:"bytes,4,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *RemovePermissionRequest) Reset() {
	*x = RemovePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePermissionRequest) ProtoMessage() {}

func (x *RemovePermissionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RemovePermissionRequest.ProtoReflect.Descriptor instead.
func (*RemovePermissionRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{4}
}

func (x *RemovePermissionRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *RemovePermissionRequest) GetCompanyid() string {
	if x != nil {
		return x.Companyid
	}
	return ""
}

func (x *RemovePermissionRequest) GetResourceid() string {
	if x != nil {
		return x.Resourceid
	}
	return ""
}

func (x *RemovePermissionRequest) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

type PermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []*PermissionItem `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Status      *string           `protobuf:"bytes,99,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *PermissionResponse) Reset() {
	*x = PermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionResponse) ProtoMessage() {}

func (x *PermissionResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PermissionResponse.ProtoReflect.Descriptor instead.
func (*PermissionResponse) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{5}
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

var File_v1_permissions_proto protoreflect.FileDescriptor

var file_v1_permissions_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x3e, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x69, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x44, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x69, 0x64, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x69, 0x64, 0x22, 0xa6, 0x01, 0x0a,
	0x14, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x7e, 0x0a, 0x12, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xf8, 0x02, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1e, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x05, 0x43, 0x61, 0x6e, 0x44, 0x6f, 0x12, 0x1c,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x6e, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x38, 0x73, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_v1_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_permissions_proto_goTypes = []interface{}{
	(*PermissionItem)(nil),          // 0: permissions.v1.PermissionItem
	(*AllowedRequest)(nil),          // 1: permissions.v1.AllowedRequest
	(*CanDoRequest)(nil),            // 2: permissions.v1.CanDoRequest
	(*AddPermissionRequest)(nil),    // 3: permissions.v1.AddPermissionRequest
	(*RemovePermissionRequest)(nil), // 4: permissions.v1.RemovePermissionRequest
	(*PermissionResponse)(nil),      // 5: permissions.v1.PermissionResponse
}
var file_v1_permissions_proto_depIdxs = []int32{
	0, // 0: permissions.v1.PermissionResponse.permissions:type_name -> permissions.v1.PermissionItem
	1, // 1: permissions.v1.PermissionService.GetPermissions:input_type -> permissions.v1.AllowedRequest
	3, // 2: permissions.v1.PermissionService.AddPermission:input_type -> permissions.v1.AddPermissionRequest
	4, // 3: permissions.v1.PermissionService.RemovePermission:input_type -> permissions.v1.RemovePermissionRequest
	2, // 4: permissions.v1.PermissionService.CanDo:input_type -> permissions.v1.CanDoRequest
	5, // 5: permissions.v1.PermissionService.GetPermissions:output_type -> permissions.v1.PermissionResponse
	5, // 6: permissions.v1.PermissionService.AddPermission:output_type -> permissions.v1.PermissionResponse
	5, // 7: permissions.v1.PermissionService.RemovePermission:output_type -> permissions.v1.PermissionResponse
	5, // 8: permissions.v1.PermissionService.CanDo:output_type -> permissions.v1.PermissionResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_permissions_proto_init() }
func file_v1_permissions_proto_init() {
	if File_v1_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_permissions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_permissions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanDoRequest); i {
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
			switch v := v.(*AddPermissionRequest); i {
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
			switch v := v.(*RemovePermissionRequest); i {
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
	}
	file_v1_permissions_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_v1_permissions_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_v1_permissions_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_permissions_proto_goTypes,
		DependencyIndexes: file_v1_permissions_proto_depIdxs,
		MessageInfos:      file_v1_permissions_proto_msgTypes,
	}.Build()
	File_v1_permissions_proto = out.File
	file_v1_permissions_proto_rawDesc = nil
	file_v1_permissions_proto_goTypes = nil
	file_v1_permissions_proto_depIdxs = nil
}
