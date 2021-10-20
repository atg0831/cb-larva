// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cb_larva

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CloudAdaptiveNetworkClient is the client API for CloudAdaptiveNetwork service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudAdaptiveNetworkClient interface {
	// Return Say Hello
	SayHello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	// Return a specific CLADNet
	GetCLADNet(ctx context.Context, in *CLADNetID, opts ...grpc.CallOption) (*CLADNetSpecification, error)
	// Return a specific CLADNet
	GetCLADNetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CLADNetSpecifications, error)
	// Create a new CLADNet
	CreateCLADNet(ctx context.Context, in *CLADNetSpecification, opts ...grpc.CallOption) (*CLADNetSpecification, error)
	// Returns available IPv4 private address spaces
	RecommendAvailableIPv4PrivateAddressSpaces(ctx context.Context, in *IPNetworks, opts ...grpc.CallOption) (*AvailableIPv4PrivateAddressSpaces, error)
	// Return a result of deleting a specific CLADNet
	DeleteCLADNet(ctx context.Context, in *CLADNetID, opts ...grpc.CallOption) (*DeletionResult, error)
	// Return a updated CLADNet
	UpdateCLADNet(ctx context.Context, in *CLADNetSpecification, opts ...grpc.CallOption) (*CLADNetSpecification, error)
}

type cloudAdaptiveNetworkClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudAdaptiveNetworkClient(cc grpc.ClientConnInterface) CloudAdaptiveNetworkClient {
	return &cloudAdaptiveNetworkClient{cc}
}

