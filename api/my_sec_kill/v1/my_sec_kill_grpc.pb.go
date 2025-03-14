// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: my_sec_kill/v1/my_sec_kill.proto

package v1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

// UserXClient is the client API for UserX service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserXClient interface {
}

type userXClient struct {
	cc grpc.ClientConnInterface
}

func NewUserXClient(cc grpc.ClientConnInterface) UserXClient {
	return &userXClient{cc}
}

// UserXServer is the server API for UserX service.
// All implementations must embed UnimplementedUserXServer
// for forward compatibility.
type UserXServer interface {
	mustEmbedUnimplementedUserXServer()
}

// UnimplementedUserXServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserXServer struct{}

func (UnimplementedUserXServer) mustEmbedUnimplementedUserXServer() {}
func (UnimplementedUserXServer) testEmbeddedByValue()               {}

// UnsafeUserXServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserXServer will
// result in compilation errors.
type UnsafeUserXServer interface {
	mustEmbedUnimplementedUserXServer()
}

func RegisterUserXServer(s grpc.ServiceRegistrar, srv UserXServer) {
	// If the following call pancis, it indicates UnimplementedUserXServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserX_ServiceDesc, srv)
}

// UserX_ServiceDesc is the grpc.ServiceDesc for UserX service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserX_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sec_kell.v1.UserX",
	HandlerType: (*UserXServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "my_sec_kill/v1/my_sec_kill.proto",
}
