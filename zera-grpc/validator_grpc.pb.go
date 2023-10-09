// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: validator.proto

package zera_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ValidatorServiceClient is the client API for ValidatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValidatorServiceClient interface {
	Broadcast(ctx context.Context, in *Block, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SyncBlockchain(ctx context.Context, in *BlockSync, opts ...grpc.CallOption) (*BlockBatch, error)
	ValidatorRegistration(ctx context.Context, in *ValidatorRegistrationMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SyncValidatorList(ctx context.Context, in *ValidatorSyncRequest, opts ...grpc.CallOption) (*ValidatorSync, error)
	ValidatorCoin(ctx context.Context, in *CoinTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorMint(ctx context.Context, in *MintTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorItemMint(ctx context.Context, in *ItemizedMintTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorContract(ctx context.Context, in *InstrumentContract, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorGovernProposal(ctx context.Context, in *GovernanceProposal, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorGovernVote(ctx context.Context, in *GovernanceVote, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorSmartContract(ctx context.Context, in *SmartContractTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorSmartContractExecute(ctx context.Context, in *SmartContractExecuteTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorCurrencyEquiv(ctx context.Context, in *SelfCurrencyEquiv, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorAuthCurrencyEquiv(ctx context.Context, in *AuthorizedCurrencyEquiv, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidatorExpenseRatio(ctx context.Context, in *ExpenseRatioTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type validatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidatorServiceClient(cc grpc.ClientConnInterface) ValidatorServiceClient {
	return &validatorServiceClient{cc}
}

func (c *validatorServiceClient) Broadcast(ctx context.Context, in *Block, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) SyncBlockchain(ctx context.Context, in *BlockSync, opts ...grpc.CallOption) (*BlockBatch, error) {
	out := new(BlockBatch)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/SyncBlockchain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorRegistration(ctx context.Context, in *ValidatorRegistrationMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) SyncValidatorList(ctx context.Context, in *ValidatorSyncRequest, opts ...grpc.CallOption) (*ValidatorSync, error) {
	out := new(ValidatorSync)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/SyncValidatorList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorCoin(ctx context.Context, in *CoinTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorMint(ctx context.Context, in *MintTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorMint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorItemMint(ctx context.Context, in *ItemizedMintTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorItemMint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorContract(ctx context.Context, in *InstrumentContract, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorGovernProposal(ctx context.Context, in *GovernanceProposal, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorGovernProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorGovernVote(ctx context.Context, in *GovernanceVote, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorGovernVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorSmartContract(ctx context.Context, in *SmartContractTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorSmartContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorSmartContractExecute(ctx context.Context, in *SmartContractExecuteTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorSmartContractExecute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorCurrencyEquiv(ctx context.Context, in *SelfCurrencyEquiv, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorCurrencyEquiv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorAuthCurrencyEquiv(ctx context.Context, in *AuthorizedCurrencyEquiv, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorAuthCurrencyEquiv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validatorServiceClient) ValidatorExpenseRatio(ctx context.Context, in *ExpenseRatioTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zera_validator.ValidatorService/ValidatorExpenseRatio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorServiceServer is the server API for ValidatorService service.
// All implementations must embed UnimplementedValidatorServiceServer
// for forward compatibility
type ValidatorServiceServer interface {
	Broadcast(context.Context, *Block) (*emptypb.Empty, error)
	SyncBlockchain(context.Context, *BlockSync) (*BlockBatch, error)
	ValidatorRegistration(context.Context, *ValidatorRegistrationMessage) (*emptypb.Empty, error)
	SyncValidatorList(context.Context, *ValidatorSyncRequest) (*ValidatorSync, error)
	ValidatorCoin(context.Context, *CoinTXN) (*emptypb.Empty, error)
	ValidatorMint(context.Context, *MintTXN) (*emptypb.Empty, error)
	ValidatorItemMint(context.Context, *ItemizedMintTXN) (*emptypb.Empty, error)
	ValidatorContract(context.Context, *InstrumentContract) (*emptypb.Empty, error)
	ValidatorGovernProposal(context.Context, *GovernanceProposal) (*emptypb.Empty, error)
	ValidatorGovernVote(context.Context, *GovernanceVote) (*emptypb.Empty, error)
	ValidatorSmartContract(context.Context, *SmartContractTXN) (*emptypb.Empty, error)
	ValidatorSmartContractExecute(context.Context, *SmartContractExecuteTXN) (*emptypb.Empty, error)
	ValidatorCurrencyEquiv(context.Context, *SelfCurrencyEquiv) (*emptypb.Empty, error)
	ValidatorAuthCurrencyEquiv(context.Context, *AuthorizedCurrencyEquiv) (*emptypb.Empty, error)
	ValidatorExpenseRatio(context.Context, *ExpenseRatioTXN) (*emptypb.Empty, error)
	mustEmbedUnimplementedValidatorServiceServer()
}

// UnimplementedValidatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedValidatorServiceServer struct {
}

func (UnimplementedValidatorServiceServer) Broadcast(context.Context, *Block) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedValidatorServiceServer) SyncBlockchain(context.Context, *BlockSync) (*BlockBatch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncBlockchain not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorRegistration(context.Context, *ValidatorRegistrationMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorRegistration not implemented")
}
func (UnimplementedValidatorServiceServer) SyncValidatorList(context.Context, *ValidatorSyncRequest) (*ValidatorSync, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncValidatorList not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorCoin(context.Context, *CoinTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorCoin not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorMint(context.Context, *MintTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorMint not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorItemMint(context.Context, *ItemizedMintTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorItemMint not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorContract(context.Context, *InstrumentContract) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorContract not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorGovernProposal(context.Context, *GovernanceProposal) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorGovernProposal not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorGovernVote(context.Context, *GovernanceVote) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorGovernVote not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorSmartContract(context.Context, *SmartContractTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorSmartContract not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorSmartContractExecute(context.Context, *SmartContractExecuteTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorSmartContractExecute not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorCurrencyEquiv(context.Context, *SelfCurrencyEquiv) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorCurrencyEquiv not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorAuthCurrencyEquiv(context.Context, *AuthorizedCurrencyEquiv) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorAuthCurrencyEquiv not implemented")
}
func (UnimplementedValidatorServiceServer) ValidatorExpenseRatio(context.Context, *ExpenseRatioTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorExpenseRatio not implemented")
}
func (UnimplementedValidatorServiceServer) mustEmbedUnimplementedValidatorServiceServer() {}

// UnsafeValidatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValidatorServiceServer will
// result in compilation errors.
type UnsafeValidatorServiceServer interface {
	mustEmbedUnimplementedValidatorServiceServer()
}

func RegisterValidatorServiceServer(s grpc.ServiceRegistrar, srv ValidatorServiceServer) {
	s.RegisterService(&ValidatorService_ServiceDesc, srv)
}

func _ValidatorService_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).Broadcast(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_SyncBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockSync)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).SyncBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/SyncBlockchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).SyncBlockchain(ctx, req.(*BlockSync))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatorRegistrationMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorRegistration(ctx, req.(*ValidatorRegistrationMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_SyncValidatorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatorSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).SyncValidatorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/SyncValidatorList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).SyncValidatorList(ctx, req.(*ValidatorSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorCoin(ctx, req.(*CoinTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorMint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorMint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorMint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorMint(ctx, req.(*MintTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorItemMint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemizedMintTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorItemMint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorItemMint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorItemMint(ctx, req.(*ItemizedMintTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorContract(ctx, req.(*InstrumentContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorGovernProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GovernanceProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorGovernProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorGovernProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorGovernProposal(ctx, req.(*GovernanceProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorGovernVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GovernanceVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorGovernVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorGovernVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorGovernVote(ctx, req.(*GovernanceVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorSmartContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmartContractTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorSmartContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorSmartContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorSmartContract(ctx, req.(*SmartContractTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorSmartContractExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmartContractExecuteTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorSmartContractExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorSmartContractExecute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorSmartContractExecute(ctx, req.(*SmartContractExecuteTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorCurrencyEquiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelfCurrencyEquiv)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorCurrencyEquiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorCurrencyEquiv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorCurrencyEquiv(ctx, req.(*SelfCurrencyEquiv))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorAuthCurrencyEquiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizedCurrencyEquiv)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorAuthCurrencyEquiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorAuthCurrencyEquiv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorAuthCurrencyEquiv(ctx, req.(*AuthorizedCurrencyEquiv))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidatorService_ValidatorExpenseRatio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseRatioTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidatorExpenseRatio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zera_validator.ValidatorService/ValidatorExpenseRatio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidatorExpenseRatio(ctx, req.(*ExpenseRatioTXN))
	}
	return interceptor(ctx, in, info, handler)
}

// ValidatorService_ServiceDesc is the grpc.ServiceDesc for ValidatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValidatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zera_validator.ValidatorService",
	HandlerType: (*ValidatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _ValidatorService_Broadcast_Handler,
		},
		{
			MethodName: "SyncBlockchain",
			Handler:    _ValidatorService_SyncBlockchain_Handler,
		},
		{
			MethodName: "ValidatorRegistration",
			Handler:    _ValidatorService_ValidatorRegistration_Handler,
		},
		{
			MethodName: "SyncValidatorList",
			Handler:    _ValidatorService_SyncValidatorList_Handler,
		},
		{
			MethodName: "ValidatorCoin",
			Handler:    _ValidatorService_ValidatorCoin_Handler,
		},
		{
			MethodName: "ValidatorMint",
			Handler:    _ValidatorService_ValidatorMint_Handler,
		},
		{
			MethodName: "ValidatorItemMint",
			Handler:    _ValidatorService_ValidatorItemMint_Handler,
		},
		{
			MethodName: "ValidatorContract",
			Handler:    _ValidatorService_ValidatorContract_Handler,
		},
		{
			MethodName: "ValidatorGovernProposal",
			Handler:    _ValidatorService_ValidatorGovernProposal_Handler,
		},
		{
			MethodName: "ValidatorGovernVote",
			Handler:    _ValidatorService_ValidatorGovernVote_Handler,
		},
		{
			MethodName: "ValidatorSmartContract",
			Handler:    _ValidatorService_ValidatorSmartContract_Handler,
		},
		{
			MethodName: "ValidatorSmartContractExecute",
			Handler:    _ValidatorService_ValidatorSmartContractExecute_Handler,
		},
		{
			MethodName: "ValidatorCurrencyEquiv",
			Handler:    _ValidatorService_ValidatorCurrencyEquiv_Handler,
		},
		{
			MethodName: "ValidatorAuthCurrencyEquiv",
			Handler:    _ValidatorService_ValidatorAuthCurrencyEquiv_Handler,
		},
		{
			MethodName: "ValidatorExpenseRatio",
			Handler:    _ValidatorService_ValidatorExpenseRatio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "validator.proto",
}
