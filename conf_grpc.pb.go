// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package configurator

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

// ConfiguratorClient is the client API for Configurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfiguratorClient interface {
	Get(ctx context.Context, in *Conf, opts ...grpc.CallOption) (*Conf, error)
	SetAdminConf(ctx context.Context, in *AdminConf, opts ...grpc.CallOption) (*AdminConf, error)
}

type configuratorClient struct {
	cc grpc.ClientConnInterface
}

func NewConfiguratorClient(cc grpc.ClientConnInterface) ConfiguratorClient {
	return &configuratorClient{cc}
}

func (c *configuratorClient) Get(ctx context.Context, in *Conf, opts ...grpc.CallOption) (*Conf, error) {
	out := new(Conf)
	err := c.cc.Invoke(ctx, "/configurator.configurator/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configuratorClient) SetAdminConf(ctx context.Context, in *AdminConf, opts ...grpc.CallOption) (*AdminConf, error) {
	out := new(AdminConf)
	err := c.cc.Invoke(ctx, "/configurator.configurator/SetAdminConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfiguratorServer is the server API for Configurator service.
// All implementations must embed UnimplementedConfiguratorServer
// for forward compatibility
type ConfiguratorServer interface {
	Get(context.Context, *Conf) (*Conf, error)
	SetAdminConf(context.Context, *AdminConf) (*AdminConf, error)
	mustEmbedUnimplementedConfiguratorServer()
}

// UnimplementedConfiguratorServer must be embedded to have forward compatible implementations.
type UnimplementedConfiguratorServer struct {
}

func (UnimplementedConfiguratorServer) Get(context.Context, *Conf) (*Conf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedConfiguratorServer) SetAdminConf(context.Context, *AdminConf) (*AdminConf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAdminConf not implemented")
}
func (UnimplementedConfiguratorServer) mustEmbedUnimplementedConfiguratorServer() {}

// UnsafeConfiguratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfiguratorServer will
// result in compilation errors.
type UnsafeConfiguratorServer interface {
	mustEmbedUnimplementedConfiguratorServer()
}

func RegisterConfiguratorServer(s grpc.ServiceRegistrar, srv ConfiguratorServer) {
	s.RegisterService(&Configurator_ServiceDesc, srv)
}

func _Configurator_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configurator.configurator/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).Get(ctx, req.(*Conf))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configurator_SetAdminConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).SetAdminConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configurator.configurator/SetAdminConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).SetAdminConf(ctx, req.(*AdminConf))
	}
	return interceptor(ctx, in, info, handler)
}

// Configurator_ServiceDesc is the grpc.ServiceDesc for Configurator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Configurator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "configurator.configurator",
	HandlerType: (*ConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Configurator_Get_Handler,
		},
		{
			MethodName: "SetAdminConf",
			Handler:    _Configurator_SetAdminConf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conf.proto",
}
