// Code generated by protoc-gen-gogo.
// source: workerservice.proto
// DO NOT EDIT!

package echopb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import echopb_openstack1 "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WorkerService service

type WorkerServiceClient interface {
	StartPacketCapturing(ctx context.Context, in *echopb_openstack1.OSILayersReqRespData, opts ...grpc.CallOption) (*echopb_openstack1.OSILayersReqRespData, error)
	StopPacketCapturing(ctx context.Context, in *echopb_openstack1.OSILayersReqRespData, opts ...grpc.CallOption) (*echopb_openstack1.OSILayersReqRespData, error)
	GetLibvirtDomainInfo(ctx context.Context, in *echopb_openstack1.LibvirtDomainInfo, opts ...grpc.CallOption) (*echopb_openstack1.LibvirtDomainInfo, error)
	Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
}

type workerServiceClient struct {
	cc *grpc.ClientConn
}

func NewWorkerServiceClient(cc *grpc.ClientConn) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) StartPacketCapturing(ctx context.Context, in *echopb_openstack1.OSILayersReqRespData, opts ...grpc.CallOption) (*echopb_openstack1.OSILayersReqRespData, error) {
	out := new(echopb_openstack1.OSILayersReqRespData)
	err := grpc.Invoke(ctx, "/echopb.WorkerService/StartPacketCapturing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) StopPacketCapturing(ctx context.Context, in *echopb_openstack1.OSILayersReqRespData, opts ...grpc.CallOption) (*echopb_openstack1.OSILayersReqRespData, error) {
	out := new(echopb_openstack1.OSILayersReqRespData)
	err := grpc.Invoke(ctx, "/echopb.WorkerService/StopPacketCapturing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetLibvirtDomainInfo(ctx context.Context, in *echopb_openstack1.LibvirtDomainInfo, opts ...grpc.CallOption) (*echopb_openstack1.LibvirtDomainInfo, error) {
	out := new(echopb_openstack1.LibvirtDomainInfo)
	err := grpc.Invoke(ctx, "/echopb.WorkerService/GetLibvirtDomainInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := grpc.Invoke(ctx, "/echopb.WorkerService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkerService service

type WorkerServiceServer interface {
	StartPacketCapturing(context.Context, *echopb_openstack1.OSILayersReqRespData) (*echopb_openstack1.OSILayersReqRespData, error)
	StopPacketCapturing(context.Context, *echopb_openstack1.OSILayersReqRespData) (*echopb_openstack1.OSILayersReqRespData, error)
	GetLibvirtDomainInfo(context.Context, *echopb_openstack1.LibvirtDomainInfo) (*echopb_openstack1.LibvirtDomainInfo, error)
	Echo(context.Context, *EchoMessage) (*EchoMessage, error)
}

func RegisterWorkerServiceServer(s *grpc.Server, srv WorkerServiceServer) {
	s.RegisterService(&_WorkerService_serviceDesc, srv)
}

func _WorkerService_StartPacketCapturing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(echopb_openstack1.OSILayersReqRespData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).StartPacketCapturing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.WorkerService/StartPacketCapturing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).StartPacketCapturing(ctx, req.(*echopb_openstack1.OSILayersReqRespData))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_StopPacketCapturing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(echopb_openstack1.OSILayersReqRespData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).StopPacketCapturing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.WorkerService/StopPacketCapturing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).StopPacketCapturing(ctx, req.(*echopb_openstack1.OSILayersReqRespData))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetLibvirtDomainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(echopb_openstack1.LibvirtDomainInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetLibvirtDomainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.WorkerService/GetLibvirtDomainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetLibvirtDomainInfo(ctx, req.(*echopb_openstack1.LibvirtDomainInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.WorkerService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).Echo(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echopb.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartPacketCapturing",
			Handler:    _WorkerService_StartPacketCapturing_Handler,
		},
		{
			MethodName: "StopPacketCapturing",
			Handler:    _WorkerService_StopPacketCapturing_Handler,
		},
		{
			MethodName: "GetLibvirtDomainInfo",
			Handler:    _WorkerService_GetLibvirtDomainInfo_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _WorkerService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workerservice.proto",
}

func init() { proto.RegisterFile("workerservice.proto", fileDescriptorWorkerservice) }

var fileDescriptorWorkerservice = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0xd1, 0x31, 0x4a, 0xc4, 0x40,
	0x14, 0x06, 0xe0, 0xac, 0xc8, 0x16, 0x03, 0x0b, 0x32, 0x49, 0x35, 0x48, 0x8a, 0x08, 0x96, 0x09,
	0xac, 0x37, 0xd0, 0x15, 0x59, 0x58, 0x51, 0x36, 0x85, 0xf5, 0xcb, 0xf8, 0x9c, 0x0c, 0xd1, 0x79,
	0xe3, 0xcc, 0x73, 0xc5, 0x9b, 0x78, 0x24, 0x4b, 0x8f, 0x20, 0xf1, 0x16, 0x56, 0xe2, 0x06, 0x05,
	0x59, 0x8b, 0x6d, 0x2c, 0xe7, 0xe7, 0x9b, 0xff, 0x2f, 0x9e, 0x48, 0x1f, 0x29, 0x74, 0x18, 0x22,
	0x86, 0x95, 0xd5, 0x58, 0xfa, 0x40, 0x4c, 0x72, 0x8c, 0xba, 0x25, 0xdf, 0xa8, 0x7d, 0x43, 0x64,
	0x6e, 0xb1, 0x02, 0x6f, 0x2b, 0x70, 0x8e, 0x18, 0xd8, 0x92, 0x8b, 0x83, 0x52, 0x93, 0x5f, 0x9f,
	0x54, 0x46, 0x1e, 0x5d, 0x64, 0xd0, 0x5d, 0x75, 0x0d, 0x0c, 0x43, 0x3a, 0xfd, 0xd8, 0x11, 0x93,
	0xab, 0xf5, 0x44, 0x3d, 0x68, 0xd9, 0x8a, 0xac, 0x66, 0x08, 0x7c, 0x09, 0xba, 0x43, 0x3e, 0x01,
	0xcf, 0x0f, 0xc1, 0x3a, 0x23, 0x0f, 0xcb, 0x61, 0xb5, 0xfc, 0xe9, 0x29, 0x2f, 0xea, 0xf9, 0x02,
	0x9e, 0x30, 0xc4, 0x25, 0xde, 0x2f, 0x31, 0xfa, 0x19, 0x30, 0xa8, 0x2d, 0x5d, 0x91, 0x48, 0x23,
	0xd2, 0x9a, 0xc9, 0xff, 0xff, 0x50, 0x23, 0xb2, 0x33, 0xe4, 0x85, 0x6d, 0x56, 0x36, 0xf0, 0x8c,
	0xee, 0xc0, 0xba, 0xb9, 0xbb, 0x21, 0x79, 0xb0, 0xd9, 0xb0, 0x81, 0xd4, 0x36, 0xa8, 0x48, 0xe4,
	0x54, 0xec, 0x9e, 0xea, 0x96, 0x64, 0xfa, 0xcd, 0xbf, 0x5e, 0xe7, 0x18, 0x23, 0x18, 0x54, 0x7f,
	0x85, 0x45, 0x72, 0xbc, 0xf7, 0xd2, 0xe7, 0xa3, 0xd7, 0x3e, 0x1f, 0xbd, 0xf5, 0xf9, 0xe8, 0xf9,
	0x3d, 0x4f, 0x9a, 0xf1, 0xfa, 0x2a, 0x47, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0xfb, 0xac,
	0xd6, 0xf7, 0x01, 0x00, 0x00,
}
