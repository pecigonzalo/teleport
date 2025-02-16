// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: teleport/trust/v1/trust_service.proto

package trustv1

import (
	context "context"
	types "github.com/gravitational/teleport/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TrustService_GetCertAuthority_FullMethodName            = "/teleport.trust.v1.TrustService/GetCertAuthority"
	TrustService_GetCertAuthorities_FullMethodName          = "/teleport.trust.v1.TrustService/GetCertAuthorities"
	TrustService_DeleteCertAuthority_FullMethodName         = "/teleport.trust.v1.TrustService/DeleteCertAuthority"
	TrustService_UpsertCertAuthority_FullMethodName         = "/teleport.trust.v1.TrustService/UpsertCertAuthority"
	TrustService_RotateCertAuthority_FullMethodName         = "/teleport.trust.v1.TrustService/RotateCertAuthority"
	TrustService_RotateExternalCertAuthority_FullMethodName = "/teleport.trust.v1.TrustService/RotateExternalCertAuthority"
	TrustService_GenerateHostCert_FullMethodName            = "/teleport.trust.v1.TrustService/GenerateHostCert"
	TrustService_UpsertTrustedCluster_FullMethodName        = "/teleport.trust.v1.TrustService/UpsertTrustedCluster"
	TrustService_CreateTrustedCluster_FullMethodName        = "/teleport.trust.v1.TrustService/CreateTrustedCluster"
	TrustService_UpdateTrustedCluster_FullMethodName        = "/teleport.trust.v1.TrustService/UpdateTrustedCluster"
)

