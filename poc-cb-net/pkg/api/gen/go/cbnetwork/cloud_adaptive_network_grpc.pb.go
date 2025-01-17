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

// CloudAdaptiveNetworkServiceClient is the client API for CloudAdaptiveNetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudAdaptiveNetworkServiceClient interface {
	// Used to say hello (for testing).
	// Pass in nothing and
	// return a say-hello message.
	SayHello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	// Used to get a Cloud Adaptive Network specification.
	// Pass in an ID of Cloud Adaptive Networkand
	// return a Cloud Adaptive Network specification.
	GetCLADNet(ctx context.Context, in *CLADNetID, opts ...grpc.CallOption) (*CLADNetSpecification, error)
	// Used to get a list of Cloud Adaptive Network specifications.
	// Pass in nothing and
	// return a list of Cloud Adaptive Network specifications.
	GetCLADNetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CLADNetSpecifications, error)
	// Used to create a new Cloud Adaptive Network.
	// Pass in a specification of Cloud Adaptive Network and
	// return the specification of Cloud Adaptive Network.
	CreateCLADNet(ctx context.Context, in *CLADNetSpecification, opts ...grpc.CallOption) (*CLADNetSpecification, error)
	// Used to recommend available IPv4 private address spaces for Cloud Adaptive Network.
	// Pass in a list of IP networks (e.g., ["10.10.10.10/14", "192.168.20.20/26", ....]) and
	// return available IPv4 private address spaces
	RecommendAvailableIPv4PrivateAddressSpaces(ctx context.Context, in *IPNetworks, opts ...grpc.CallOption) (*AvailableIPv4PrivateAddressSpaces, error)
	// [To be provided] Used to delete a Cloud Adaptive Network
	// Pass in an ID of Cloud Adaptive Network and
	// return a result of attempt to delete a Cloud Adaptive Network.
	DeleteCLADNet(ctx context.Context, in *CLADNetID, opts ...grpc.CallOption) (*DeletionResult, error)
	// [To be provided] Used to update a Cloud Adaptive Network
	// Pass in a specification of Cloud Adaptive Network and
	// return the specification of Cloud Adaptive Network.
	UpdateCLADNet(ctx context.Context, in *CLADNetSpecification, opts ...grpc.CallOption) (*CLADNetSpecification, error)
}

type cloudAdaptiveNetworkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudAdaptiveNetworkServiceClient(cc grpc.ClientConnInterface) CloudAdaptiveNetworkServiceClient {
	return &cloudAdaptiveNetworkServiceClient{cc}
}

