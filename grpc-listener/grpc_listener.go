package grpc_network_listener

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	zera_grpc "github.com/ZeraVision/go-zera-network/zera-grpc"
)

// Define handler functions for each gRPC method.
type BroadcastHandler func(ctx context.Context, block *zera_grpc.Block) (*emptypb.Empty, error)
type SyncBlockchainHandler func(ctx context.Context, blockSync *zera_grpc.BlockSync) (*zera_grpc.BlockBatch, error)
type ValidatorRegistrationHandler func(ctx context.Context, message *zera_grpc.ValidatorRegistrationMessage) (*emptypb.Empty, error)
type SyncValidatorListHandler func(ctx context.Context, request *zera_grpc.ValidatorSyncRequest) (*zera_grpc.ValidatorSync, error)
type ValidatorCoinHandler func(ctx context.Context, txn *zera_grpc.CoinTXN) (*emptypb.Empty, error)
type ValidatorMintHandler func(ctx context.Context, txn *zera_grpc.MintTXN) (*emptypb.Empty, error)
type ValidatorItemMintHandler func(ctx context.Context, txn *zera_grpc.ItemizedMintTXN) (*emptypb.Empty, error)
type ValidatorContractHandler func(ctx context.Context, contract *zera_grpc.InstrumentContract) (*emptypb.Empty, error)
type ValidatorGovernProposalHandler func(ctx context.Context, proposal *zera_grpc.GovernanceProposal) (*emptypb.Empty, error)
type ValidatorGovernVoteHandler func(ctx context.Context, vote *zera_grpc.GovernanceVote) (*emptypb.Empty, error)
type ValidatorSmartContractHandler func(ctx context.Context, txn *zera_grpc.SmartContractTXN) (*emptypb.Empty, error)
type ValidatorSmartContractExecuteHandler func(ctx context.Context, txn *zera_grpc.SmartContractExecuteTXN) (*emptypb.Empty, error)
type ValidatorCurrencyEquivHandler func(ctx context.Context, equiv *zera_grpc.SelfCurrencyEquiv) (*emptypb.Empty, error)
type ValidatorAuthCurrencyEquivHandler func(ctx context.Context, equiv *zera_grpc.AuthorizedCurrencyEquiv) (*emptypb.Empty, error)
type ValidatorExpenseRatioHandler func(ctx context.Context, ratio *zera_grpc.ExpenseRatioTXN) (*emptypb.Empty, error)

// ValidatorService is an implementation of the ValidatorServiceClient interface.
type ValidatorService struct {
	HandleBroadcast                  BroadcastHandler
	HandleSyncBlockchain             SyncBlockchainHandler
	HandleValidatorRegistration      ValidatorRegistrationHandler
	HandleSyncValidatorList          SyncValidatorListHandler
	HandleValidatorCoin              ValidatorCoinHandler
	HandleValidatorMint              ValidatorMintHandler
	HandleValidatorItemMint          ValidatorItemMintHandler
	HandleValidatorContract          ValidatorContractHandler
	HandleValidatorGovernProposal    ValidatorGovernProposalHandler
	HandleValidatorGovernVote        ValidatorGovernVoteHandler
	HandleValidatorSmartContract     ValidatorSmartContractHandler
	HandleValidatorSmartContractExec ValidatorSmartContractExecuteHandler
	HandleValidatorCurrencyEquiv     ValidatorCurrencyEquivHandler
	HandleValidatorAuthCurrencyEquiv ValidatorAuthCurrencyEquivHandler
	HandleValidatorExpenseRatio      ValidatorExpenseRatioHandler
	zera_grpc.UnimplementedValidatorServiceServer
	Service *grpc.Server
}

