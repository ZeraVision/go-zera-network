package grpc_network_client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	zera_grpc "github.com/ZeraVision/go-zera-network-grpc/zera-grpc"
)

const GRPC_NETWORK_HIT_ADDRESS = "dev-validator-grpc.zera.vision:50051"

// struct for client implementation of grpcs
type ValidatorNetworkClient struct {
	client zera_grpc.ValidatorServiceClient
}

// struct for client implementation of grpcs
type NetworkClient struct {
	client zera_grpc.TXNServiceClient
}

// constructor for client implementation of grpcs
func NewNetworkClient(conn *grpc.ClientConn) *NetworkClient {
	client := zera_grpc.NewTXNServiceClient(conn)
	return &NetworkClient{client: client}
}

// constructor for client implementation of grpcs
func NewValidatorNetworkClient(conn *grpc.ClientConn) *ValidatorNetworkClient {
	client := zera_grpc.NewValidatorServiceClient(conn)
	return &ValidatorNetworkClient{client: client}
}

// SendBlocksyncRequest sends a block sync request and returns the received block batch.
func SendBlocksyncRequest(blockSync *zera_grpc.BlockSync) (*zera_grpc.BlockBatch, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(GRPC_NETWORK_HIT_ADDRESS, grpc.WithInsecure()) // 137.184.181.31 dev // 146.190.163.205 proto // 146.190.139.36
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewValidatorNetworkClient(conn)

	// Call the SyncBlockchain RPC
	blockBatch, err := client.SyncBlockchain(context.Background(), blockSync)
	if err != nil {
		return blockBatch, nil
	}

	return blockBatch, nil
}

// SyncBlockchain sends a block sync request and returns the received block batch.
func (v *ValidatorNetworkClient) SyncBlockchain(ctx context.Context, request *zera_grpc.BlockSync) (*zera_grpc.BlockBatch, error) {
	response, err := v.client.SyncBlockchain(ctx, request)
	return response, err
}

// SendMintTXN sends a mint txn and returns an empty response.
func SendMintTXN(mintTXN *zera_grpc.MintTXN) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial("proto-txn-sender-v2.zera.vision:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewNetworkClient(conn)

	response, err := client.client.Mint(context.Background(), mintTXN)

	if err != nil {
		return nil, err
	}

	return response, nil
}
