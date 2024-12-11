// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: relationships/v1/subject_ispropagatedto_object_service.proto

package relationships

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ObjectSubjectRelationshipService_CreateObjectSubjectRelationship_FullMethodName = "/relationships.v1.ObjectSubjectRelationshipService/CreateObjectSubjectRelationship"
	ObjectSubjectRelationshipService_UpdateObjectSubjectRelationship_FullMethodName = "/relationships.v1.ObjectSubjectRelationshipService/UpdateObjectSubjectRelationship"
	ObjectSubjectRelationshipService_DeleteObjectSubjectRelationship_FullMethodName = "/relationships.v1.ObjectSubjectRelationshipService/DeleteObjectSubjectRelationship"
)

// ObjectSubjectRelationshipServiceClient is the client API for ObjectSubjectRelationshipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectSubjectRelationshipServiceClient interface {
	CreateObjectSubjectRelationship(ctx context.Context, in *CreateObjectSubjectRelationshipRequest, opts ...grpc.CallOption) (*CreateObjectSubjectRelationshipResponse, error)
	UpdateObjectSubjectRelationship(ctx context.Context, in *UpdateObjectSubjectRelationshipRequest, opts ...grpc.CallOption) (*UpdateObjectSubjectRelationshipResponse, error)
	DeleteObjectSubjectRelationship(ctx context.Context, in *DeleteObjectSubjectRelationshipRequest, opts ...grpc.CallOption) (*DeleteObjectSubjectRelationshipResponse, error)
}

type objectSubjectRelationshipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectSubjectRelationshipServiceClient(cc grpc.ClientConnInterface) ObjectSubjectRelationshipServiceClient {
	return &objectSubjectRelationshipServiceClient{cc}
}

func (c *objectSubjectRelationshipServiceClient) CreateObjectSubjectRelationship(ctx context.Context, in *CreateObjectSubjectRelationshipRequest, opts ...grpc.CallOption) (*CreateObjectSubjectRelationshipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateObjectSubjectRelationshipResponse)
	err := c.cc.Invoke(ctx, ObjectSubjectRelationshipService_CreateObjectSubjectRelationship_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectSubjectRelationshipServiceClient) UpdateObjectSubjectRelationship(ctx context.Context, in *UpdateObjectSubjectRelationshipRequest, opts ...grpc.CallOption) (*UpdateObjectSubjectRelationshipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateObjectSubjectRelationshipResponse)
	err := c.cc.Invoke(ctx, ObjectSubjectRelationshipService_UpdateObjectSubjectRelationship_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectSubjectRelationshipServiceClient) DeleteObjectSubjectRelationship(ctx context.Context, in *DeleteObjectSubjectRelationshipRequest, opts ...grpc.CallOption) (*DeleteObjectSubjectRelationshipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteObjectSubjectRelationshipResponse)
	err := c.cc.Invoke(ctx, ObjectSubjectRelationshipService_DeleteObjectSubjectRelationship_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectSubjectRelationshipServiceServer is the server API for ObjectSubjectRelationshipService service.
// All implementations must embed UnimplementedObjectSubjectRelationshipServiceServer
// for forward compatibility.
type ObjectSubjectRelationshipServiceServer interface {
	CreateObjectSubjectRelationship(context.Context, *CreateObjectSubjectRelationshipRequest) (*CreateObjectSubjectRelationshipResponse, error)
	UpdateObjectSubjectRelationship(context.Context, *UpdateObjectSubjectRelationshipRequest) (*UpdateObjectSubjectRelationshipResponse, error)
	DeleteObjectSubjectRelationship(context.Context, *DeleteObjectSubjectRelationshipRequest) (*DeleteObjectSubjectRelationshipResponse, error)
	mustEmbedUnimplementedObjectSubjectRelationshipServiceServer()
}

// UnimplementedObjectSubjectRelationshipServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedObjectSubjectRelationshipServiceServer struct{}

func (UnimplementedObjectSubjectRelationshipServiceServer) CreateObjectSubjectRelationship(context.Context, *CreateObjectSubjectRelationshipRequest) (*CreateObjectSubjectRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectSubjectRelationship not implemented")
}
func (UnimplementedObjectSubjectRelationshipServiceServer) UpdateObjectSubjectRelationship(context.Context, *UpdateObjectSubjectRelationshipRequest) (*UpdateObjectSubjectRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObjectSubjectRelationship not implemented")
}
func (UnimplementedObjectSubjectRelationshipServiceServer) DeleteObjectSubjectRelationship(context.Context, *DeleteObjectSubjectRelationshipRequest) (*DeleteObjectSubjectRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjectSubjectRelationship not implemented")
}
func (UnimplementedObjectSubjectRelationshipServiceServer) mustEmbedUnimplementedObjectSubjectRelationshipServiceServer() {
}
func (UnimplementedObjectSubjectRelationshipServiceServer) testEmbeddedByValue() {}

// UnsafeObjectSubjectRelationshipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectSubjectRelationshipServiceServer will
// result in compilation errors.
type UnsafeObjectSubjectRelationshipServiceServer interface {
	mustEmbedUnimplementedObjectSubjectRelationshipServiceServer()
}

func RegisterObjectSubjectRelationshipServiceServer(s grpc.ServiceRegistrar, srv ObjectSubjectRelationshipServiceServer) {
	// If the following call pancis, it indicates UnimplementedObjectSubjectRelationshipServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ObjectSubjectRelationshipService_ServiceDesc, srv)
}

func _ObjectSubjectRelationshipService_CreateObjectSubjectRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectSubjectRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectSubjectRelationshipServiceServer).CreateObjectSubjectRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectSubjectRelationshipService_CreateObjectSubjectRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectSubjectRelationshipServiceServer).CreateObjectSubjectRelationship(ctx, req.(*CreateObjectSubjectRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectSubjectRelationshipService_UpdateObjectSubjectRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObjectSubjectRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectSubjectRelationshipServiceServer).UpdateObjectSubjectRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectSubjectRelationshipService_UpdateObjectSubjectRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectSubjectRelationshipServiceServer).UpdateObjectSubjectRelationship(ctx, req.(*UpdateObjectSubjectRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectSubjectRelationshipService_DeleteObjectSubjectRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectSubjectRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectSubjectRelationshipServiceServer).DeleteObjectSubjectRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectSubjectRelationshipService_DeleteObjectSubjectRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectSubjectRelationshipServiceServer).DeleteObjectSubjectRelationship(ctx, req.(*DeleteObjectSubjectRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectSubjectRelationshipService_ServiceDesc is the grpc.ServiceDesc for ObjectSubjectRelationshipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectSubjectRelationshipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relationships.v1.ObjectSubjectRelationshipService",
	HandlerType: (*ObjectSubjectRelationshipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjectSubjectRelationship",
			Handler:    _ObjectSubjectRelationshipService_CreateObjectSubjectRelationship_Handler,
		},
		{
			MethodName: "UpdateObjectSubjectRelationship",
			Handler:    _ObjectSubjectRelationshipService_UpdateObjectSubjectRelationship_Handler,
		},
		{
			MethodName: "DeleteObjectSubjectRelationship",
			Handler:    _ObjectSubjectRelationshipService_DeleteObjectSubjectRelationship_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relationships/v1/subject_ispropagatedto_object_service.proto",
}