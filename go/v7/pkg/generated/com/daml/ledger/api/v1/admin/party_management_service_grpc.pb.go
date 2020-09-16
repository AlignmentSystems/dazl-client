// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package admin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PartyManagementServiceClient is the client API for PartyManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartyManagementServiceClient interface {
	// Return the identifier of the backing participant.
	// All horizontally scaled replicas should return the same id.
	// This method is expected to succeed provided the backing participant is
	// healthy, otherwise it responds with INTERNAL grpc error.
	// daml-on-sql: returns an identifier supplied on command line at launch time
	// daml-on-kv-ledger: as above
	// canton: returns globally unique identifier of the backing participant
	GetParticipantId(ctx context.Context, in *GetParticipantIdRequest, opts ...grpc.CallOption) (*GetParticipantIdResponse, error)
	// Get the party details of the given parties. Only known parties will be
	// returned in the list.
	// This request will always succeed.
	GetParties(ctx context.Context, in *GetPartiesRequest, opts ...grpc.CallOption) (*GetPartiesResponse, error)
	// List the parties known by the backing participant.
	// The list returned contains parties whose ledger access is facilitated by
	// backing participant and the ones maintained elsewhere.
	// This request will always succeed.
	ListKnownParties(ctx context.Context, in *ListKnownPartiesRequest, opts ...grpc.CallOption) (*ListKnownPartiesResponse, error)
	// Adds a new party to the set managed by the backing participant.
	// Caller specifies a party identifier suggestion, the actual identifier
	// allocated might be different and is implementation specific.
	// This call may:
	// - Succeed, in which case the actual allocated identifier is visible in
	//   the response.
	// - Respond with UNIMPLEMENTED if synchronous party allocation is not
	//   supported by the backing participant.
	// - Respond with INVALID_ARGUMENT if the provided hint and/or display name
	//   is invalid on the given ledger (see below).
	// daml-on-sql: suggestion's uniqueness is checked and call rejected if the
	// identifier is already present
	// daml-on-kv-ledger: suggestion's uniqueness is checked by the validators in
	// the consensus layer and call rejected if the identifier is already present.
	// canton: completely different globally unique identifier is allocated.
	// Behind the scenes calls to an internal protocol are made. As that protocol
	// is richer than the the surface protocol, the arguments take implicit values
	AllocateParty(ctx context.Context, in *AllocatePartyRequest, opts ...grpc.CallOption) (*AllocatePartyResponse, error)
}

type partyManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartyManagementServiceClient(cc grpc.ClientConnInterface) PartyManagementServiceClient {
	return &partyManagementServiceClient{cc}
}

var partyManagementServiceGetParticipantIdStreamDesc = &grpc.StreamDesc{
	StreamName: "GetParticipantId",
}

