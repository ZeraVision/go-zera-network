package client

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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

// struct for client implementation of grpcs
type APIClient struct {
	client zera_pb.APIServiceClient
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

// constructor for client implementation of grpcs
func APINetworkClient(conn *grpc.ClientConn) *APIClient {
	client := zera_pb.NewAPIServiceClient(conn)
	return &APIClient{client: client}
}

// SendBlocksyncRequest sends a block sync request and returns the received block batch.
func SendBlocksyncRequest(blockSync *zera_pb.BlockSync, destAddr string) (*zera_pb.BlockBatch, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewValidatorNetworkClient(conn)

	stream, err := client.client.SyncBlockchain(context.Background(), blockSync)

	if err != nil {
		fmt.Printf("Failed to call SyncBlockchain: %v", err)
		return nil, err
	}

	var aggregatedData []byte // use bytes.Buffer or similar for performance in a real-world scenario
	for {
		batchChunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("SendBlocksyncRequest: Failed to receive a batch chunk: %v", err)
			return nil, err
		}
		aggregatedData = append(aggregatedData, batchChunk.GetChunkData()...)
	}

	blockBatch := &zera_pb.BlockBatch{}
	if err := proto.Unmarshal(aggregatedData, blockBatch); err != nil {
		fmt.Printf("SendBlocksyncRequest: Failed to deserialize BlockBatch: %v", err)
		return nil, err
	}

	return blockBatch, nil
}

// SendMintTXN sends a mint txn and returns an empty response.
func SendMintTXN(mintTXN *zera_pb.MintTXN, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
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

func SendNFTTXN(nftTxn *zera_pb.NFTTXN, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewNetworkClient(conn)

	response, err := client.client.NFT(context.Background(), nftTxn)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func SendItemMintTXN(mintTXN *zera_pb.ItemizedMintTXN, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewNetworkClient(conn)

	response, err := client.client.ItemMint(context.Background(), mintTXN)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func SendGovernanceProposal(proposal *zera_pb.GovernanceProposal, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewNetworkClient(conn)

	response, err := client.client.GovernProposal(context.Background(), proposal)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func SendGovernanceVote(vote *zera_pb.GovernanceVote, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewNetworkClient(conn)

	response, err := client.client.GovernVote(context.Background(), vote)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func SendKyc(compliance *zera_pb.ComplianceTXN, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewNetworkClient(conn)

	response, err := client.client.Compliance(context.Background(), compliance)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func SendExpenseRatio(expenseRatio *zera_pb.ExpenseRatioTXN, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewNetworkClient(conn)

	response, err := client.client.ExpenseRatio(context.Background(), expenseRatio)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func SendCoinTXN(coin *zera_pb.CoinTXN, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to the server: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	client := NewNetworkClient(conn)

	response, err := client.client.Coin(context.Background(), coin)

	if err != nil {
		return nil, err
	}

	return response, nil
}



