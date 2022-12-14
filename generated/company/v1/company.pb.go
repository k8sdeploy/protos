// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: v1/company.proto

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

type CompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserKey      string  `protobuf:"bytes,1,opt,name=user_key,json=userKey,proto3" json:"user_key,omitempty"`
	CompanyKey   *string `protobuf:"bytes,2,opt,name=company_key,json=companyKey,proto3,oneof" json:"company_key,omitempty"`
	CompanyName  *string `protobuf:"bytes,3,opt,name=company_name,json=companyName,proto3,oneof" json:"company_name,omitempty"`
	CompanyEmail *string `protobuf:"bytes,4,opt,name=company_email,json=companyEmail,proto3,oneof" json:"company_email,omitempty"`
}

func (x *CompanyRequest) Reset() {
	*x = CompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyRequest) ProtoMessage() {}

func (x *CompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyRequest.ProtoReflect.Descriptor instead.
func (*CompanyRequest) Descriptor() ([]byte, []int) {
	return file_v1_company_proto_rawDescGZIP(), []int{0}
}

func (x *CompanyRequest) GetUserKey() string {
	if x != nil {
		return x.UserKey
	}
	return ""
}

func (x *CompanyRequest) GetCompanyKey() string {
	if x != nil && x.CompanyKey != nil {
		return *x.CompanyKey
	}
	return ""
}

func (x *CompanyRequest) GetCompanyName() string {
	if x != nil && x.CompanyName != nil {
		return *x.CompanyName
	}
	return ""
}

func (x *CompanyRequest) GetCompanyEmail() string {
	if x != nil && x.CompanyEmail != nil {
		return *x.CompanyEmail
	}
	return ""
}

type CompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyKey   string  `protobuf:"bytes,1,opt,name=company_key,json=companyKey,proto3" json:"company_key,omitempty"`
	CompanyName  string  `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	CompanyEmail string  `protobuf:"bytes,3,opt,name=company_email,json=companyEmail,proto3" json:"company_email,omitempty"`
	Status       *string `protobuf:"bytes,99,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *CompanyResponse) Reset() {
	*x = CompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyResponse) ProtoMessage() {}

func (x *CompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyResponse.ProtoReflect.Descriptor instead.
func (*CompanyResponse) Descriptor() ([]byte, []int) {
	return file_v1_company_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyResponse) GetCompanyKey() string {
	if x != nil {
		return x.CompanyKey
	}
	return ""
}

func (x *CompanyResponse) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *CompanyResponse) GetCompanyEmail() string {
	if x != nil {
		return x.CompanyEmail
	}
	return ""
}

func (x *CompanyResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

var File_v1_company_proto protoreflect.FileDescriptor

var file_v1_company_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x22, 0xd6, 0x01, 0x0a, 0x0e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x26,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6b, 0x65, 0x79,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0xa2, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x95, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x16, 0x2e, 0x6b,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x16, 0x2e, 0x6b, 0x65,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x16, 0x2e,
	0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x16, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x38, 0x73, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_company_proto_rawDescOnce sync.Once
	file_v1_company_proto_rawDescData = file_v1_company_proto_rawDesc
)

func file_v1_company_proto_rawDescGZIP() []byte {
	file_v1_company_proto_rawDescOnce.Do(func() {
		file_v1_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_company_proto_rawDescData)
	})
	return file_v1_company_proto_rawDescData
}

var file_v1_company_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_company_proto_goTypes = []interface{}{
	(*CompanyRequest)(nil),  // 0: key.v1.CompanyRequest
	(*CompanyResponse)(nil), // 1: key.v1.CompanyResponse
}
var file_v1_company_proto_depIdxs = []int32{
	0, // 0: key.v1.CompanyService.CreateCompany:input_type -> key.v1.CompanyRequest
	0, // 1: key.v1.CompanyService.GetCompany:input_type -> key.v1.CompanyRequest
	0, // 2: key.v1.CompanyService.UpdateCompany:input_type -> key.v1.CompanyRequest
	0, // 3: key.v1.CompanyService.DeleteCompany:input_type -> key.v1.CompanyRequest
	1, // 4: key.v1.CompanyService.CreateCompany:output_type -> key.v1.CompanyResponse
	1, // 5: key.v1.CompanyService.GetCompany:output_type -> key.v1.CompanyResponse
	1, // 6: key.v1.CompanyService.UpdateCompany:output_type -> key.v1.CompanyResponse
	1, // 7: key.v1.CompanyService.DeleteCompany:output_type -> key.v1.CompanyResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_company_proto_init() }
func file_v1_company_proto_init() {
	if File_v1_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyRequest); i {
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
		file_v1_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyResponse); i {
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
	file_v1_company_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_company_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_company_proto_goTypes,
		DependencyIndexes: file_v1_company_proto_depIdxs,
		MessageInfos:      file_v1_company_proto_msgTypes,
	}.Build()
	File_v1_company_proto = out.File
	file_v1_company_proto_rawDesc = nil
	file_v1_company_proto_goTypes = nil
	file_v1_company_proto_depIdxs = nil
}