func (c *partyManagementServiceClient) GetParticipantId(ctx context.Context, in *GetParticipantIdRequest, opts ...grpc.CallOption) (*GetParticipantIdResponse, error) {
	out := new(GetParticipantIdResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.PartyManagementService/GetParticipantId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var partyManagementServiceGetPartiesStreamDesc = &grpc.StreamDesc{
	StreamName: "GetParties",
}

func (c *partyManagementServiceClient) GetParties(ctx context.Context, in *GetPartiesRequest, opts ...grpc.CallOption) (*GetPartiesResponse, error) {
	out := new(GetPartiesResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.PartyManagementService/GetParties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var partyManagementServiceListKnownPartiesStreamDesc = &grpc.StreamDesc{
	StreamName: "ListKnownParties",
}

func (c *partyManagementServiceClient) ListKnownParties(ctx context.Context, in *ListKnownPartiesRequest, opts ...grpc.CallOption) (*ListKnownPartiesResponse, error) {
	out := new(ListKnownPartiesResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.PartyManagementService/ListKnownParties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var partyManagementServiceAllocatePartyStreamDesc = &grpc.StreamDesc{
	StreamName: "AllocateParty",
}

func (c *partyManagementServiceClient) AllocateParty(ctx context.Context, in *AllocatePartyRequest, opts ...grpc.CallOption) (*AllocatePartyResponse, error) {
	out := new(AllocatePartyResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.PartyManagementService/AllocateParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartyManagementServiceService is the service API for PartyManagementService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterPartyManagementServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type PartyManagementServiceService struct {
	// Return the identifier of the backing participant.
	// All horizontally scaled replicas should return the same id.
	// This method is expected to succeed provided the backing participant is
	// healthy, otherwise it responds with INTERNAL grpc error.
	// daml-on-sql: returns an identifier supplied on command line at launch time
	// daml-on-kv-ledger: as above
	// canton: returns globally unique identifier of the backing participant
	GetParticipantId func(context.Context, *GetParticipantIdRequest) (*GetParticipantIdResponse, error)
	// Get the party details of the given parties. Only known parties will be
	// returned in the list.
	// This request will always succeed.
	GetParties func(context.Context, *GetPartiesRequest) (*GetPartiesResponse, error)
	// List the parties known by the backing participant.
	// The list returned contains parties whose ledger access is facilitated by
	// backing participant and the ones maintained elsewhere.
	// This request will always succeed.
	ListKnownParties func(context.Context, *ListKnownPartiesRequest) (*ListKnownPartiesResponse, error)
	// Adds a new party to the set managed by the backing participant.
	// Caller specifies a party identifier suggestion, the actual identifier
	// allocated might be different and is implementation specific.
	// This call may:
	// - Succeed, in which case the actual allocated identifier is visible in
	//   the response.
	// - Respond with UNIMPLEMENTED if synchronous party allocation is not
	//   supported by the backing participant.
	// - Respond with INVALID_ARGUMENT if the provided hint and/or display name
	//   is invalid on the given ledger (see below).
	// daml-on-sql: suggestion's uniqueness is checked and call rejected if the
	// identifier is already present
	// daml-on-kv-ledger: suggestion's uniqueness is checked by the validators in
	// the consensus layer and call rejected if the identifier is already present.
	// canton: completely different globally unique identifier is allocated.
	// Behind the scenes calls to an internal protocol are made. As that protocol
	// is richer than the the surface protocol, the arguments take implicit values
	AllocateParty func(context.Context, *AllocatePartyRequest) (*AllocatePartyResponse, error)
}

func (s *PartyManagementServiceService) getParticipantId(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.GetParticipantId == nil {
		return nil, status.Errorf(codes.Unimplemented, "method GetParticipantId not implemented")
	}
	in := new(GetParticipantIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetParticipantId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.daml.ledger.api.v1.admin.PartyManagementService/GetParticipantId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetParticipantId(ctx, req.(*GetParticipantIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *PartyManagementServiceService) getParties(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.GetParties == nil {
		return nil, status.Errorf(codes.Unimplemented, "method GetParties not implemented")
	}
	in := new(GetPartiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetParties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.daml.ledger.api.v1.admin.PartyManagementService/GetParties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetParties(ctx, req.(*GetPartiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *PartyManagementServiceService) listKnownParties(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.ListKnownParties == nil {
		return nil, status.Errorf(codes.Unimplemented, "method ListKnownParties not implemented")
	}
	in := new(ListKnownPartiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ListKnownParties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.daml.ledger.api.v1.admin.PartyManagementService/ListKnownParties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListKnownParties(ctx, req.(*ListKnownPartiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *PartyManagementServiceService) allocateParty(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.AllocateParty == nil {
		return nil, status.Errorf(codes.Unimplemented, "method AllocateParty not implemented")
	}
	in := new(AllocatePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.AllocateParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.daml.ledger.api.v1.admin.PartyManagementService/AllocateParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.AllocateParty(ctx, req.(*AllocatePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterPartyManagementServiceService registers a service implementation with a gRPC server.
func RegisterPartyManagementServiceService(s grpc.ServiceRegistrar, srv *PartyManagementServiceService) {
	sd := grpc.ServiceDesc{
		ServiceName: "com.daml.ledger.api.v1.admin.PartyManagementService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "GetParticipantId",
				Handler:    srv.getParticipantId,
			},
			{
				MethodName: "GetParties",
				Handler:    srv.getParties,
			},
			{
				MethodName: "ListKnownParties",
				Handler:    srv.listKnownParties,
			},
			{
				MethodName: "AllocateParty",
				Handler:    srv.allocateParty,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "com/daml/ledger/api/v1/admin/party_management_service.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewPartyManagementServiceService creates a new PartyManagementServiceService containing the
// implemented methods of the PartyManagementService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewPartyManagementServiceService(s interface{}) *PartyManagementServiceService {
	ns := &PartyManagementServiceService{}
	if h, ok := s.(interface {
		GetParticipantId(context.Context, *GetParticipantIdRequest) (*GetParticipantIdResponse, error)
	}); ok {
		ns.GetParticipantId = h.GetParticipantId
	}
	if h, ok := s.(interface {
		GetParties(context.Context, *GetPartiesRequest) (*GetPartiesResponse, error)
	}); ok {
		ns.GetParties = h.GetParties
	}
	if h, ok := s.(interface {
		ListKnownParties(context.Context, *ListKnownPartiesRequest) (*ListKnownPartiesResponse, error)
	}); ok {
		ns.ListKnownParties = h.ListKnownParties
	}
	if h, ok := s.(interface {
		AllocateParty(context.Context, *AllocatePartyRequest) (*AllocatePartyResponse, error)
	}); ok {
		ns.AllocateParty = h.AllocateParty
	}
	return ns
}

// UnstablePartyManagementServiceService is the service API for PartyManagementService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstablePartyManagementServiceService interface {
	// Return the identifier of the backing participant.
	// All horizontally scaled replicas should return the same id.
	// This method is expected to succeed provided the backing participant is
	// healthy, otherwise it responds with INTERNAL grpc error.
	// daml-on-sql: returns an identifier supplied on command line at launch time
	// daml-on-kv-ledger: as above
	// canton: returns globally unique identifier of the backing participant
	GetParticipantId(context.Context, *GetParticipantIdRequest) (*GetParticipantIdResponse, error)
	// Get the party details of the given parties. Only known parties will be
	// returned in the list.
	// This request will always succeed.
	GetParties(context.Context, *GetPartiesRequest) (*GetPartiesResponse, error)
	// List the parties known by the backing participant.
	// The list returned contains parties whose ledger access is facilitated by
	// backing participant and the ones maintained elsewhere.
	// This request will always succeed.
	ListKnownParties(context.Context, *ListKnownPartiesRequest) (*ListKnownPartiesResponse, error)
	// Adds a new party to the set managed by the backing participant.
	// Caller specifies a party identifier suggestion, the actual identifier
	// allocated might be different and is implementation specific.
	// This call may:
	// - Succeed, in which case the actual allocated identifier is visible in
	//   the response.
	// - Respond with UNIMPLEMENTED if synchronous party allocation is not
	//   supported by the backing participant.
	// - Respond with INVALID_ARGUMENT if the provided hint and/or display name
	//   is invalid on the given ledger (see below).
	// daml-on-sql: suggestion's uniqueness is checked and call rejected if the
	// identifier is already present
	// daml-on-kv-ledger: suggestion's uniqueness is checked by the validators in
	// the consensus layer and call rejected if the identifier is already present.
	// canton: completely different globally unique identifier is allocated.
	// Behind the scenes calls to an internal protocol are made. As that protocol
	// is richer than the the surface protocol, the arguments take implicit values
	AllocateParty(context.Context, *AllocatePartyRequest) (*AllocatePartyResponse, error)
}