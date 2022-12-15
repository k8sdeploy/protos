// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: v1/permissions.proto

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

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	// Modification
	AddPermission(ctx context.Context, in *PermissionItem, opts ...grpc.CallOption) (*PermissionResponse, error)
	AddPermissions(ctx context.Context, in *MultiplePermissionRequest, opts ...grpc.CallOption) (*PermissionResponse, error)
	RemovePermission(ctx context.Context, in *PermissionItem, opts ...grpc.CallOption) (*PermissionResponse, error)
	RemovePermissions(ctx context.Context, in *MultiplePermissionRequest, opts ...grpc.CallOption) (*PermissionResponse, error)
	// Retrieval
	GetPermissions(ctx context.Context, in *AllPermissionsGetRequest, opts ...grpc.CallOption) (*PermissionResponse, error)
	CanDo(ctx context.Context, in *PermissionItem, opts ...grpc.CallOption) (*PermissionResponse, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) AddPermission(ctx context.Context, in *PermissionItem, opts ...grpc.CallOption) (*PermissionResponse, error) {
	out := new(PermissionResponse)
	err := c.cc.Invoke(ctx, "/permissions.v1.PermissionService/AddPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) AddPermissions(ctx context.Context, in *MultiplePermissionRequest, opts ...grpc.CallOption) (*PermissionResponse, error) {
	out := new(PermissionResponse)
	err := c.cc.Invoke(ctx, "/permissions.v1.PermissionService/AddPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) RemovePermission(ctx context.Context, in *PermissionItem, opts ...grpc.CallOption) (*PermissionResponse, error) {
	out := new(PermissionResponse)
	err := c.cc.Invoke(ctx, "/permissions.v1.PermissionService/RemovePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) RemovePermissions(ctx context.Context, in *MultiplePermissionRequest, opts ...grpc.CallOption) (*PermissionResponse, error) {
	out := new(PermissionResponse)
	err := c.cc.Invoke(ctx, "/permissions.v1.PermissionService/RemovePermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetPermissions(ctx context.Context, in *AllPermissionsGetRequest, opts ...grpc.CallOption) (*PermissionResponse, error) {
	out := new(PermissionResponse)
	err := c.cc.Invoke(ctx, "/permissions.v1.PermissionService/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) CanDo(ctx context.Context, in *PermissionItem, opts ...grpc.CallOption) (*PermissionResponse, error) {
	out := new(PermissionResponse)
	err := c.cc.Invoke(ctx, "/permissions.v1.PermissionService/CanDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations must embed UnimplementedPermissionServiceServer
// for forward compatibility
type PermissionServiceServer interface {
	// Modification
	AddPermission(context.Context, *PermissionItem) (*PermissionResponse, error)
	AddPermissions(context.Context, *MultiplePermissionRequest) (*PermissionResponse, error)
	RemovePermission(context.Context, *PermissionItem) (*PermissionResponse, error)
	RemovePermissions(context.Context, *MultiplePermissionRequest) (*PermissionResponse, error)
	// Retrieval
	GetPermissions(context.Context, *AllPermissionsGetRequest) (*PermissionResponse, error)
	CanDo(context.Context, *PermissionItem) (*PermissionResponse, error)
	mustEmbedUnimplementedPermissionServiceServer()
}

// UnimplementedPermissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionServiceServer struct {
}

func (UnimplementedPermissionServiceServer) AddPermission(context.Context, *PermissionItem) (*PermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermission not implemented")
}
func (UnimplementedPermissionServiceServer) AddPermissions(context.Context, *MultiplePermissionRequest) (*PermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermissions not implemented")
}
func (UnimplementedPermissionServiceServer) RemovePermission(context.Context, *PermissionItem) (*PermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermission not implemented")
}
func (UnimplementedPermissionServiceServer) RemovePermissions(context.Context, *MultiplePermissionRequest) (*PermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermissions not implemented")
}
func (UnimplementedPermissionServiceServer) GetPermissions(context.Context, *AllPermissionsGetRequest) (*PermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (UnimplementedPermissionServiceServer) CanDo(context.Context, *PermissionItem) (*PermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanDo not implemented")
}
func (UnimplementedPermissionServiceServer) mustEmbedUnimplementedPermissionServiceServer() {}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_AddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).AddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.v1.PermissionService/AddPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).AddPermission(ctx, req.(*PermissionItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_AddPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).AddPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.v1.PermissionService/AddPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).AddPermissions(ctx, req.(*MultiplePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_RemovePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).RemovePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.v1.PermissionService/RemovePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).RemovePermission(ctx, req.(*PermissionItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_RemovePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).RemovePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.v1.PermissionService/RemovePermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).RemovePermissions(ctx, req.(*MultiplePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllPermissionsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.v1.PermissionService/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetPermissions(ctx, req.(*AllPermissionsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_CanDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).CanDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.v1.PermissionService/CanDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).CanDo(ctx, req.(*PermissionItem))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "permissions.v1.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPermission",
			Handler:    _PermissionService_AddPermission_Handler,
		},
		{
			MethodName: "AddPermissions",
			Handler:    _PermissionService_AddPermissions_Handler,
		},
		{
			MethodName: "RemovePermission",
			Handler:    _PermissionService_RemovePermission_Handler,
		},
		{
			MethodName: "RemovePermissions",
			Handler:    _PermissionService_RemovePermissions_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _PermissionService_GetPermissions_Handler,
		},
		{
			MethodName: "CanDo",
			Handler:    _PermissionService_CanDo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/permissions.proto",
}
