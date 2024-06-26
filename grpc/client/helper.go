package client

import (
	"context"
	"fmt"

	zera_pb "github.com/ZeraVision/go-zera-network/grpc/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)


// SendMintTXN sends a mint txn and returns an empty response.
func SendCreateContractTXN(contract *zera_pb.InstrumentContract, destAddr string) (*emptypb.Empty, error) {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial(destAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to connect to the server: " + err.Error())
		return nil, err
	}
	defer conn.Close()

	// Create a new instance of ValidatorNetworkClient
	//var client *grpc_client.NetworkClient
	client := NewNetworkClient(conn)

	response, err := client.client.Contract(context.Background(), contract)

	if err != nil {
		fmt.Println("Failed to connect to the server: ", err.Error())
		return nil, err
	}
	return response, nil
}
