package api

import (
	"context"
	"fmt"

	apirpc "onboarding.com/api/grpcmodules"
	numrpc "onboarding.com/number/grpcmodules"

	"google.golang.org/grpc"
)

type apiRpcServer struct {
	client numrpc.NumberRpcClient
}

func NewRpcServer() (*apiRpcServer, error) {
	conn, err := grpc.Dial("localhost:1000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	client := numrpc.NewNumberRpcClient(conn)
	return &apiRpcServer{client: client}, nil
}

func (s *apiRpcServer) GuessNumber(ctx context.Context, guessStream apirpc.ApiRpc_GuessNumberServer) error {

	numStream, err := s.client.IsNumberExist(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for {
		guess, err := guessStream.Recv()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		numStream.Send(&numrpc.GuessNumber{Num: guess.GetNum(), Id: guess.GetId()})
		// numSream.Recv()
	}
}