// TrustServiceClient is the client API for TrustService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// TrustService provides methods to manage certificate authorities.
type TrustServiceClient interface {
	// GetCertAuthority returns a cert authority by type and domain.
	GetCertAuthority(ctx context.Context, in *GetCertAuthorityRequest, opts ...grpc.CallOption) (*types.CertAuthorityV2, error)
	// GetCertAuthorities returns all cert authorities with the specified type.
	GetCertAuthorities(ctx context.Context, in *GetCertAuthoritiesRequest, opts ...grpc.CallOption) (*GetCertAuthoritiesResponse, error)
	// DeleteCertAuthority deletes the matching cert authority.
	DeleteCertAuthority(ctx context.Context, in *DeleteCertAuthorityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UpsertCertAuthority creates or updates the provided cert authority.
	UpsertCertAuthority(ctx context.Context, in *UpsertCertAuthorityRequest, opts ...grpc.CallOption) (*types.CertAuthorityV2, error)
	// RotateCertAuthority is a request to start rotation of the certificate authority.
	RotateCertAuthority(ctx context.Context, in *RotateCertAuthorityRequest, opts ...grpc.CallOption) (*RotateCertAuthorityResponse, error)
	// RotateExternalCertAuthority rotates an external cert authority.
	RotateExternalCertAuthority(ctx context.Context, in *RotateExternalCertAuthorityRequest, opts ...grpc.CallOption) (*RotateExternalCertAuthorityResponse, error)
	// GenerateHostCert takes a public key in the OpenSSH `authorized_keys` format and returns
	// a SSH certificate signed by the Host CA.
	GenerateHostCert(ctx context.Context, in *GenerateHostCertRequest, opts ...grpc.CallOption) (*GenerateHostCertResponse, error)
	// UpsertTrustedCluster upserts a Trusted Cluster in a backend.
	UpsertTrustedCluster(ctx context.Context, in *UpsertTrustedClusterRequest, opts ...grpc.CallOption) (*types.TrustedClusterV2, error)
	// CreateTrustedCluster creates a Trusted Cluster in a backend.
	CreateTrustedCluster(ctx context.Context, in *CreateTrustedClusterRequest, opts ...grpc.CallOption) (*types.TrustedClusterV2, error)
	// UpdateTrustedCluster updates a Trusted Cluster in a backend.
	UpdateTrustedCluster(ctx context.Context, in *UpdateTrustedClusterRequest, opts ...grpc.CallOption) (*types.TrustedClusterV2, error)
}

type trustServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrustServiceClient(cc grpc.ClientConnInterface) TrustServiceClient {
	return &trustServiceClient{cc}
}

func (c *trustServiceClient) GetCertAuthority(ctx context.Context, in *GetCertAuthorityRequest, opts ...grpc.CallOption) (*types.CertAuthorityV2, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.CertAuthorityV2)
	err := c.cc.Invoke(ctx, TrustService_GetCertAuthority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustServiceClient) GetCertAuthorities(ctx context.Context, in *GetCertAuthoritiesRequest, opts ...grpc.CallOption) (*GetCertAuthoritiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCertAuthoritiesResponse)
	err := c.cc.Invoke(ctx, TrustService_GetCertAuthorities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustServiceClient) DeleteCertAuthority(ctx context.Context, in *DeleteCertAuthorityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TrustService_DeleteCertAuthority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustServiceClient) UpsertCertAuthority(ctx context.Context, in *UpsertCertAuthorityRequest, opts ...grpc.CallOption) (*types.CertAuthorityV2, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.CertAuthorityV2)
	err := c.cc.Invoke(ctx, TrustService_UpsertCertAuthority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustServiceClient) RotateCertAuthority(ctx context.Context, in *RotateCertAuthorityRequest, opts ...grpc.CallOption) (*RotateCertAuthorityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RotateCertAuthorityResponse)
	err := c.cc.Invoke(ctx, TrustService_RotateCertAuthority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustServiceClient) RotateExternalCertAuthority(ctx context.Context, in *RotateExternalCertAuthorityRequest, opts ...grpc.CallOption) (*RotateExternalCertAuthorityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RotateExternalCertAuthorityResponse)
	err := c.cc.Invoke(ctx, TrustService_RotateExternalCertAuthority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustServiceClient) GenerateHostCert(ctx context.Context, in *GenerateHostCertRequest, opts ...grpc.CallOption) (*GenerateHostCertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateHostCertResponse)
	err := c.cc.Invoke(ctx, TrustService_GenerateHostCert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustServiceClient) UpsertTrustedCluster(ctx context.Context, in *UpsertTrustedClusterRequest, opts ...grpc.CallOption) (*types.TrustedClusterV2, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.TrustedClusterV2)
	err := c.cc.Invoke(ctx, TrustService_UpsertTrustedCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustServiceClient) CreateTrustedCluster(ctx context.Context, in *CreateTrustedClusterRequest, opts ...grpc.CallOption) (*types.TrustedClusterV2, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.TrustedClusterV2)
	err := c.cc.Invoke(ctx, TrustService_CreateTrustedCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustServiceClient) UpdateTrustedCluster(ctx context.Context, in *UpdateTrustedClusterRequest, opts ...grpc.CallOption) (*types.TrustedClusterV2, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.TrustedClusterV2)
	err := c.cc.Invoke(ctx, TrustService_UpdateTrustedCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrustServiceServer is the server API for TrustService service.
// All implementations must embed UnimplementedTrustServiceServer
// for forward compatibility.
//
// TrustService provides methods to manage certificate authorities.
type TrustServiceServer interface {
	// GetCertAuthority returns a cert authority by type and domain.
	GetCertAuthority(context.Context, *GetCertAuthorityRequest) (*types.CertAuthorityV2, error)
	// GetCertAuthorities returns all cert authorities with the specified type.
	GetCertAuthorities(context.Context, *GetCertAuthoritiesRequest) (*GetCertAuthoritiesResponse, error)
	// DeleteCertAuthority deletes the matching cert authority.
	DeleteCertAuthority(context.Context, *DeleteCertAuthorityRequest) (*emptypb.Empty, error)
	// UpsertCertAuthority creates or updates the provided cert authority.
	UpsertCertAuthority(context.Context, *UpsertCertAuthorityRequest) (*types.CertAuthorityV2, error)
	// RotateCertAuthority is a request to start rotation of the certificate authority.
	RotateCertAuthority(context.Context, *RotateCertAuthorityRequest) (*RotateCertAuthorityResponse, error)
	// RotateExternalCertAuthority rotates an external cert authority.
	RotateExternalCertAuthority(context.Context, *RotateExternalCertAuthorityRequest) (*RotateExternalCertAuthorityResponse, error)
	// GenerateHostCert takes a public key in the OpenSSH `authorized_keys` format and returns
	// a SSH certificate signed by the Host CA.
	GenerateHostCert(context.Context, *GenerateHostCertRequest) (*GenerateHostCertResponse, error)
	// UpsertTrustedCluster upserts a Trusted Cluster in a backend.
	UpsertTrustedCluster(context.Context, *UpsertTrustedClusterRequest) (*types.TrustedClusterV2, error)
	// CreateTrustedCluster creates a Trusted Cluster in a backend.
	CreateTrustedCluster(context.Context, *CreateTrustedClusterRequest) (*types.TrustedClusterV2, error)
	// UpdateTrustedCluster updates a Trusted Cluster in a backend.
	UpdateTrustedCluster(context.Context, *UpdateTrustedClusterRequest) (*types.TrustedClusterV2, error)
	mustEmbedUnimplementedTrustServiceServer()
}

// UnimplementedTrustServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrustServiceServer struct{}

func (UnimplementedTrustServiceServer) GetCertAuthority(context.Context, *GetCertAuthorityRequest) (*types.CertAuthorityV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertAuthority not implemented")
}
func (UnimplementedTrustServiceServer) GetCertAuthorities(context.Context, *GetCertAuthoritiesRequest) (*GetCertAuthoritiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertAuthorities not implemented")
}
func (UnimplementedTrustServiceServer) DeleteCertAuthority(context.Context, *DeleteCertAuthorityRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCertAuthority not implemented")
}
func (UnimplementedTrustServiceServer) UpsertCertAuthority(context.Context, *UpsertCertAuthorityRequest) (*types.CertAuthorityV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertCertAuthority not implemented")
}
func (UnimplementedTrustServiceServer) RotateCertAuthority(context.Context, *RotateCertAuthorityRequest) (*RotateCertAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateCertAuthority not implemented")
}
func (UnimplementedTrustServiceServer) RotateExternalCertAuthority(context.Context, *RotateExternalCertAuthorityRequest) (*RotateExternalCertAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateExternalCertAuthority not implemented")
}
func (UnimplementedTrustServiceServer) GenerateHostCert(context.Context, *GenerateHostCertRequest) (*GenerateHostCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateHostCert not implemented")
}
func (UnimplementedTrustServiceServer) UpsertTrustedCluster(context.Context, *UpsertTrustedClusterRequest) (*types.TrustedClusterV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertTrustedCluster not implemented")
}
func (UnimplementedTrustServiceServer) CreateTrustedCluster(context.Context, *CreateTrustedClusterRequest) (*types.TrustedClusterV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrustedCluster not implemented")
}
func (UnimplementedTrustServiceServer) UpdateTrustedCluster(context.Context, *UpdateTrustedClusterRequest) (*types.TrustedClusterV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrustedCluster not implemented")
}
func (UnimplementedTrustServiceServer) mustEmbedUnimplementedTrustServiceServer() {}
func (UnimplementedTrustServiceServer) testEmbeddedByValue()                      {}

// UnsafeTrustServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrustServiceServer will
// result in compilation errors.
type UnsafeTrustServiceServer interface {
	mustEmbedUnimplementedTrustServiceServer()
}

func RegisterTrustServiceServer(s grpc.ServiceRegistrar, srv TrustServiceServer) {
	// If the following call pancis, it indicates UnimplementedTrustServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TrustService_ServiceDesc, srv)
}

func _TrustService_GetCertAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).GetCertAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_GetCertAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).GetCertAuthority(ctx, req.(*GetCertAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustService_GetCertAuthorities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertAuthoritiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).GetCertAuthorities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_GetCertAuthorities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).GetCertAuthorities(ctx, req.(*GetCertAuthoritiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustService_DeleteCertAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCertAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).DeleteCertAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_DeleteCertAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).DeleteCertAuthority(ctx, req.(*DeleteCertAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustService_UpsertCertAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertCertAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).UpsertCertAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_UpsertCertAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).UpsertCertAuthority(ctx, req.(*UpsertCertAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustService_RotateCertAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateCertAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).RotateCertAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_RotateCertAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).RotateCertAuthority(ctx, req.(*RotateCertAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustService_RotateExternalCertAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateExternalCertAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).RotateExternalCertAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_RotateExternalCertAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).RotateExternalCertAuthority(ctx, req.(*RotateExternalCertAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustService_GenerateHostCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateHostCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).GenerateHostCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_GenerateHostCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).GenerateHostCert(ctx, req.(*GenerateHostCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustService_UpsertTrustedCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertTrustedClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).UpsertTrustedCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_UpsertTrustedCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).UpsertTrustedCluster(ctx, req.(*UpsertTrustedClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustService_CreateTrustedCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrustedClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).CreateTrustedCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_CreateTrustedCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).CreateTrustedCluster(ctx, req.(*CreateTrustedClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustService_UpdateTrustedCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrustedClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustServiceServer).UpdateTrustedCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustService_UpdateTrustedCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustServiceServer).UpdateTrustedCluster(ctx, req.(*UpdateTrustedClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrustService_ServiceDesc is the grpc.ServiceDesc for TrustService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrustService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.trust.v1.TrustService",
	HandlerType: (*TrustServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCertAuthority",
			Handler:    _TrustService_GetCertAuthority_Handler,
		},
		{
			MethodName: "GetCertAuthorities",
			Handler:    _TrustService_GetCertAuthorities_Handler,
		},
		{
			MethodName: "DeleteCertAuthority",
			Handler:    _TrustService_DeleteCertAuthority_Handler,
		},
		{
			MethodName: "UpsertCertAuthority",
			Handler:    _TrustService_UpsertCertAuthority_Handler,
		},
		{
			MethodName: "RotateCertAuthority",
			Handler:    _TrustService_RotateCertAuthority_Handler,
		},
		{
			MethodName: "RotateExternalCertAuthority",
			Handler:    _TrustService_RotateExternalCertAuthority_Handler,
		},
		{
			MethodName: "GenerateHostCert",
			Handler:    _TrustService_GenerateHostCert_Handler,
		},
		{
			MethodName: "UpsertTrustedCluster",
			Handler:    _TrustService_UpsertTrustedCluster_Handler,
		},
		{
			MethodName: "CreateTrustedCluster",
			Handler:    _TrustService_CreateTrustedCluster_Handler,
		},
		{
			MethodName: "UpdateTrustedCluster",
			Handler:    _TrustService_UpdateTrustedCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/trust/v1/trust_service.proto",
}