func (c *cloudAdaptiveNetworkServiceClient) SayHello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetworkService/sayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkServiceClient) GetCLADNet(ctx context.Context, in *CLADNetID, opts ...grpc.CallOption) (*CLADNetSpecification, error) {
	out := new(CLADNetSpecification)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetworkService/getCLADNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkServiceClient) GetCLADNetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CLADNetSpecifications, error) {
	out := new(CLADNetSpecifications)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetworkService/getCLADNetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkServiceClient) CreateCLADNet(ctx context.Context, in *CLADNetSpecification, opts ...grpc.CallOption) (*CLADNetSpecification, error) {
	out := new(CLADNetSpecification)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetworkService/createCLADNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkServiceClient) RecommendAvailableIPv4PrivateAddressSpaces(ctx context.Context, in *IPNetworks, opts ...grpc.CallOption) (*AvailableIPv4PrivateAddressSpaces, error) {
	out := new(AvailableIPv4PrivateAddressSpaces)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetworkService/recommendAvailableIPv4PrivateAddressSpaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkServiceClient) DeleteCLADNet(ctx context.Context, in *CLADNetID, opts ...grpc.CallOption) (*DeletionResult, error) {
	out := new(DeletionResult)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetworkService/deleteCLADNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudAdaptiveNetworkServiceClient) UpdateCLADNet(ctx context.Context, in *CLADNetSpecification, opts ...grpc.CallOption) (*CLADNetSpecification, error) {
	out := new(CLADNetSpecification)
	err := c.cc.Invoke(ctx, "/cbnet.CloudAdaptiveNetworkService/updateCLADNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudAdaptiveNetworkServiceServer is the server API for CloudAdaptiveNetworkService service.
// All implementations must embed UnimplementedCloudAdaptiveNetworkServiceServer
// for forward compatibility
type CloudAdaptiveNetworkServiceServer interface {
	// Used to say hello (for testing).
	// Pass in nothing and
	// return a say-hello message.
	SayHello(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error)
	// Used to get a Cloud Adaptive Network specification.
	// Pass in an ID of Cloud Adaptive Networkand
	// return a Cloud Adaptive Network specification.
	GetCLADNet(context.Context, *CLADNetID) (*CLADNetSpecification, error)
	// Used to get a list of Cloud Adaptive Network specifications.
	// Pass in nothing and
	// return a list of Cloud Adaptive Network specifications.
	GetCLADNetList(context.Context, *emptypb.Empty) (*CLADNetSpecifications, error)
	// Used to create a new Cloud Adaptive Network.
	// Pass in a specification of Cloud Adaptive Network and
	// return the specification of Cloud Adaptive Network.
	CreateCLADNet(context.Context, *CLADNetSpecification) (*CLADNetSpecification, error)
	// Used to recommend available IPv4 private address spaces for Cloud Adaptive Network.
	// Pass in a list of IP networks (e.g., ["10.10.10.10/14", "192.168.20.20/26", ....]) and
	// return available IPv4 private address spaces
	RecommendAvailableIPv4PrivateAddressSpaces(context.Context, *IPNetworks) (*AvailableIPv4PrivateAddressSpaces, error)
	// [To be provided] Used to delete a Cloud Adaptive Network
	// Pass in an ID of Cloud Adaptive Network and
	// return a result of attempt to delete a Cloud Adaptive Network.
	DeleteCLADNet(context.Context, *CLADNetID) (*DeletionResult, error)
	// [To be provided] Used to update a Cloud Adaptive Network
	// Pass in a specification of Cloud Adaptive Network and
	// return the specification of Cloud Adaptive Network.
	UpdateCLADNet(context.Context, *CLADNetSpecification) (*CLADNetSpecification, error)
	mustEmbedUnimplementedCloudAdaptiveNetworkServiceServer()
}

// UnimplementedCloudAdaptiveNetworkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCloudAdaptiveNetworkServiceServer struct {
}

func (UnimplementedCloudAdaptiveNetworkServiceServer) SayHello(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServiceServer) GetCLADNet(context.Context, *CLADNetID) (*CLADNetSpecification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCLADNet not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServiceServer) GetCLADNetList(context.Context, *emptypb.Empty) (*CLADNetSpecifications, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCLADNetList not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServiceServer) CreateCLADNet(context.Context, *CLADNetSpecification) (*CLADNetSpecification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCLADNet not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServiceServer) RecommendAvailableIPv4PrivateAddressSpaces(context.Context, *IPNetworks) (*AvailableIPv4PrivateAddressSpaces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendAvailableIPv4PrivateAddressSpaces not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServiceServer) DeleteCLADNet(context.Context, *CLADNetID) (*DeletionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCLADNet not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServiceServer) UpdateCLADNet(context.Context, *CLADNetSpecification) (*CLADNetSpecification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCLADNet not implemented")
}
func (UnimplementedCloudAdaptiveNetworkServiceServer) mustEmbedUnimplementedCloudAdaptiveNetworkServiceServer() {
}

// UnsafeCloudAdaptiveNetworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudAdaptiveNetworkServiceServer will
// result in compilation errors.
type UnsafeCloudAdaptiveNetworkServiceServer interface {
	mustEmbedUnimplementedCloudAdaptiveNetworkServiceServer()
}

func RegisterCloudAdaptiveNetworkServiceServer(s grpc.ServiceRegistrar, srv CloudAdaptiveNetworkServiceServer) {
	s.RegisterService(&CloudAdaptiveNetworkService_ServiceDesc, srv)
}

func _CloudAdaptiveNetworkService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetworkService/sayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServiceServer).SayHello(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetworkService_GetCLADNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLADNetID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServiceServer).GetCLADNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetworkService/getCLADNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServiceServer).GetCLADNet(ctx, req.(*CLADNetID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetworkService_GetCLADNetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServiceServer).GetCLADNetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetworkService/getCLADNetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServiceServer).GetCLADNetList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetworkService_CreateCLADNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLADNetSpecification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServiceServer).CreateCLADNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetworkService/createCLADNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServiceServer).CreateCLADNet(ctx, req.(*CLADNetSpecification))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetworkService_RecommendAvailableIPv4PrivateAddressSpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPNetworks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServiceServer).RecommendAvailableIPv4PrivateAddressSpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetworkService/recommendAvailableIPv4PrivateAddressSpaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServiceServer).RecommendAvailableIPv4PrivateAddressSpaces(ctx, req.(*IPNetworks))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetworkService_DeleteCLADNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLADNetID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServiceServer).DeleteCLADNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetworkService/deleteCLADNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServiceServer).DeleteCLADNet(ctx, req.(*CLADNetID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudAdaptiveNetworkService_UpdateCLADNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLADNetSpecification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudAdaptiveNetworkServiceServer).UpdateCLADNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbnet.CloudAdaptiveNetworkService/updateCLADNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudAdaptiveNetworkServiceServer).UpdateCLADNet(ctx, req.(*CLADNetSpecification))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudAdaptiveNetworkService_ServiceDesc is the grpc.ServiceDesc for CloudAdaptiveNetworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudAdaptiveNetworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cbnet.CloudAdaptiveNetworkService",
	HandlerType: (*CloudAdaptiveNetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayHello",
			Handler:    _CloudAdaptiveNetworkService_SayHello_Handler,
		},
		{
			MethodName: "getCLADNet",
			Handler:    _CloudAdaptiveNetworkService_GetCLADNet_Handler,
		},
		{
			MethodName: "getCLADNetList",
			Handler:    _CloudAdaptiveNetworkService_GetCLADNetList_Handler,
		},
		{
			MethodName: "createCLADNet",
			Handler:    _CloudAdaptiveNetworkService_CreateCLADNet_Handler,
		},
		{
			MethodName: "recommendAvailableIPv4PrivateAddressSpaces",
			Handler:    _CloudAdaptiveNetworkService_RecommendAvailableIPv4PrivateAddressSpaces_Handler,
		},
		{
			MethodName: "deleteCLADNet",
			Handler:    _CloudAdaptiveNetworkService_DeleteCLADNet_Handler,
		},
		{
			MethodName: "updateCLADNet",
			Handler:    _CloudAdaptiveNetworkService_UpdateCLADNet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cbnetwork/cloud_adaptive_network.proto",
}
