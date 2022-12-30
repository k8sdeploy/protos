// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
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

// eg user:email:edit
// eg user:view
type PermissionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId     string  `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Resource    string  `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"` // user
	Action      string  `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`     // edit
	Allowed     *bool   `protobuf:"varint,4,opt,name=allowed,proto3,oneof" json:"allowed,omitempty"`
	SubResource *string `protobuf:"bytes,5,opt,name=sub_resource,json=subResource,proto3,oneof" json:"sub_resource,omitempty"` // email
	// Types that are assignable to Subject:
	//
	//	*PermissionItem_UserId
	//	*PermissionItem_CompanyId
	//	*PermissionItem_AssetId
	Subject isPermissionItem_Subject `protobuf_oneof:"subject"`
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

func (x *PermissionItem) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *PermissionItem) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *PermissionItem) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *PermissionItem) GetAllowed() bool {
	if x != nil && x.Allowed != nil {
		return *x.Allowed
	}
	return false
}

func (x *PermissionItem) GetSubResource() string {
	if x != nil && x.SubResource != nil {
		return *x.SubResource
	}
	return ""
}

func (m *PermissionItem) GetSubject() isPermissionItem_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (x *PermissionItem) GetUserId() string {
	if x, ok := x.GetSubject().(*PermissionItem_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *PermissionItem) GetCompanyId() string {
	if x, ok := x.GetSubject().(*PermissionItem_CompanyId); ok {
		return x.CompanyId
	}
	return ""
}

func (x *PermissionItem) GetAssetId() string {
	if x, ok := x.GetSubject().(*PermissionItem_AssetId); ok {
		return x.AssetId
	}
	return ""
}

type isPermissionItem_Subject interface {
	isPermissionItem_Subject()
}

type PermissionItem_UserId struct {
	UserId string `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3,oneof"`
}

type PermissionItem_CompanyId struct {
	CompanyId string `protobuf:"bytes,7,opt,name=company_id,json=companyId,proto3,oneof"`
}

type PermissionItem_AssetId struct {
	AssetId string `protobuf:"bytes,8,opt,name=asset_id,json=assetId,proto3,oneof"`
}

func (*PermissionItem_UserId) isPermissionItem_Subject() {}

func (*PermissionItem_CompanyId) isPermissionItem_Subject() {}

func (*PermissionItem_AssetId) isPermissionItem_Subject() {}

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
		mi := &file_v1_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionResponse) ProtoMessage() {}

func (x *PermissionResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PermissionResponse.ProtoReflect.Descriptor instead.
func (*PermissionResponse) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{1}
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

type MultiplePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []*PermissionItem `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *MultiplePermissionRequest) Reset() {
	*x = MultiplePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplePermissionRequest) ProtoMessage() {}

func (x *MultiplePermissionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MultiplePermissionRequest.ProtoReflect.Descriptor instead.
func (*MultiplePermissionRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *MultiplePermissionRequest) GetPermissions() []*PermissionItem {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type AllPermissionsGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *AllPermissionsGetRequest) Reset() {
	*x = AllPermissionsGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_permissions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllPermissionsGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllPermissionsGetRequest) ProtoMessage() {}

func (x *AllPermissionsGetRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AllPermissionsGetRequest.ProtoReflect.Descriptor instead.
func (*AllPermissionsGetRequest) Descriptor() ([]byte, []int) {
	return file_v1_permissions_proto_rawDescGZIP(), []int{3}
}

func (x *AllPermissionsGetRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

var File_v1_permissions_proto protoreflect.FileDescriptor

var file_v1_permissions_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xa7, 0x02, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x07, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x0b, 0x73, 0x75, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x7e, 0x0a, 0x12, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x5d, 0x0a, 0x19, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x35, 0x0a, 0x18, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x32, 0xbe, 0x04, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x22, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x22, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x64, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x05, 0x43, 0x61, 0x6e, 0x44,
	0x6f, 0x12, 0x1e, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65,
	0x6d, 0x1a, 0x22, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x38, 0x73, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_v1_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_permissions_proto_goTypes = []interface{}{
	(*PermissionItem)(nil),            // 0: permissions.v1.PermissionItem
	(*PermissionResponse)(nil),        // 1: permissions.v1.PermissionResponse
	(*MultiplePermissionRequest)(nil), // 2: permissions.v1.MultiplePermissionRequest
	(*AllPermissionsGetRequest)(nil),  // 3: permissions.v1.AllPermissionsGetRequest
}
var file_v1_permissions_proto_depIdxs = []int32{
	0, // 0: permissions.v1.PermissionResponse.permissions:type_name -> permissions.v1.PermissionItem
	0, // 1: permissions.v1.MultiplePermissionRequest.permissions:type_name -> permissions.v1.PermissionItem
	0, // 2: permissions.v1.PermissionService.AddPermission:input_type -> permissions.v1.PermissionItem
	2, // 3: permissions.v1.PermissionService.AddPermissions:input_type -> permissions.v1.MultiplePermissionRequest
	0, // 4: permissions.v1.PermissionService.RemovePermission:input_type -> permissions.v1.PermissionItem
	2, // 5: permissions.v1.PermissionService.RemovePermissions:input_type -> permissions.v1.MultiplePermissionRequest
	3, // 6: permissions.v1.PermissionService.GetPermissions:input_type -> permissions.v1.AllPermissionsGetRequest
	0, // 7: permissions.v1.PermissionService.CanDo:input_type -> permissions.v1.PermissionItem
	1, // 8: permissions.v1.PermissionService.AddPermission:output_type -> permissions.v1.PermissionResponse
	1, // 9: permissions.v1.PermissionService.AddPermissions:output_type -> permissions.v1.PermissionResponse
	1, // 10: permissions.v1.PermissionService.RemovePermission:output_type -> permissions.v1.PermissionResponse
	1, // 11: permissions.v1.PermissionService.RemovePermissions:output_type -> permissions.v1.PermissionResponse
	1, // 12: permissions.v1.PermissionService.GetPermissions:output_type -> permissions.v1.PermissionResponse
	1, // 13: permissions.v1.PermissionService.CanDo:output_type -> permissions.v1.PermissionResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
		file_v1_permissions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplePermissionRequest); i {
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
			switch v := v.(*AllPermissionsGetRequest); i {
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
	file_v1_permissions_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PermissionItem_UserId)(nil),
		(*PermissionItem_CompanyId)(nil),
		(*PermissionItem_AssetId)(nil),
	}
	file_v1_permissions_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
