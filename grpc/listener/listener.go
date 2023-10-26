package listener

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	zera_pb "github.com/ZeraVision/go-zera-network/grpc/protobuf"
)

// Define handler functions for each gRPC method.
type BroadcastHandler func(ctx context.Context, block *zera_pb.Block) (*emptypb.Empty, error)
type SyncBlockchainHandler func(ctx context.Context, blockSync *zera_pb.BlockSync) (*zera_pb.BlockBatch, error)
type ValidatorRegistrationHandler func(ctx context.Context, message *zera_pb.ValidatorRegistration) (*emptypb.Empty, error)
type SyncValidatorListHandler func(ctx context.Context, request *zera_pb.ValidatorSyncRequest) (*zera_pb.ValidatorSync, error)
type ValidatorCoinHandler func(ctx context.Context, txn *zera_pb.CoinTXN) (*emptypb.Empty, error)
type ValidatorMintHandler func(ctx context.Context, txn *zera_pb.MintTXN) (*emptypb.Empty, error)
type ValidatorItemMintHandler func(ctx context.Context, txn *zera_pb.ItemizedMintTXN) (*emptypb.Empty, error)
type ValidatorContractHandler func(ctx context.Context, contract *zera_pb.InstrumentContract) (*emptypb.Empty, error)
type ValidatorGovernProposalHandler func(ctx context.Context, proposal *zera_pb.GovernanceProposal) (*emptypb.Empty, error)
type ValidatorGovernVoteHandler func(ctx context.Context, vote *zera_pb.GovernanceVote) (*emptypb.Empty, error)
type ValidatorSmartContractHandler func(ctx context.Context, txn *zera_pb.SmartContractTXN) (*emptypb.Empty, error)
type ValidatorSmartContractExecuteHandler func(ctx context.Context, txn *zera_pb.SmartContractExecuteTXN) (*emptypb.Empty, error)
type ValidatorCurrencyEquivHandler func(ctx context.Context, equiv *zera_pb.SelfCurrencyEquiv) (*emptypb.Empty, error)
type ValidatorAuthCurrencyEquivHandler func(ctx context.Context, equiv *zera_pb.AuthorizedCurrencyEquiv) (*emptypb.Empty, error)
type ValidatorExpenseRatioHandler func(ctx context.Context, ratio *zera_pb.ExpenseRatioTXN) (*emptypb.Empty, error)

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
	zera_pb.UnimplementedValidatorServiceServer
	Service *grpc.Server
}

