// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
// source: rpc/io.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/rand"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import google_protobuf "google.golang.org/protobuf/types/known/emptypb"
import livekit4 "github.com/livekit/protocol/livekit"
import livekit5 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_6

// =======================
// IOInfo Client Interface
// =======================

type IOInfoClient interface {
	// egress
	CreateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	UpdateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit4.EgressInfo, error)

	ListEgress(ctx context.Context, req *livekit4.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit4.ListEgressResponse, error)

	UpdateMetrics(ctx context.Context, req *UpdateMetricsRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// ingress
	CreateIngress(ctx context.Context, req *livekit5.IngressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error)

	UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// sip
	GetSIPTrunkAuthentication(ctx context.Context, req *GetSIPTrunkAuthenticationRequest, opts ...psrpc.RequestOption) (*GetSIPTrunkAuthenticationResponse, error)

	EvaluateSIPDispatchRules(ctx context.Context, req *EvaluateSIPDispatchRulesRequest, opts ...psrpc.RequestOption) (*EvaluateSIPDispatchRulesResponse, error)

	UpdateSIPCallState(ctx context.Context, req *UpdateSIPCallStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ===========================
// IOInfo ServerImpl Interface
// ===========================

type IOInfoServerImpl interface {
	// egress
	CreateEgress(context.Context, *livekit4.EgressInfo) (*google_protobuf.Empty, error)

	UpdateEgress(context.Context, *livekit4.EgressInfo) (*google_protobuf.Empty, error)

	GetEgress(context.Context, *GetEgressRequest) (*livekit4.EgressInfo, error)

	ListEgress(context.Context, *livekit4.ListEgressRequest) (*livekit4.ListEgressResponse, error)

	UpdateMetrics(context.Context, *UpdateMetricsRequest) (*google_protobuf.Empty, error)

	// ingress
	CreateIngress(context.Context, *livekit5.IngressInfo) (*google_protobuf.Empty, error)

	GetIngressInfo(context.Context, *GetIngressInfoRequest) (*GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *UpdateIngressStateRequest) (*google_protobuf.Empty, error)

	// sip
	GetSIPTrunkAuthentication(context.Context, *GetSIPTrunkAuthenticationRequest) (*GetSIPTrunkAuthenticationResponse, error)

	EvaluateSIPDispatchRules(context.Context, *EvaluateSIPDispatchRulesRequest) (*EvaluateSIPDispatchRulesResponse, error)

	UpdateSIPCallState(context.Context, *UpdateSIPCallStateRequest) (*google_protobuf.Empty, error)
}

// =======================
// IOInfo Server Interface
// =======================

type IOInfoServer interface {

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// IOInfo Client
// =============

type iOInfoClient struct {
	client *client.RPCClient
}

// NewIOInfoClient creates a psrpc client that implements the IOInfoClient interface.
func NewIOInfoClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IOInfoClient, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CreateEgress", false, false, true, true)
	sd.RegisterMethod("UpdateEgress", false, false, true, true)
	sd.RegisterMethod("GetEgress", false, false, true, true)
	sd.RegisterMethod("ListEgress", false, false, true, true)
	sd.RegisterMethod("UpdateMetrics", false, false, true, true)
	sd.RegisterMethod("CreateIngress", false, false, true, true)
	sd.RegisterMethod("GetIngressInfo", false, false, true, true)
	sd.RegisterMethod("UpdateIngressState", false, false, true, true)
	sd.RegisterMethod("GetSIPTrunkAuthentication", false, false, true, true)
	sd.RegisterMethod("EvaluateSIPDispatchRules", false, false, true, true)
	sd.RegisterMethod("UpdateSIPCallState", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &iOInfoClient{
		client: rpcClient,
	}, nil
}

func (c *iOInfoClient) CreateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "CreateEgress", nil, req, opts...)
}

func (c *iOInfoClient) UpdateEgress(ctx context.Context, req *livekit4.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateEgress", nil, req, opts...)
}

func (c *iOInfoClient) GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit4.EgressInfo, error) {
	return client.RequestSingle[*livekit4.EgressInfo](ctx, c.client, "GetEgress", nil, req, opts...)
}

