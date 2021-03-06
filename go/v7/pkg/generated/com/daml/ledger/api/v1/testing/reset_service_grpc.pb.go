// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package testing

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ResetServiceClient is the client API for ResetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResetServiceClient interface {
	// Resets the ledger state. Note that loaded DARs won't be removed -- this only rolls back the
	// ledger to genesis.
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type resetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResetServiceClient(cc grpc.ClientConnInterface) ResetServiceClient {
	return &resetServiceClient{cc}
}

var resetServiceResetStreamDesc = &grpc.StreamDesc{
	StreamName: "Reset",
}

func (c *resetServiceClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.testing.ResetService/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResetServiceService is the service API for ResetService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterResetServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type ResetServiceService struct {
	// Resets the ledger state. Note that loaded DARs won't be removed -- this only rolls back the
	// ledger to genesis.
	Reset func(context.Context, *ResetRequest) (*empty.Empty, error)
}

func (s *ResetServiceService) reset(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.Reset == nil {
		return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
	}
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.daml.ledger.api.v1.testing.ResetService/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterResetServiceService registers a service implementation with a gRPC server.
func RegisterResetServiceService(s grpc.ServiceRegistrar, srv *ResetServiceService) {
	sd := grpc.ServiceDesc{
		ServiceName: "com.daml.ledger.api.v1.testing.ResetService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Reset",
				Handler:    srv.reset,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "com/daml/ledger/api/v1/testing/reset_service.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewResetServiceService creates a new ResetServiceService containing the
// implemented methods of the ResetService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewResetServiceService(s interface{}) *ResetServiceService {
	ns := &ResetServiceService{}
	if h, ok := s.(interface {
		Reset(context.Context, *ResetRequest) (*empty.Empty, error)
	}); ok {
		ns.Reset = h.Reset
	}
	return ns
}

// UnstableResetServiceService is the service API for ResetService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableResetServiceService interface {
	// Resets the ledger state. Note that loaded DARs won't be removed -- this only rolls back the
	// ledger to genesis.
	Reset(context.Context, *ResetRequest) (*empty.Empty, error)
}
