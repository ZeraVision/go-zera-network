package grpc_network_listener

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	zera_grpc "github.com/ZeraVision/go-zera-network-grpc/zera-grpc"
)

// struct for server implementation of grpcs
type ValidatorNetworkServer struct {
	zera_grpc.UnimplementedValidatorServiceServer //! TODO from network?
}

func StartgRPCService() {
	// Create a gRPC server
	server := grpc.NewServer()
	validatorNetworkServer := &ValidatorNetworkServer{}

	// Register the ValidatorNetworkServer with the gRPC server
	zera_grpc.RegisterValidatorServiceServer(server, validatorNetworkServer) //! TODO from network?

	// Start listening on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	// Start the gRPC server
	fmt.Println("Starting gRPC server on port 50051")
	if err := server.Serve(listener); err != nil {
		fmt.Printf("Failed to server: %v", err)
	}

}

// listener service for broadcasts (new blocks)
func (s *ValidatorNetworkServer) Broadcast(ctx context.Context, block *zera_grpc.Block) (*emptypb.Empty, error) {

	//TODO abstract to not take input method and not require db connection
	// if block_processing_vars.IsBlockSyncActive() {
	// 	return &emptypb.Empty{}, nil
	// }

	// // Block being sent from server doesn't contain hash (fixed in dev), so sync blocks for now // TODO network
	//go block_processing_sync.DoBlockSync(database.GetEnv())

	//go block_processing_insert.RecieveNewBlock()

	return &emptypb.Empty{}, nil

	// TODO implement when direct data is corrected

	// Check if blocks are up to date
	// if bytes.Compare(block.BlockHeader.GetPreviousBlockHash(), GetLastBlockHashByte(dbEnv)) != 0 && !isBlockSyncActive {
	// 	fmt.Println("Behind Blocks, Syncing Blocks!")
	// 	DoBlockSync(dbEnv)
	// } else {
	// 	// insert block
	// 	fmt.Print("Block Recieved: ")
	// 	fmt.Println(block.BlockHeader.GetBlockHeight())
	// 	go InsertBlock(block, dbEnv)
	// }

	// return &emptypb.Empty{}, nil

}

// listener service for new validators
func (s *ValidatorNetworkServer) ValidatorRegistration(ctx context.Context, valReg *zera_grpc.ValidatorRegistrationMessage) (*emptypb.Empty, error) {
	// TODO: Implement the logic for the ValidatorRegistration RPC
	return &emptypb.Empty{}, nil
}

//TODO - May need to delete, dont think we need this functionality

// func (s *ValidatorNetworkServer) SyncValidatorList(context.Context, *zera_grpc.ValidatorSyncRequest) (*zera_grpc.ValidatorSync, error) {
// 	// TODO: Implement the logic for the SyncValidatorList RPC
// 	return &zera_grpc.ValidatorSync{}, nil
// }

// listener service for coin txns
func (s *ValidatorNetworkServer) ValidatorCoin(ctx context.Context, txn *zera_grpc.CoinTXN) (*emptypb.Empty, error) {
	// TODO: Implement the logic for the ValidatorCoin RPC
	return &emptypb.Empty{}, nil
}

//TODO - Will add this service when functionality is complete on validator side

func (s *ValidatorNetworkServer) ValidatorMint(ctx context.Context, txn *zera_grpc.MintTXN) (*emptypb.Empty, error) {
	// TODO: Implement the logic for the ValidatorMint RPC
	return &emptypb.Empty{}, nil
}

//TODO - Will add this service when functionality is complete on validator side

func (s *ValidatorNetworkServer) ValidatorContract(ctx context.Context, txn *zera_grpc.InstrumentContract) (*emptypb.Empty, error) {
	// TODO: Implement the logic for the ValidatorContract RPC
	return &emptypb.Empty{}, nil
}

func (s *ValidatorNetworkServer) ValidatorGovernProposal(ctx context.Context, txn *zera_grpc.GovernanceProposal) (*emptypb.Empty, error) {
	// TODO: Implement the logic for the ValidatorContract RPC
	return &emptypb.Empty{}, nil
}

func (s *ValidatorNetworkServer) ValidatorGovernVoting(ctx context.Context, txn *zera_grpc.GovernanceVote) (*emptypb.Empty, error) {
	// TODO: Implement the logic for the ValidatorContract RPC
	return &emptypb.Empty{}, nil
}