func (c *ValidatorService) StartService() {
	// Create a gRPC server
	c.Service = grpc.NewServer()

	// Register the ValidatorNetworkServer with the gRPC server
	zera_pb.RegisterValidatorServiceServer(c.Service, c) //! TODO from network?

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
func (c *ValidatorService) Broadcast(ctx context.Context, in *zera_pb.Block) (*emptypb.Empty, error) {
	if c == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleBroadcast(ctx, in)
}

// Implement each of the methods by delegating to the corresponding handler functions.
func (c *ValidatorService) StreamBroadcast(stream zera_pb.ValidatorService_StreamBroadcastServer) error {

	var aggregatedData []byte // use bytes.Buffer or similar for performance in a real-world scenario
	// Listen for messages from client
	for {
		dataChunk, err := stream.Recv()
		if err == io.EOF {
			// All data has been received
			break
		}
		if err != nil {
			log.Fatalf("StreamBroadcast: Failed to receive a block chunk: %v", err)
			return err
		}
		aggregatedData = append(aggregatedData, dataChunk.GetChunkData()...)
	}
	// Close the stream immediately after receipt of all data
	err := stream.SendAndClose(&emptypb.Empty{})
	if err != nil {
		log.Fatalf("StreamBroadcast: Failed to close the stream: %v", err)
		return err
	}

	block := &zera_pb.Block{}
	if err := proto.Unmarshal(aggregatedData, block); err != nil {
		log.Fatalf("SendBlocksyncRequest: Failed to deserialize BlockBatch: %v", err)
		return err
	}

	_, err1 := c.HandleBroadcast(stream.Context(), block)

	return err1
}

// // func (c *ValidatorService) SyncBlockchain(ctx context.Context, in *zera_pb.BlockSync) (*zera_pb.BlockBatch, error) {
// // 	if c.HandleSyncBlockchain == nil {
// // 		return nil, errors.New("ValidatorService is not initialized")
// // 	}
// // 	return c.HandleSyncBlockchain(ctx, in)
// // }

func (c *ValidatorService) ValidatorRegistration(ctx context.Context, in *zera_pb.ValidatorRegistration) (*emptypb.Empty, error) {
	if c.HandleValidatorRegistration == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorRegistration(ctx, in)
}

func (c *ValidatorService) SyncValidatorList(ctx context.Context, in *zera_pb.ValidatorSyncRequest) (*zera_pb.ValidatorSync, error) {
	if c.HandleSyncValidatorList == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleSyncValidatorList(ctx, in)
}

func (c *ValidatorService) ValidatorCoin(ctx context.Context, in *zera_pb.CoinTXN) (*emptypb.Empty, error) {
	if c.HandleValidatorCoin == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorCoin(ctx, in)
}

func (c *ValidatorService) ValidatorMint(ctx context.Context, in *zera_pb.MintTXN) (*emptypb.Empty, error) {
	if c.HandleValidatorMint == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorMint(ctx, in)
}

func (c *ValidatorService) ValidatorItemMint(ctx context.Context, in *zera_pb.ItemizedMintTXN) (*emptypb.Empty, error) {
	if c.HandleValidatorItemMint == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorItemMint(ctx, in)
}

func (c *ValidatorService) ValidatorContract(ctx context.Context, in *zera_pb.InstrumentContract) (*emptypb.Empty, error) {
	if c.HandleValidatorContract == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorContract(ctx, in)
}

func (c *ValidatorService) ValidatorGovernProposal(ctx context.Context, in *zera_pb.GovernanceProposal) (*emptypb.Empty, error) {
	if c.HandleValidatorGovernProposal == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorGovernProposal(ctx, in)
}

func (c *ValidatorService) ValidatorGovernVote(ctx context.Context, in *zera_pb.GovernanceVote) (*emptypb.Empty, error) {
	if c.HandleValidatorGovernVote == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorGovernVote(ctx, in)
}

func (c *ValidatorService) ValidatorSmartContract(ctx context.Context, in *zera_pb.SmartContractTXN) (*emptypb.Empty, error) {
	if c.HandleValidatorSmartContract == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorSmartContract(ctx, in)
}

func (c *ValidatorService) ValidatorSmartContractExecute(ctx context.Context, in *zera_pb.SmartContractExecuteTXN) (*emptypb.Empty, error) {
	if c.HandleValidatorSmartContractExec == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorSmartContractExec(ctx, in)
}

func (c *ValidatorService) ValidatorCurrencyEquiv(ctx context.Context, in *zera_pb.SelfCurrencyEquiv) (*emptypb.Empty, error) {
	if c.HandleValidatorCurrencyEquiv == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorCurrencyEquiv(ctx, in)
}

func (c *ValidatorService) ValidatorAuthCurrencyEquiv(ctx context.Context, in *zera_pb.AuthorizedCurrencyEquiv) (*emptypb.Empty, error) {
	if c.HandleValidatorAuthCurrencyEquiv == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorAuthCurrencyEquiv(ctx, in)
}

func (c *ValidatorService) ValidatorExpenseRatio(ctx context.Context, in *zera_pb.ExpenseRatioTXN) (*emptypb.Empty, error) {
	if c.HandleValidatorExpenseRatio == nil {
		return nil, errors.New("ValidatorService is not initialized")
	}
	return c.HandleValidatorExpenseRatio(ctx, in)
}
