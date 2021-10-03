package api

import (
	"context"
	"fmt"
	"net"
	"os"

	apirpc "onboarding.com/api/grpcmodules"
	numrpc "onboarding.com/number/grpcmodules"

	"google.golang.org/grpc"
)

type rpcServer struct {
	apirpc.UnimplementedApiRpcServer
	client numrpc.NumberRpcClient
}

func NewRpcServer() {
	numPort := os.Getenv("NUMBER_GRPC_PORT")
	conn, err := grpc.Dial(":"+numPort, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	client := numrpc.NewNumberRpcClient(conn)
	apiRpcServer := rpcServer{client: client}

	grpcPort := os.Getenv("API_GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := grpc.NewServer()
	apirpc.RegisterApiRpcServer(s, &apiRpcServer)
	s.Serve(lis)
}

func (s *rpcServer) GuessNumber(guessStream apirpc.ApiRpc_GuessNumberServer) error {

	numStream, err := s.client.IsExist(context.Background())
	if err != nil {
		fmt.Println("IsExist get stream failed ", err.Error())
		return err
	}

	for {
		guess, err := guessStream.Recv()
		if err != nil {
			fmt.Println("guessStream Recv error: ", err.Error())
			return err
		}

		numStream.Send(&numrpc.GuessNumber{Num: guess.GetNum(), Id: guess.GetId()})
		numStream.Recv()
	}
}
