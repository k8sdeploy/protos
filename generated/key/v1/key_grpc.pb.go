// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: v1/key.proto

package v1

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

// KeyServiceClient is the client API for KeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyServiceClient interface {
	CreateAgentKeys(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*KeyResponse, error)
	GetAgentKeys(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*KeyResponse, error)
	ValidateAgentKey(ctx context.Context, in *ValidateSystemKeyRequest, opts ...grpc.CallOption) (*ValidKeyResponse, error)
	CreateHookKeys(ctx context.Context, in *HooksRequest, opts ...grpc.CallOption) (*KeyResponse, error)
	GetHookKeysForCompany(ctx context.Context, in *HooksRequest, opts ...grpc.CallOption) (*MultipleHooksResponse, error)
	ValidateHookKey(ctx context.Context, in *ValidateSystemKeyRequest, opts ...grpc.CallOption) (*ValidKeyResponse, error)
	CreateUserKeys(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*KeyResponse, error)
	ValidateUserKeys(ctx context.Context, in *ValidateUserKeyRequest, opts ...grpc.CallOption) (*ValidKeyResponse, error)
}

type keyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyServiceClient(cc grpc.ClientConnInterface) KeyServiceClient {
	return &keyServiceClient{cc}
}

func (c *keyServiceClient) CreateAgentKeys(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/key.v1.KeyService/CreateAgentKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) GetAgentKeys(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/key.v1.KeyService/GetAgentKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) ValidateAgentKey(ctx context.Context, in *ValidateSystemKeyRequest, opts ...grpc.CallOption) (*ValidKeyResponse, error) {
	out := new(ValidKeyResponse)
	err := c.cc.Invoke(ctx, "/key.v1.KeyService/ValidateAgentKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) CreateHookKeys(ctx context.Context, in *HooksRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/key.v1.KeyService/CreateHookKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) GetHookKeysForCompany(ctx context.Context, in *HooksRequest, opts ...grpc.CallOption) (*MultipleHooksResponse, error) {
	out := new(MultipleHooksResponse)
	err := c.cc.Invoke(ctx, "/key.v1.KeyService/GetHookKeysForCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) ValidateHookKey(ctx context.Context, in *ValidateSystemKeyRequest, opts ...grpc.CallOption) (*ValidKeyResponse, error) {
	out := new(ValidKeyResponse)
	err := c.cc.Invoke(ctx, "/key.v1.KeyService/ValidateHookKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) CreateUserKeys(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/key.v1.KeyService/CreateUserKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyServiceClient) ValidateUserKeys(ctx context.Context, in *ValidateUserKeyRequest, opts ...grpc.CallOption) (*ValidKeyResponse, error) {
	out := new(ValidKeyResponse)
	err := c.cc.Invoke(ctx, "/key.v1.KeyService/ValidateUserKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyServiceServer is the server API for KeyService service.
// All implementations must embed UnimplementedKeyServiceServer
// for forward compatibility
type KeyServiceServer interface {
	CreateAgentKeys(context.Context, *AgentRequest) (*KeyResponse, error)
	GetAgentKeys(context.Context, *AgentRequest) (*KeyResponse, error)
	ValidateAgentKey(context.Context, *ValidateSystemKeyRequest) (*ValidKeyResponse, error)
	CreateHookKeys(context.Context, *HooksRequest) (*KeyResponse, error)
	GetHookKeysForCompany(context.Context, *HooksRequest) (*MultipleHooksResponse, error)
	ValidateHookKey(context.Context, *ValidateSystemKeyRequest) (*ValidKeyResponse, error)
	CreateUserKeys(context.Context, *UserRequest) (*KeyResponse, error)
	ValidateUserKeys(context.Context, *ValidateUserKeyRequest) (*ValidKeyResponse, error)
	mustEmbedUnimplementedKeyServiceServer()
}

// UnimplementedKeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyServiceServer struct {
}

func (UnimplementedKeyServiceServer) CreateAgentKeys(context.Context, *AgentRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgentKeys not implemented")
}
func (UnimplementedKeyServiceServer) GetAgentKeys(context.Context, *AgentRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentKeys not implemented")
}
func (UnimplementedKeyServiceServer) ValidateAgentKey(context.Context, *ValidateSystemKeyRequest) (*ValidKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAgentKey not implemented")
}
func (UnimplementedKeyServiceServer) CreateHookKeys(context.Context, *HooksRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHookKeys not implemented")
}
func (UnimplementedKeyServiceServer) GetHookKeysForCompany(context.Context, *HooksRequest) (*MultipleHooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHookKeysForCompany not implemented")
}
func (UnimplementedKeyServiceServer) ValidateHookKey(context.Context, *ValidateSystemKeyRequest) (*ValidKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateHookKey not implemented")
}
func (UnimplementedKeyServiceServer) CreateUserKeys(context.Context, *UserRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserKeys not implemented")
}
func (UnimplementedKeyServiceServer) ValidateUserKeys(context.Context, *ValidateUserKeyRequest) (*ValidKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserKeys not implemented")
}
func (UnimplementedKeyServiceServer) mustEmbedUnimplementedKeyServiceServer() {}

// UnsafeKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyServiceServer will
// result in compilation errors.
type UnsafeKeyServiceServer interface {
	mustEmbedUnimplementedKeyServiceServer()
}

func RegisterKeyServiceServer(s grpc.ServiceRegistrar, srv KeyServiceServer) {
	s.RegisterService(&KeyService_ServiceDesc, srv)
}

func _KeyService_CreateAgentKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).CreateAgentKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.v1.KeyService/CreateAgentKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).CreateAgentKeys(ctx, req.(*AgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_GetAgentKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GetAgentKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.v1.KeyService/GetAgentKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GetAgentKeys(ctx, req.(*AgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_ValidateAgentKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateSystemKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).ValidateAgentKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.v1.KeyService/ValidateAgentKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).ValidateAgentKey(ctx, req.(*ValidateSystemKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_CreateHookKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).CreateHookKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.v1.KeyService/CreateHookKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).CreateHookKeys(ctx, req.(*HooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_GetHookKeysForCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).GetHookKeysForCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.v1.KeyService/GetHookKeysForCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).GetHookKeysForCompany(ctx, req.(*HooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_ValidateHookKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateSystemKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).ValidateHookKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.v1.KeyService/ValidateHookKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).ValidateHookKey(ctx, req.(*ValidateSystemKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_CreateUserKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).CreateUserKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.v1.KeyService/CreateUserKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).CreateUserKeys(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyService_ValidateUserKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateUserKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyServiceServer).ValidateUserKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/key.v1.KeyService/ValidateUserKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyServiceServer).ValidateUserKeys(ctx, req.(*ValidateUserKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyService_ServiceDesc is the grpc.ServiceDesc for KeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "key.v1.KeyService",
	HandlerType: (*KeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgentKeys",
			Handler:    _KeyService_CreateAgentKeys_Handler,
		},
		{
			MethodName: "GetAgentKeys",
			Handler:    _KeyService_GetAgentKeys_Handler,
		},
		{
			MethodName: "ValidateAgentKey",
			Handler:    _KeyService_ValidateAgentKey_Handler,
		},
		{
			MethodName: "CreateHookKeys",
			Handler:    _KeyService_CreateHookKeys_Handler,
		},
		{
			MethodName: "GetHookKeysForCompany",
			Handler:    _KeyService_GetHookKeysForCompany_Handler,
		},
		{
			MethodName: "ValidateHookKey",
			Handler:    _KeyService_ValidateHookKey_Handler,
		},
		{
			MethodName: "CreateUserKeys",
			Handler:    _KeyService_CreateUserKeys_Handler,
		},
		{
			MethodName: "ValidateUserKeys",
			Handler:    _KeyService_ValidateUserKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/key.proto",
}
