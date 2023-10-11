// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc1
// source: txn.proto

package protobuf

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

const (
	TXNService_Coin_FullMethodName                 = "/zera_txn.TXNService/Coin"
	TXNService_Mint_FullMethodName                 = "/zera_txn.TXNService/Mint"
	TXNService_ItemMint_FullMethodName             = "/zera_txn.TXNService/ItemMint"
	TXNService_Contract_FullMethodName             = "/zera_txn.TXNService/Contract"
	TXNService_GovernProposal_FullMethodName       = "/zera_txn.TXNService/GovernProposal"
	TXNService_GovernVote_FullMethodName           = "/zera_txn.TXNService/GovernVote"
	TXNService_SmartContract_FullMethodName        = "/zera_txn.TXNService/SmartContract"
	TXNService_SmartContractExecute_FullMethodName = "/zera_txn.TXNService/SmartContractExecute"
	TXNService_CurrencyEquiv_FullMethodName        = "/zera_txn.TXNService/CurrencyEquiv"
	TXNService_AuthCurrencyEquiv_FullMethodName    = "/zera_txn.TXNService/AuthCurrencyEquiv"
	TXNService_ExpenseRatio_FullMethodName         = "/zera_txn.TXNService/ExpenseRatio"
)

// TXNServiceClient is the client API for TXNService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TXNServiceClient interface {
	Coin(ctx context.Context, in *CoinTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Mint(ctx context.Context, in *MintTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ItemMint(ctx context.Context, in *ItemizedMintTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Contract(ctx context.Context, in *InstrumentContract, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GovernProposal(ctx context.Context, in *GovernanceProposal, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GovernVote(ctx context.Context, in *GovernanceVote, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SmartContract(ctx context.Context, in *SmartContractTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SmartContractExecute(ctx context.Context, in *SmartContractExecuteTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CurrencyEquiv(ctx context.Context, in *SelfCurrencyEquiv, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AuthCurrencyEquiv(ctx context.Context, in *AuthorizedCurrencyEquiv, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ExpenseRatio(ctx context.Context, in *ExpenseRatioTXN, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tXNServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTXNServiceClient(cc grpc.ClientConnInterface) TXNServiceClient {
	return &tXNServiceClient{cc}
}

func (c *tXNServiceClient) Coin(ctx context.Context, in *CoinTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_Coin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) Mint(ctx context.Context, in *MintTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_Mint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) ItemMint(ctx context.Context, in *ItemizedMintTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_ItemMint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) Contract(ctx context.Context, in *InstrumentContract, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_Contract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) GovernProposal(ctx context.Context, in *GovernanceProposal, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_GovernProposal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) GovernVote(ctx context.Context, in *GovernanceVote, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_GovernVote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) SmartContract(ctx context.Context, in *SmartContractTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_SmartContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) SmartContractExecute(ctx context.Context, in *SmartContractExecuteTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_SmartContractExecute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) CurrencyEquiv(ctx context.Context, in *SelfCurrencyEquiv, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_CurrencyEquiv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) AuthCurrencyEquiv(ctx context.Context, in *AuthorizedCurrencyEquiv, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_AuthCurrencyEquiv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXNServiceClient) ExpenseRatio(ctx context.Context, in *ExpenseRatioTXN, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TXNService_ExpenseRatio_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TXNServiceServer is the server API for TXNService service.
// All implementations must embed UnimplementedTXNServiceServer
// for forward compatibility
type TXNServiceServer interface {
	Coin(context.Context, *CoinTXN) (*emptypb.Empty, error)
	Mint(context.Context, *MintTXN) (*emptypb.Empty, error)
	ItemMint(context.Context, *ItemizedMintTXN) (*emptypb.Empty, error)
	Contract(context.Context, *InstrumentContract) (*emptypb.Empty, error)
	GovernProposal(context.Context, *GovernanceProposal) (*emptypb.Empty, error)
	GovernVote(context.Context, *GovernanceVote) (*emptypb.Empty, error)
	SmartContract(context.Context, *SmartContractTXN) (*emptypb.Empty, error)
	SmartContractExecute(context.Context, *SmartContractExecuteTXN) (*emptypb.Empty, error)
	CurrencyEquiv(context.Context, *SelfCurrencyEquiv) (*emptypb.Empty, error)
	AuthCurrencyEquiv(context.Context, *AuthorizedCurrencyEquiv) (*emptypb.Empty, error)
	ExpenseRatio(context.Context, *ExpenseRatioTXN) (*emptypb.Empty, error)
	mustEmbedUnimplementedTXNServiceServer()
}

// UnimplementedTXNServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTXNServiceServer struct {
}

func (UnimplementedTXNServiceServer) Coin(context.Context, *CoinTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Coin not implemented")
}
func (UnimplementedTXNServiceServer) Mint(context.Context, *MintTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mint not implemented")
}
func (UnimplementedTXNServiceServer) ItemMint(context.Context, *ItemizedMintTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ItemMint not implemented")
}
func (UnimplementedTXNServiceServer) Contract(context.Context, *InstrumentContract) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Contract not implemented")
}
func (UnimplementedTXNServiceServer) GovernProposal(context.Context, *GovernanceProposal) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovernProposal not implemented")
}
func (UnimplementedTXNServiceServer) GovernVote(context.Context, *GovernanceVote) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovernVote not implemented")
}
func (UnimplementedTXNServiceServer) SmartContract(context.Context, *SmartContractTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmartContract not implemented")
}
func (UnimplementedTXNServiceServer) SmartContractExecute(context.Context, *SmartContractExecuteTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmartContractExecute not implemented")
}
func (UnimplementedTXNServiceServer) CurrencyEquiv(context.Context, *SelfCurrencyEquiv) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrencyEquiv not implemented")
}
func (UnimplementedTXNServiceServer) AuthCurrencyEquiv(context.Context, *AuthorizedCurrencyEquiv) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthCurrencyEquiv not implemented")
}
func (UnimplementedTXNServiceServer) ExpenseRatio(context.Context, *ExpenseRatioTXN) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpenseRatio not implemented")
}
func (UnimplementedTXNServiceServer) mustEmbedUnimplementedTXNServiceServer() {}

// UnsafeTXNServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TXNServiceServer will
// result in compilation errors.
type UnsafeTXNServiceServer interface {
	mustEmbedUnimplementedTXNServiceServer()
}

func RegisterTXNServiceServer(s grpc.ServiceRegistrar, srv TXNServiceServer) {
	s.RegisterService(&TXNService_ServiceDesc, srv)
}

func _TXNService_Coin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).Coin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_Coin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).Coin(ctx, req.(*CoinTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_Mint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).Mint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_Mint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).Mint(ctx, req.(*MintTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_ItemMint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemizedMintTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).ItemMint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_ItemMint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).ItemMint(ctx, req.(*ItemizedMintTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_Contract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).Contract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_Contract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).Contract(ctx, req.(*InstrumentContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_GovernProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GovernanceProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).GovernProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_GovernProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).GovernProposal(ctx, req.(*GovernanceProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_GovernVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GovernanceVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).GovernVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_GovernVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).GovernVote(ctx, req.(*GovernanceVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_SmartContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmartContractTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).SmartContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_SmartContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).SmartContract(ctx, req.(*SmartContractTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_SmartContractExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmartContractExecuteTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).SmartContractExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_SmartContractExecute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).SmartContractExecute(ctx, req.(*SmartContractExecuteTXN))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_CurrencyEquiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelfCurrencyEquiv)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).CurrencyEquiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_CurrencyEquiv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).CurrencyEquiv(ctx, req.(*SelfCurrencyEquiv))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_AuthCurrencyEquiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizedCurrencyEquiv)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).AuthCurrencyEquiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_AuthCurrencyEquiv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).AuthCurrencyEquiv(ctx, req.(*AuthorizedCurrencyEquiv))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXNService_ExpenseRatio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseRatioTXN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXNServiceServer).ExpenseRatio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TXNService_ExpenseRatio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXNServiceServer).ExpenseRatio(ctx, req.(*ExpenseRatioTXN))
	}
	return interceptor(ctx, in, info, handler)
}

