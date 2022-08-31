// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: v1/orchestrator.proto

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

type DeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceKey    string  `protobuf:"bytes,1,opt,name=service_key,json=serviceKey,proto3" json:"service_key,omitempty"`
	TriggerKey    string  `protobuf:"bytes,2,opt,name=trigger_key,json=triggerKey,proto3" json:"trigger_key,omitempty"`
	CommitMessage *string `protobuf:"bytes,3,opt,name=commit_message,json=commitMessage,proto3,oneof" json:"commit_message,omitempty"`
	CommitId      *string `protobuf:"bytes,4,opt,name=commit_id,json=commitId,proto3,oneof" json:"commit_id,omitempty"`
	CommitUrl     *string `protobuf:"bytes,5,opt,name=commit_url,json=commitUrl,proto3,oneof" json:"commit_url,omitempty"`
}

func (x *DeploymentRequest) Reset() {
	*x = DeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orchestrator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentRequest) ProtoMessage() {}

func (x *DeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orchestrator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentRequest.ProtoReflect.Descriptor instead.
func (*DeploymentRequest) Descriptor() ([]byte, []int) {
	return file_v1_orchestrator_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentRequest) GetServiceKey() string {
	if x != nil {
		return x.ServiceKey
	}
	return ""
}

func (x *DeploymentRequest) GetTriggerKey() string {
	if x != nil {
		return x.TriggerKey
	}
	return ""
}

func (x *DeploymentRequest) GetCommitMessage() string {
	if x != nil && x.CommitMessage != nil {
		return *x.CommitMessage
	}
	return ""
}

func (x *DeploymentRequest) GetCommitId() string {
	if x != nil && x.CommitId != nil {
		return *x.CommitId
	}
	return ""
}

func (x *DeploymentRequest) GetCommitUrl() string {
	if x != nil && x.CommitUrl != nil {
		return *x.CommitUrl
	}
	return ""
}

type DeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployed bool    `protobuf:"varint,1,opt,name=deployed,proto3" json:"deployed,omitempty"`
	Status   *string `protobuf:"bytes,99,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *DeploymentResponse) Reset() {
	*x = DeploymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orchestrator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentResponse) ProtoMessage() {}

func (x *DeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orchestrator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentResponse.ProtoReflect.Descriptor instead.
func (*DeploymentResponse) Descriptor() ([]byte, []int) {
	return file_v1_orchestrator_proto_rawDescGZIP(), []int{1}
}

func (x *DeploymentResponse) GetDeployed() bool {
	if x != nil {
		return x.Deployed
	}
	return false
}

func (x *DeploymentResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

type DeploymentRequestK8SDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageHash        string `protobuf:"bytes,1,opt,name=image_hash,json=imageHash,proto3" json:"image_hash,omitempty"`
	ImageTag         string `protobuf:"bytes,2,opt,name=image_tag,json=imageTag,proto3" json:"image_tag,omitempty"`
	ServiceName      string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ServiceNamespace string `protobuf:"bytes,4,opt,name=service_namespace,json=serviceNamespace,proto3" json:"service_namespace,omitempty"`
}

func (x *DeploymentRequestK8SDetails) Reset() {
	*x = DeploymentRequestK8SDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orchestrator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentRequestK8SDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentRequestK8SDetails) ProtoMessage() {}

func (x *DeploymentRequestK8SDetails) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orchestrator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentRequestK8SDetails.ProtoReflect.Descriptor instead.
func (*DeploymentRequestK8SDetails) Descriptor() ([]byte, []int) {
	return file_v1_orchestrator_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DeploymentRequestK8SDetails) GetImageHash() string {
	if x != nil {
		return x.ImageHash
	}
	return ""
}

func (x *DeploymentRequestK8SDetails) GetImageTag() string {
	if x != nil {
		return x.ImageTag
	}
	return ""
}

func (x *DeploymentRequestK8SDetails) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *DeploymentRequestK8SDetails) GetServiceNamespace() string {
	if x != nil {
		return x.ServiceNamespace
	}
	return ""
}

type DeploymentRequest_AuthorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeploymentRequest_AuthorDetails) Reset() {
	*x = DeploymentRequest_AuthorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orchestrator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentRequest_AuthorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentRequest_AuthorDetails) ProtoMessage() {}

func (x *DeploymentRequest_AuthorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orchestrator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentRequest_AuthorDetails.ProtoReflect.Descriptor instead.
func (*DeploymentRequest_AuthorDetails) Descriptor() ([]byte, []int) {
	return file_v1_orchestrator_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DeploymentRequest_AuthorDetails) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DeploymentRequest_AuthorDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_v1_orchestrator_proto protoreflect.FileDescriptor

var file_v1_orchestrator_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xd4, 0x03, 0x0a, 0x11, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x55, 0x72, 0x6c, 0x88,
	0x01, 0x01, 0x1a, 0x99, 0x01, 0x0a, 0x0b, 0x6b, 0x38, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x61, 0x67, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x3f,
	0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x22,
	0x58, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x63, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x63, 0x0a, 0x0c, 0x4f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x53, 0x0a, 0x06, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x12, 0x22, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x38, 0x73,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_orchestrator_proto_rawDescOnce sync.Once
	file_v1_orchestrator_proto_rawDescData = file_v1_orchestrator_proto_rawDesc
)

func file_v1_orchestrator_proto_rawDescGZIP() []byte {
	file_v1_orchestrator_proto_rawDescOnce.Do(func() {
		file_v1_orchestrator_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_orchestrator_proto_rawDescData)
	})
	return file_v1_orchestrator_proto_rawDescData
}

var file_v1_orchestrator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_orchestrator_proto_goTypes = []interface{}{
	(*DeploymentRequest)(nil),               // 0: orchestrator.v1.DeploymentRequest
	(*DeploymentResponse)(nil),              // 1: orchestrator.v1.DeploymentResponse
	(*DeploymentRequestK8SDetails)(nil),     // 2: orchestrator.v1.DeploymentRequest.k8s_details
	(*DeploymentRequest_AuthorDetails)(nil), // 3: orchestrator.v1.DeploymentRequest.AuthorDetails
}
var file_v1_orchestrator_proto_depIdxs = []int32{
	0, // 0: orchestrator.v1.Orchestrator.Deploy:input_type -> orchestrator.v1.DeploymentRequest
	1, // 1: orchestrator.v1.Orchestrator.Deploy:output_type -> orchestrator.v1.DeploymentResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_orchestrator_proto_init() }
func file_v1_orchestrator_proto_init() {
	if File_v1_orchestrator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_orchestrator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentRequest); i {
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
		file_v1_orchestrator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentResponse); i {
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
		file_v1_orchestrator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentRequestK8SDetails); i {
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
		file_v1_orchestrator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentRequest_AuthorDetails); i {
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
	file_v1_orchestrator_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_orchestrator_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_orchestrator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_orchestrator_proto_goTypes,
		DependencyIndexes: file_v1_orchestrator_proto_depIdxs,
		MessageInfos:      file_v1_orchestrator_proto_msgTypes,
	}.Build()
	File_v1_orchestrator_proto = out.File
	file_v1_orchestrator_proto_rawDesc = nil
	file_v1_orchestrator_proto_goTypes = nil
	file_v1_orchestrator_proto_depIdxs = nil
}