func (c *iOInfoClient) ListEgress(ctx context.Context, req *livekit4.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit4.ListEgressResponse, error) {
	return client.RequestSingle[*livekit4.ListEgressResponse](ctx, c.client, "ListEgress", nil, req, opts...)
}

func (c *iOInfoClient) UpdateMetrics(ctx context.Context, req *UpdateMetricsRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateMetrics", nil, req, opts...)
}

func (c *iOInfoClient) CreateIngress(ctx context.Context, req *livekit5.IngressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "CreateIngress", nil, req, opts...)
}

func (c *iOInfoClient) GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error) {
	return client.RequestSingle[*GetIngressInfoResponse](ctx, c.client, "GetIngressInfo", nil, req, opts...)
}

func (c *iOInfoClient) UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateIngressState", nil, req, opts...)
}

func (c *iOInfoClient) GetSIPTrunkAuthentication(ctx context.Context, req *GetSIPTrunkAuthenticationRequest, opts ...psrpc.RequestOption) (*GetSIPTrunkAuthenticationResponse, error) {
	return client.RequestSingle[*GetSIPTrunkAuthenticationResponse](ctx, c.client, "GetSIPTrunkAuthentication", nil, req, opts...)
}

func (c *iOInfoClient) EvaluateSIPDispatchRules(ctx context.Context, req *EvaluateSIPDispatchRulesRequest, opts ...psrpc.RequestOption) (*EvaluateSIPDispatchRulesResponse, error) {
	return client.RequestSingle[*EvaluateSIPDispatchRulesResponse](ctx, c.client, "EvaluateSIPDispatchRules", nil, req, opts...)
}

func (c *iOInfoClient) UpdateSIPCallState(ctx context.Context, req *UpdateSIPCallStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateSIPCallState", nil, req, opts...)
}

func (s *iOInfoClient) Close() {
	s.client.Close()
}

// =============
// IOInfo Server
// =============

type iOInfoServer struct {
	svc IOInfoServerImpl
	rpc *server.RPCServer
}

// NewIOInfoServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIOInfoServer(svc IOInfoServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IOInfoServer, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CreateEgress", false, false, true, true)
	var err error
	err = server.RegisterHandler(s, "CreateEgress", nil, svc.CreateEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateEgress", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateEgress", nil, svc.UpdateEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetEgress", false, false, true, true)
	err = server.RegisterHandler(s, "GetEgress", nil, svc.GetEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("ListEgress", false, false, true, true)
	err = server.RegisterHandler(s, "ListEgress", nil, svc.ListEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateMetrics", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateMetrics", nil, svc.UpdateMetrics, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("CreateIngress", false, false, true, true)
	err = server.RegisterHandler(s, "CreateIngress", nil, svc.CreateIngress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetIngressInfo", false, false, true, true)
	err = server.RegisterHandler(s, "GetIngressInfo", nil, svc.GetIngressInfo, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateIngressState", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateIngressState", nil, svc.UpdateIngressState, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetSIPTrunkAuthentication", false, false, true, true)
	err = server.RegisterHandler(s, "GetSIPTrunkAuthentication", nil, svc.GetSIPTrunkAuthentication, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("EvaluateSIPDispatchRules", false, false, true, true)
	err = server.RegisterHandler(s, "EvaluateSIPDispatchRules", nil, svc.EvaluateSIPDispatchRules, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateSIPCallState", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateSIPCallState", nil, svc.UpdateSIPCallState, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &iOInfoServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *iOInfoServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *iOInfoServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor4 = []byte{
	// 1538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x58, 0xdd, 0x52, 0x1b, 0x47,
	0x16, 0x5e, 0x49, 0x20, 0xa3, 0x23, 0x90, 0x44, 0x23, 0xb1, 0x83, 0xa8, 0xb5, 0xb1, 0xbc, 0xde,
	0x65, 0x77, 0xab, 0xc4, 0x9a, 0x5c, 0x24, 0xb1, 0x2b, 0x2e, 0x83, 0x2c, 0xf0, 0x24, 0x18, 0x94,
	0x01, 0x2e, 0x92, 0x9b, 0x49, 0x33, 0xd3, 0x88, 0x31, 0xa3, 0xe9, 0x49, 0x77, 0x0f, 0x98, 0x27,
	0x48, 0x1e, 0x23, 0x8f, 0x90, 0x27, 0x48, 0x55, 0x5e, 0x2a, 0xd7, 0xa9, 0xfe, 0x19, 0x69, 0x84,
	0x24, 0x03, 0xb9, 0xd3, 0x7c, 0xe7, 0xff, 0xa7, 0xbb, 0xcf, 0x11, 0x2c, 0xb2, 0xd8, 0xdb, 0x0a,
	0x68, 0x3b, 0x66, 0x54, 0x50, 0x54, 0x60, 0xb1, 0xd7, 0xac, 0x87, 0xc1, 0x15, 0xb9, 0x0c, 0x84,
	0x4b, 0xfa, 0x8c, 0x70, 0xae, 0x49, 0xcd, 0x46, 0x8a, 0x06, 0x51, 0x16, 0x5e, 0x4e, 0x61, 0x1e,
	0xc4, 0x06, 0x42, 0x29, 0xc4, 0x28, 0x1d, 0x18, 0x6c, 0xbd, 0x4f, 0x69, 0x3f, 0x24, 0x5b, 0xea,
	0xeb, 0x2c, 0x39, 0xdf, 0x22, 0x83, 0x58, 0xdc, 0x18, 0xe2, 0xe3, 0xdb, 0x44, 0x3f, 0x61, 0x58,
	0x04, 0x34, 0xd2, 0xf4, 0xd6, 0x16, 0xd4, 0xf6, 0x89, 0xe8, 0x2a, 0xb3, 0x0e, 0xf9, 0x31, 0x21,
	0x5c, 0xa0, 0x75, 0x28, 0x69, 0xf7, 0xdc, 0xc0, 0xb7, 0x72, 0x1b, 0xb9, 0xcd, 0x92, 0xb3, 0xa0,
	0x01, 0xdb, 0x6f, 0xfd, 0x94, 0x83, 0xfa, 0x69, 0xec, 0x63, 0x41, 0xde, 0x13, 0xc1, 0x02, 0x6f,
	0x28, 0xf5, 0x6f, 0x98, 0x0b, 0xa2, 0x73, 0xaa, 0x04, 0xca, 0xdb, 0x2b, 0x6d, 0xe3, 0x69, 0x5b,
	0xeb, 0xb6, 0xa3, 0x73, 0xea, 0x28, 0x06, 0xd4, 0x82, 0x25, 0x7c, 0xd5, 0x77, 0xbd, 0x38, 0x71,
	0x13, 0x8e, 0xfb, 0xc4, 0x2a, 0x6c, 0xe4, 0x36, 0xf3, 0x4e, 0x19, 0x5f, 0xf5, 0x3b, 0x71, 0x72,
	0x2a, 0x21, 0xc9, 0x33, 0xc0, 0x1f, 0x33, 0x3c, 0x73, 0x9a, 0x67, 0x80, 0x3f, 0xa6, 0x3c, 0xad,
	0x53, 0x68, 0xec, 0x13, 0x61, 0x47, 0x23, 0xfd, 0xc6, 0x93, 0x7f, 0x00, 0x98, 0x44, 0x8e, 0x02,
	0x28, 0x19, 0xc4, 0xf6, 0x25, 0x99, 0x0b, 0x46, 0xf0, 0xc0, 0xbd, 0x24, 0x37, 0x56, 0x5e, 0x93,
	0x35, 0xf2, 0x0d, 0xb9, 0x69, 0xfd, 0x9c, 0x87, 0xd5, 0xdb, 0x7a, 0x79, 0x4c, 0x23, 0x4e, 0xd0,
	0xe6, 0x58, 0x88, 0xf5, 0x61, 0x88, 0x59, 0x5e, 0x1d, 0x63, 0x1d, 0xe6, 0x05, 0xbd, 0x24, 0x91,
	0x51, 0xaf, 0x3f, 0x50, 0x03, 0x8a, 0xd7, 0xdc, 0x4d, 0x58, 0xa8, 0x42, 0x2e, 0x39, 0xf3, 0xd7,
	0xfc, 0x94, 0x85, 0xe8, 0x14, 0x2a, 0x21, 0xed, 0xf7, 0x83, 0xa8, 0xef, 0x9e, 0x07, 0x24, 0xf4,
	0xb9, 0x35, 0xb7, 0x51, 0xd8, 0x2c, 0x6f, 0xb7, 0xdb, 0x2c, 0xf6, 0xda, 0xd3, 0x7d, 0x69, 0x1f,
	0x68, 0x89, 0x3d, 0x25, 0xd0, 0x8d, 0x04, 0xbb, 0x71, 0x96, 0xc2, 0x2c, 0xd6, 0x7c, 0x03, 0x68,
	0x92, 0x09, 0xd5, 0xa0, 0x20, 0xc3, 0xd6, 0x59, 0x91, 0x3f, 0xa5, 0xaf, 0x57, 0x38, 0x4c, 0x48,
	0xea, 0xab, 0xfa, 0x78, 0x99, 0xff, 0x22, 0xd7, 0xea, 0xc3, 0x9a, 0x2e, 0xb5, 0x71, 0xe0, 0x58,
	0x60, 0x41, 0xee, 0x99, 0xe5, 0xff, 0xc1, 0x3c, 0x97, 0xec, 0x4a, 0x6b, 0x79, 0xbb, 0x71, 0x3b,
	0x59, 0x5a, 0x97, 0xe6, 0x69, 0xfd, 0x92, 0x83, 0x8d, 0x7d, 0x22, 0x8e, 0xed, 0xde, 0x09, 0x4b,
	0xa2, 0xcb, 0x9d, 0x44, 0x5c, 0x90, 0x48, 0x04, 0x9e, 0xea, 0xd4, 0xd4, 0xe0, 0x63, 0x28, 0xf3,
	0x20, 0x76, 0x3d, 0x1c, 0x86, 0xd2, 0x62, 0xd1, 0x14, 0x2e, 0x88, 0x3b, 0x38, 0x0c, 0x6d, 0x1f,
	0x21, 0x98, 0x3b, 0x67, 0x74, 0x60, 0xc2, 0x50, 0xbf, 0x51, 0x05, 0xf2, 0x82, 0x9a, 0x6c, 0xe7,
	0x05, 0x45, 0x4f, 0xa0, 0xcc, 0x99, 0xe7, 0x62, 0xdf, 0x97, 0x3e, 0xa8, 0xae, 0x2a, 0x39, 0xc0,
	0x99, 0xb7, 0xa3, 0x11, 0xf4, 0x77, 0x78, 0x24, 0xa8, 0x7b, 0x41, 0xb9, 0xb0, 0xe6, 0x15, 0xb1,
	0x28, 0xe8, 0x3b, 0xca, 0x45, 0xeb, 0xd7, 0x1c, 0x3c, 0xfd, 0x84, 0x8b, 0xa6, 0x43, 0x9a, 0xb0,
	0x90, 0x70, 0xc2, 0x22, 0x3c, 0x20, 0xe9, 0xc9, 0x49, 0xbf, 0x25, 0x2d, 0xc6, 0x9c, 0x5f, 0x53,
	0xe6, 0x1b, 0x1f, 0x87, 0xdf, 0xd2, 0x77, 0x9f, 0xd1, 0x58, 0x79, 0xba, 0xe0, 0xa8, 0xdf, 0x68,
	0x03, 0x16, 0x65, 0xbc, 0x42, 0x9a, 0x93, 0x01, 0xa7, 0xce, 0x06, 0xb1, 0xf2, 0x40, 0x77, 0x72,
	0xcc, 0xe8, 0x07, 0xe2, 0x09, 0x49, 0xd7, 0xfe, 0x96, 0x0c, 0x62, 0xfb, 0xad, 0x3f, 0x0a, 0xf0,
	0xa4, 0x2b, 0xab, 0x89, 0x05, 0x39, 0xb6, 0x7b, 0x6f, 0x03, 0x1e, 0x63, 0xe1, 0x5d, 0x38, 0x49,
	0x48, 0xf8, 0x8c, 0xa4, 0x2e, 0xdc, 0x4e, 0xea, 0xff, 0x01, 0x49, 0x7a, 0x8c, 0x99, 0x08, 0xbc,
	0x20, 0xc6, 0x91, 0x18, 0x56, 0x7b, 0x37, 0x6f, 0xe5, 0x9c, 0x1a, 0x0f, 0xe2, 0xde, 0x88, 0x68,
	0xfb, 0x13, 0x6e, 0xc3, 0x84, 0xdb, 0xcf, 0xa1, 0x22, 0xed, 0xc9, 0x7e, 0x8f, 0x92, 0xc1, 0x19,
	0x61, 0x26, 0x1d, 0x4b, 0x06, 0x3d, 0x54, 0x20, 0x7a, 0x06, 0x0a, 0x20, 0x7e, 0xca, 0xa5, 0xcb,
	0xb8, 0xa8, 0x41, 0xc3, 0x74, 0x67, 0x41, 0x6b, 0x50, 0x88, 0x83, 0xc8, 0x24, 0x47, 0xfe, 0x94,
	0xa7, 0x30, 0xa2, 0xae, 0x04, 0x8b, 0x2a, 0xdb, 0xf3, 0x11, 0xed, 0x05, 0x91, 0xd4, 0x64, 0xcc,
	0xa9, 0xea, 0x3f, 0xd2, 0x9a, 0x34, 0x24, 0x3b, 0x00, 0xf9, 0x50, 0x23, 0x1f, 0x05, 0xc3, 0x2e,
	0x16, 0x82, 0x05, 0x67, 0x89, 0x20, 0xdc, 0x2a, 0xa9, 0x83, 0xfa, 0xa5, 0x3a, 0xa8, 0x77, 0xa4,
	0xba, 0xdd, 0x95, 0xc2, 0x3b, 0x43, 0x59, 0x7d, 0x66, 0xab, 0x64, 0x1c, 0x6d, 0xee, 0x42, 0x7d,
	0x1a, 0xe3, 0x83, 0xce, 0xed, 0x6f, 0x8b, 0xb0, 0x31, 0xdb, 0x1b, 0xd3, 0xaa, 0xeb, 0x50, 0x92,
	0x8f, 0x88, 0x9b, 0xed, 0x55, 0x09, 0x1c, 0xca, 0x5e, 0x7d, 0x01, 0xf5, 0xf1, 0x92, 0xcb, 0x5e,
	0x17, 0xe9, 0x6d, 0xb9, 0x12, 0x67, 0x2b, 0xae, 0x49, 0xe8, 0x3f, 0x50, 0xcb, 0x8a, 0x28, 0xb5,
	0x3a, 0x89, 0xd5, 0x0c, 0x3e, 0x4d, 0xfb, 0x80, 0x08, 0xec, 0x63, 0x81, 0x4d, 0xf7, 0x65, 0xb5,
	0xbf, 0x37, 0x24, 0x74, 0x0d, 0xab, 0x59, 0x91, 0x4c, 0x09, 0xca, 0xaa, 0x04, 0x6f, 0xee, 0x28,
	0x81, 0xb9, 0x35, 0x33, 0xad, 0x7a, 0xbb, 0x12, 0x8d, 0x78, 0x1a, 0x0d, 0x3d, 0x83, 0x32, 0xd3,
	0x05, 0x54, 0x2d, 0xa3, 0x0e, 0xa8, 0xea, 0x7c, 0x30, 0xb0, 0xec, 0x9d, 0xe1, 0x75, 0x3f, 0x37,
	0xfd, 0xba, 0x9f, 0xcf, 0x5e, 0xf7, 0x6d, 0x28, 0x32, 0xc2, 0x93, 0x50, 0xa8, 0xfe, 0xab, 0x6c,
	0xaf, 0x2a, 0xd7, 0xb3, 0x2e, 0x2b, 0xaa, 0x63, 0xb8, 0x26, 0x0e, 0x54, 0x69, 0xe2, 0x40, 0x6d,
	0x41, 0x5d, 0x72, 0xf8, 0x46, 0xde, 0x65, 0x49, 0x48, 0x46, 0x47, 0x6f, 0x99, 0x07, 0x71, 0x36,
	0x1b, 0x13, 0x17, 0xc7, 0xe2, 0xad, 0x8b, 0x03, 0x1d, 0xc0, 0xa3, 0x0b, 0x82, 0x7d, 0xc2, 0xb8,
	0xb5, 0xa4, 0xb2, 0xbb, 0x7d, 0xbf, 0xec, 0xbe, 0xd3, 0x42, 0x3a, 0x9f, 0xa9, 0x0a, 0xc4, 0xa0,
	0x61, 0x7e, 0xba, 0x82, 0x66, 0x2b, 0x57, 0x51, 0xba, 0x5f, 0x3f, 0x48, 0xf7, 0x09, 0xbd, 0x5d,
	0xb7, 0x95, 0x8b, 0x49, 0x8a, 0xb4, 0x39, 0x32, 0x24, 0xcd, 0xa6, 0xf1, 0xa0, 0x87, 0xd8, 0x1c,
	0x29, 0x3c, 0xa1, 0x63, 0xb1, 0xad, 0xe0, 0x49, 0x0a, 0xda, 0x85, 0x6a, 0x10, 0x79, 0x61, 0xe2,
	0x93, 0xa1, 0xb5, 0x15, 0x55, 0xe0, 0xb5, 0xe1, 0xdb, 0x77, 0x6c, 0xf7, 0x34, 0xf7, 0x51, 0x2c,
	0x9f, 0x0d, 0xee, 0x54, 0x8c, 0x44, 0xaa, 0xe3, 0x35, 0xd4, 0x48, 0x84, 0xcf, 0xe4, 0x2d, 0x74,
	0x4e, 0xb0, 0x48, 0x18, 0xe1, 0x56, 0x75, 0xa3, 0xb0, 0x59, 0xc9, 0x0c, 0x54, 0xc7, 0x76, 0x6f,
	0x4f, 0xd3, 0x9c, 0xaa, 0x61, 0x36, 0xdf, 0xca, 0x07, 0x16, 0x44, 0x6a, 0x94, 0x10, 0xc1, 0x80,
	0xd0, 0x44, 0x58, 0x35, 0xf5, 0xfe, 0xae, 0xb5, 0xf5, 0x20, 0xd8, 0x4e, 0x07, 0xc1, 0xf6, 0x5b,
	0x33, 0x08, 0x3a, 0x15, 0x23, 0x71, 0xa2, 0x05, 0x50, 0x17, 0x96, 0xd5, 0xec, 0x25, 0x9f, 0x84,
	0x74, 0x5a, 0xb4, 0x96, 0xef, 0xd2, 0x52, 0x95, 0xa3, 0x19, 0x0e, 0xc3, 0x14, 0x90, 0xf7, 0xa9,
	0xba, 0x5f, 0x62, 0x46, 0x38, 0x11, 0x56, 0x5d, 0x77, 0xad, 0x84, 0x7a, 0x0a, 0x41, 0xaf, 0x0c,
	0x83, 0x47, 0xa3, 0xf3, 0xa0, 0x6f, 0x35, 0x94, 0x85, 0xe6, 0x30, 0x4c, 0x87, 0xd2, 0x41, 0x47,
	0x91, 0x52, 0x13, 0x4a, 0x58, 0x43, 0x68, 0x0f, 0x6a, 0x03, 0xe2, 0x07, 0xd8, 0x25, 0x91, 0xc7,
	0x6e, 0x54, 0x36, 0xad, 0x55, 0x95, 0xed, 0xf5, 0x6c, 0xa2, 0xde, 0x4b, 0x9e, 0xee, 0x90, 0xc5,
	0xa9, 0x0e, 0xc6, 0x81, 0xe6, 0x3b, 0x68, 0xce, 0xbe, 0x13, 0x1e, 0x72, 0xe9, 0x36, 0x5f, 0xc2,
	0x62, 0xb6, 0x47, 0x1e, 0x24, 0xbb, 0x07, 0xd6, 0xac, 0xfe, 0x7e, 0xa8, 0x9e, 0x59, 0x3d, 0xfb,
	0xa0, 0x07, 0xe4, 0x30, 0x1d, 0xfc, 0x8e, 0xed, 0x9e, 0x2c, 0xea, 0xd8, 0xe0, 0xf7, 0x02, 0x4a,
	0x7a, 0x5c, 0x98, 0x36, 0x0a, 0x1b, 0x01, 0x35, 0xaa, 0x2e, 0x78, 0xe6, 0xd7, 0x7f, 0x7f, 0x80,
	0xe5, 0x89, 0xfb, 0x0d, 0x59, 0x50, 0x3f, 0xe8, 0xee, 0xef, 0x74, 0xbe, 0x73, 0x77, 0x3a, 0x9d,
	0x6e, 0xef, 0xc4, 0x3d, 0x72, 0xdc, 0x9e, 0x7d, 0x58, 0xfb, 0x1b, 0x02, 0x28, 0x6a, 0xa8, 0x96,
	0x43, 0x55, 0x28, 0x3b, 0xdd, 0x6f, 0x4f, 0xbb, 0xc7, 0x27, 0x8a, 0x98, 0x97, 0x44, 0xa7, 0xfb,
	0x75, 0xb7, 0x73, 0x52, 0x2b, 0xa0, 0x05, 0x98, 0x7b, 0xeb, 0x1c, 0xf5, 0x6a, 0x73, 0xdb, 0xbf,
	0x17, 0xa1, 0x68, 0x1f, 0x49, 0x63, 0xe8, 0x15, 0x2c, 0x76, 0x18, 0xc1, 0x82, 0xe8, 0xcd, 0x03,
	0x4d, 0x5b, 0x45, 0x9a, 0xab, 0x13, 0x9d, 0xdc, 0x95, 0x5b, 0x93, 0x14, 0xd6, 0x91, 0xff, 0x15,
	0xe1, 0xcf, 0xa1, 0x34, 0x5c, 0xa6, 0x50, 0x23, 0x9d, 0xde, 0xc7, 0x96, 0xab, 0xe6, 0x34, 0x85,
	0xa8, 0x0b, 0x70, 0x10, 0xf0, 0x54, 0x72, 0x74, 0x06, 0x46, 0x60, 0x2a, 0xbe, 0x3e, 0x95, 0x66,
	0x9e, 0xf4, 0x5d, 0x58, 0x1a, 0x5b, 0xcd, 0xd0, 0x9a, 0xf2, 0x61, 0xda, 0xba, 0x36, 0x33, 0x86,
	0xaf, 0x60, 0x49, 0x67, 0xcf, 0xcc, 0xe9, 0x68, 0xea, 0x9a, 0x33, 0x53, 0xdc, 0x86, 0xca, 0xf8,
	0xc2, 0x82, 0x9a, 0x53, 0xb7, 0x98, 0x34, 0x9a, 0xd9, 0x1b, 0x0e, 0x3a, 0x00, 0x34, 0xb9, 0x7d,
	0xa0, 0xc7, 0x99, 0x90, 0xa6, 0xac, 0x25, 0x33, 0x1d, 0xfb, 0x00, 0x6b, 0x33, 0xc7, 0x77, 0xf4,
	0x3c, 0xf5, 0xe3, 0x93, 0x1b, 0x48, 0xf3, 0x5f, 0x77, 0xb1, 0x19, 0xcf, 0xfb, 0x60, 0xcd, 0x7a,
	0x5b, 0xd0, 0x3f, 0xef, 0x33, 0x2b, 0x36, 0x9f, 0xdf, 0xeb, 0x81, 0x1a, 0xa5, 0x28, 0x7b, 0x4e,
	0xc7, 0x52, 0x34, 0xe5, 0x00, 0xcf, 0x4a, 0xd1, 0xee, 0xd3, 0xef, 0x9f, 0xf4, 0x03, 0x71, 0x91,
	0x9c, 0xb5, 0x3d, 0x3a, 0xd8, 0x32, 0x55, 0xd7, 0xff, 0x1c, 0x78, 0x34, 0xdc, 0x62, 0xb1, 0x77,
	0x56, 0x54, 0x5f, 0x9f, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x0a, 0x9a, 0x81, 0xdb, 0x10,
	0x00, 0x00,
}