func (c *cloudAdaptiveNetworkClient) SayHello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetwork/sayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkClient) GetCLADNet(ctx context.Context, in *CLADNetID, opts ...grpc.CallOption) (*CLADNetSpecification, error) {
	out := new(CLADNetSpecification)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetwork/getCLADNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkClient) GetCLADNetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CLADNetSpecifications, error) {
	out := new(CLADNetSpecifications)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetwork/getCLADNetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkClient) CreateCLADNet(ctx context.Context, in *CLADNetSpecification, opts ...grpc.CallOption) (*CLADNetSpecification, error) {
	out := new(CLADNetSpecification)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetwork/createCLADNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkClient) RecommendAvailableIPv4PrivateAddressSpaces(ctx context.Context, in *IPNetworks, opts ...grpc.CallOption) (*AvailableIPv4PrivateAddressSpaces, error) {
	out := new(AvailableIPv4PrivateAddressSpaces)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetwork/recommendAvailableIPv4PrivateAddressSpaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkClient) DeleteCLADNet(ctx context.Context, in *CLADNetID, opts ...grpc.CallOption) (*DeletionResult, error) {
	out := new(DeletionResult)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetwork/deleteCLADNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkClient) UpdateCLADNet(ctx context.Context, in *CLADNetSpecification, opts ...grpc.CallOption) (*CLADNetSpecification, error) {
	out := new(CLADNetSpecification)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetwork/updateCLADNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudAdaptiveNetworkServer is the server API for CloudAdaptiveNetwork service.
// All implementations must embed UnimplementedCloudAdaptiveNetworkServer
// for forward compatibility
type CloudAdaptiveNetworkServer interface {
	// Return Say Hello
	SayHello(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error)
	// Return a specific CLADNet
	GetCLADNet(context.Context, *CLADNetID) (*CLADNetSpecification, error)
	// Return a specific CLADNet
	GetCLADNetList(context.Context, *emptypb.Empty) (*CLADNetSpecifications, error)
	// Create a new CLADNet
	CreateCLADNet(context.Context, *CLADNetSpecification) (*CLADNetSpecification, error)
	// Returns available IPv4 private address spaces
	RecommendAvailableIPv4PrivateAddressSpaces(context.Context, *IPNetworks) (*AvailableIPv4PrivateAddressSpaces, error)
	// Return a result of deleting a specific CLADNet
	DeleteCLADNet(context.Context, *CLADNetID) (*DeletionResult, error)
	// Return a updated CLADNet
	UpdateCLADNet(context.Context, *CLADNetSpecification) (*CLADNetSpecification, error)
	mustEmbedUnimplementedCloudAdaptiveNetworkServer()
}

// UnimplementedCloudAdaptiveNetworkServer must be embedded to have forward compatible implementations.
type UnimplementedCloudAdaptiveNetworkServer struct {
}

func (UnimplementedCloudAdaptiveNetworkServer) SayHello(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServer) GetCLADNet(context.Context, *CLADNetID) (*CLADNetSpecification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCLADNet not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServer) GetCLADNetList(context.Context, *emptypb.Empty) (*CLADNetSpecifications, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCLADNetList not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServer) CreateCLADNet(context.Context, *CLADNetSpecification) (*CLADNetSpecification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCLADNet not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServer) RecommendAvailableIPv4PrivateAddressSpaces(context.Context, *IPNetworks) (*AvailableIPv4PrivateAddressSpaces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendAvailableIPv4PrivateAddressSpaces not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServer) DeleteCLADNet(context.Context, *CLADNetID) (*DeletionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCLADNet not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServer) UpdateCLADNet(context.Context, *CLADNetSpecification) (*CLADNetSpecification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCLADNet not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServer) mustEmbedUnimplementedCloudAdaptiveNetworkServer() {}

// UnsafeCloudAdaptiveNetworkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudAdaptiveNetworkServer will
// result in compilation errors.
type UnsafeCloudAdaptiveNetworkServer interface {
	mustEmbedUnimplementedCloudAdaptiveNetworkServer()
}

func RegisterCloudAdaptiveNetworkServer(s grpc.ServiceRegistrar, srv CloudAdaptiveNetworkServer) {
	s.RegisterService(&CloudAdaptiveNetwork_ServiceDesc, srv)
}

func _CloudAdaptiveNetwork_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetwork/sayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServer).SayHello(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetwork_GetCLADNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLADNetID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServer).GetCLADNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetwork/getCLADNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServer).GetCLADNet(ctx, req.(*CLADNetID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetwork_GetCLADNetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServer).GetCLADNetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetwork/getCLADNetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServer).GetCLADNetList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetwork_CreateCLADNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLADNetSpecification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServer).CreateCLADNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetwork/createCLADNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServer).CreateCLADNet(ctx, req.(*CLADNetSpecification))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetwork_RecommendAvailableIPv4PrivateAddressSpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPNetworks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServer).RecommendAvailableIPv4PrivateAddressSpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetwork/recommendAvailableIPv4PrivateAddressSpaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServer).RecommendAvailableIPv4PrivateAddressSpaces(ctx, req.(*IPNetworks))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetwork_DeleteCLADNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLADNetID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServer).DeleteCLADNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetwork/deleteCLADNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServer).DeleteCLADNet(ctx, req.(*CLADNetID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetwork_UpdateCLADNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLADNetSpecification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServer).UpdateCLADNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetwork/updateCLADNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServer).UpdateCLADNet(ctx, req.(*CLADNetSpecification))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudAdaptiveNetwork_ServiceDesc is the grpc.ServiceDesc for CloudAdaptiveNetwork service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudAdaptiveNetwork_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cbnet.CloudAdaptiveNetwork",
	HandlerType: (*CloudAdaptiveNetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayHello",
			Handler:    _CloudAdaptiveNetwork_SayHello_Handler,
		},
		{
			MethodName: "getCLADNet",
			Handler:    _CloudAdaptiveNetwork_GetCLADNet_Handler,
		},
		{
			MethodName: "getCLADNetList",
			Handler:    _CloudAdaptiveNetwork_GetCLADNetList_Handler,
		},
		{
			MethodName: "createCLADNet",
			Handler:    _CloudAdaptiveNetwork_CreateCLADNet_Handler,
		},
		{
			MethodName: "recommendAvailableIPv4PrivateAddressSpaces",
			Handler:    _CloudAdaptiveNetwork_RecommendAvailableIPv4PrivateAddressSpaces_Handler,
		},
		{
			MethodName: "deleteCLADNet",
			Handler:    _CloudAdaptiveNetwork_DeleteCLADNet_Handler,
		},
		{
			MethodName: "updateCLADNet",
			Handler:    _CloudAdaptiveNetwork_UpdateCLADNet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cbnetwork/cloud_adaptive_network.proto",
}