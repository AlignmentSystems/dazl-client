// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LedgerConfigurationServiceClient is the client API for LedgerConfigurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LedgerConfigurationServiceClient interface {
	// Returns the latest configuration as the first response, and publishes configuration updates in the same stream.
	GetLedgerConfiguration(ctx context.Context, in *GetLedgerConfigurationRequest, opts ...grpc.CallOption) (LedgerConfigurationService_GetLedgerConfigurationClient, error)
}

type ledgerConfigurationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLedgerConfigurationServiceClient(cc grpc.ClientConnInterface) LedgerConfigurationServiceClient {
	return &ledgerConfigurationServiceClient{cc}
}

var ledgerConfigurationServiceGetLedgerConfigurationStreamDesc = &grpc.StreamDesc{
	StreamName:    "GetLedgerConfiguration",
	ServerStreams: true,
}

func (c *ledgerConfigurationServiceClient) GetLedgerConfiguration(ctx context.Context, in *GetLedgerConfigurationRequest, opts ...grpc.CallOption) (LedgerConfigurationService_GetLedgerConfigurationClient, error) {
	stream, err := c.cc.NewStream(ctx, ledgerConfigurationServiceGetLedgerConfigurationStreamDesc, "/com.daml.ledger.api.v1.LedgerConfigurationService/GetLedgerConfiguration", opts...)
	if err != nil {
		return nil, err
	}
	x := &ledgerConfigurationServiceGetLedgerConfigurationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LedgerConfigurationService_GetLedgerConfigurationClient interface {
	Recv() (*GetLedgerConfigurationResponse, error)
	grpc.ClientStream
}

type ledgerConfigurationServiceGetLedgerConfigurationClient struct {
	grpc.ClientStream
}

func (x *ledgerConfigurationServiceGetLedgerConfigurationClient) Recv() (*GetLedgerConfigurationResponse, error) {
	m := new(GetLedgerConfigurationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LedgerConfigurationServiceService is the service API for LedgerConfigurationService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterLedgerConfigurationServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type LedgerConfigurationServiceService struct {
	// Returns the latest configuration as the first response, and publishes configuration updates in the same stream.
	GetLedgerConfiguration func(*GetLedgerConfigurationRequest, LedgerConfigurationService_GetLedgerConfigurationServer) error
}

func (s *LedgerConfigurationServiceService) getLedgerConfiguration(_ interface{}, stream grpc.ServerStream) error {
	if s.GetLedgerConfiguration == nil {
		return status.Errorf(codes.Unimplemented, "method GetLedgerConfiguration not implemented")
	}
	m := new(GetLedgerConfigurationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.GetLedgerConfiguration(m, &ledgerConfigurationServiceGetLedgerConfigurationServer{stream})
}

type LedgerConfigurationService_GetLedgerConfigurationServer interface {
	Send(*GetLedgerConfigurationResponse) error
	grpc.ServerStream
}

type ledgerConfigurationServiceGetLedgerConfigurationServer struct {
	grpc.ServerStream
}

func (x *ledgerConfigurationServiceGetLedgerConfigurationServer) Send(m *GetLedgerConfigurationResponse) error {
	return x.ServerStream.SendMsg(m)
}

// RegisterLedgerConfigurationServiceService registers a service implementation with a gRPC server.
func RegisterLedgerConfigurationServiceService(s grpc.ServiceRegistrar, srv *LedgerConfigurationServiceService) {
	sd := grpc.ServiceDesc{
		ServiceName: "com.daml.ledger.api.v1.LedgerConfigurationService",
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "GetLedgerConfiguration",
				Handler:       srv.getLedgerConfiguration,
				ServerStreams: true,
			},
		},
		Metadata: "com/daml/ledger/api/v1/ledger_configuration_service.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewLedgerConfigurationServiceService creates a new LedgerConfigurationServiceService containing the
// implemented methods of the LedgerConfigurationService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewLedgerConfigurationServiceService(s interface{}) *LedgerConfigurationServiceService {
	ns := &LedgerConfigurationServiceService{}
	if h, ok := s.(interface {
		GetLedgerConfiguration(*GetLedgerConfigurationRequest, LedgerConfigurationService_GetLedgerConfigurationServer) error
	}); ok {
		ns.GetLedgerConfiguration = h.GetLedgerConfiguration
	}
	return ns
}

// UnstableLedgerConfigurationServiceService is the service API for LedgerConfigurationService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableLedgerConfigurationServiceService interface {
	// Returns the latest configuration as the first response, and publishes configuration updates in the same stream.
	GetLedgerConfiguration(*GetLedgerConfigurationRequest, LedgerConfigurationService_GetLedgerConfigurationServer) error
}