// TXNService_ServiceDesc is the grpc.ServiceDesc for TXNService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TXNService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zera_txn.TXNService",
	HandlerType: (*TXNServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Coin",
			Handler:    _TXNService_Coin_Handler,
		},
		{
			MethodName: "Mint",
			Handler:    _TXNService_Mint_Handler,
		},
		{
			MethodName: "ItemMint",
			Handler:    _TXNService_ItemMint_Handler,
		},
		{
			MethodName: "Contract",
			Handler:    _TXNService_Contract_Handler,
		},
		{
			MethodName: "GovernProposal",
			Handler:    _TXNService_GovernProposal_Handler,
		},
		{
			MethodName: "GovernVote",
			Handler:    _TXNService_GovernVote_Handler,
		},
		{
			MethodName: "SmartContract",
			Handler:    _TXNService_SmartContract_Handler,
		},
		{
			MethodName: "SmartContractExecute",
			Handler:    _TXNService_SmartContractExecute_Handler,
		},
		{
			MethodName: "CurrencyEquiv",
			Handler:    _TXNService_CurrencyEquiv_Handler,
		},
		{
			MethodName: "AuthCurrencyEquiv",
			Handler:    _TXNService_AuthCurrencyEquiv_Handler,
		},
		{
			MethodName: "ExpenseRatio",
			Handler:    _TXNService_ExpenseRatio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "txn.proto",
}
