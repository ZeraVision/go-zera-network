package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	zera_pb "github.com/ZeraVision/go-zera-network/grpc/protobuf"
)

// struct for client implementation of grpcs
type ValidatorNetworkClient struct {
	client zera_pb.ValidatorServiceClient
}

// struct for client implementation of grpcs
type NetworkClient struct {
	client zera_pb.TXNServiceClient
}

// constructor for client implementation of grpcs
func NewNetworkClient(conn *grpc.ClientConn) *NetworkClient {
	client := zera_pb.NewTXNServiceClient(conn)
	return &NetworkClient{client: client}
}

// constructor for client implementation of grpcs
func NewValidatorNetworkClient(conn *grpc.ClientConn) *ValidatorNetworkClient {
	client := zera_pb.NewValidatorServiceClient(conn)
	return &ValidatorNetworkClient{client: client}
}

// SendBlocksyncRequest sends a block sync request and returns the received block batch.
func SendBlocksyncRequest(blockSync *zera_pb.BlockSync, destAddr string) (*zera_pb.BlockBatch, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure()) // 137.184.181.31 dev // 146.190.163.205 proto // 146.190.139.36
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
func (v *ValidatorNetworkClient) SyncBlockchain(ctx context.Context, request *zera_pb.BlockSync) (*zera_pb.BlockBatch, error) {
	response, err := v.client.SyncBlockchain(ctx, request)
	return response, err
}

// SendMintTXN sends a mint txn and returns an empty response.
func SendMintTXN(mintTXN *zera_pb.MintTXN, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
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