func (c *ValidatorService) StartService() {
	// Create a gRPC server
	c.Service = grpc.NewServer()

	// Register the ValidatorNetworkServer with the gRPC server
	zera_grpc.RegisterValidatorServiceServer(c.Service, c) //! TODO from network?

	// Start listening on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	// Start the gRPC server
	fmt.Println("Starting gRPC server on port 50051")
	if err := c.Service.Serve(listener); err != nil {
		fmt.Printf("Failed to server: %v", err)
	}

}

// NewValidatorService creates a new instance of ValidatorService.
func NewValidatorService() *ValidatorService {
	return &ValidatorService{}
}

// Implement each of the methods by delegating to the corresponding handler functions.

func (c *ValidatorService) Broadcast(ctx context.Context, in *zera_grpc.Block) (*emptypb.Empty, error) {
	return c.HandleBroadcast(ctx, in)
}

func (c *ValidatorService) SyncBlockchain(ctx context.Context, in *zera_grpc.BlockSync) (*zera_grpc.BlockBatch, error) {
	return c.HandleSyncBlockchain(ctx, in)
}

func (c *ValidatorService) ValidatorRegistration(ctx context.Context, in *zera_grpc.ValidatorRegistrationMessage) (*emptypb.Empty, error) {
	return c.HandleValidatorRegistration(ctx, in)
}

func (c *ValidatorService) SyncValidatorList(ctx context.Context, in *zera_grpc.ValidatorSyncRequest) (*zera_grpc.ValidatorSync, error) {
	return c.HandleSyncValidatorList(ctx, in)
}

func (c *ValidatorService) ValidatorCoin(ctx context.Context, in *zera_grpc.CoinTXN) (*emptypb.Empty, error) {
	return c.HandleValidatorCoin(ctx, in)
}

func (c *ValidatorService) ValidatorMint(ctx context.Context, in *zera_grpc.MintTXN) (*emptypb.Empty, error) {
	return c.HandleValidatorMint(ctx, in)
}

func (c *ValidatorService) ValidatorItemMint(ctx context.Context, in *zera_grpc.ItemizedMintTXN) (*emptypb.Empty, error) {
	return c.HandleValidatorItemMint(ctx, in)
}

func (c *ValidatorService) ValidatorContract(ctx context.Context, in *zera_grpc.InstrumentContract) (*emptypb.Empty, error) {
	return c.HandleValidatorContract(ctx, in)
}

func (c *ValidatorService) ValidatorGovernProposal(ctx context.Context, in *zera_grpc.GovernanceProposal) (*emptypb.Empty, error) {
	return c.HandleValidatorGovernProposal(ctx, in)
}

func (c *ValidatorService) ValidatorGovernVote(ctx context.Context, in *zera_grpc.GovernanceVote) (*emptypb.Empty, error) {
	return c.HandleValidatorGovernVote(ctx, in)
}

func (c *ValidatorService) ValidatorSmartContract(ctx context.Context, in *zera_grpc.SmartContractTXN) (*emptypb.Empty, error) {
	return c.HandleValidatorSmartContract(ctx, in)
}

func (c *ValidatorService) ValidatorSmartContractExecute(ctx context.Context, in *zera_grpc.SmartContractExecuteTXN) (*emptypb.Empty, error) {
	return c.HandleValidatorSmartContractExec(ctx, in)
}

func (c *ValidatorService) ValidatorCurrencyEquiv(ctx context.Context, in *zera_grpc.SelfCurrencyEquiv) (*emptypb.Empty, error) {
	return c.HandleValidatorCurrencyEquiv(ctx, in)
}

func (c *ValidatorService) ValidatorAuthCurrencyEquiv(ctx context.Context, in *zera_grpc.AuthorizedCurrencyEquiv) (*emptypb.Empty, error) {
	return c.HandleValidatorAuthCurrencyEquiv(ctx, in)
}

func (c *ValidatorService) ValidatorExpenseRatio(ctx context.Context, in *zera_grpc.ExpenseRatioTXN) (*emptypb.Empty, error) {
	return c.HandleValidatorExpenseRatio(ctx, in)
}
