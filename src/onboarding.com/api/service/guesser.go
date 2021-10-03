package service

import (
	"context"
	"fmt"
	"os"

	rpc "onboarding.com/guesser/grpcmodules"

	"google.golang.org/grpc"
)

type guessService struct {
	client rpc.GuesserRpcClient
}

func NewGuessService() (*guessService, error) {
	grpcPort := os.Getenv("GUESSER_GRPC_PORT")
	conn, err := grpc.Dial(":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	rpcClient := rpc.NewGuesserRpcClient(conn)
	return &guessService{client: rpcClient}, nil
}

func (s *guessService) Add(beginAt, incrementBy, sleepInterval uint32) error {
	gusserMessage := rpc.Guesser{
		BeginAt:       beginAt,
		IncrementBy:   incrementBy,
		SleepInterval: sleepInterval}
	_, err := s.client.Add(context.Background(), &gusserMessage)
	if err != nil {
		return err
	}

	return nil
}

func (s *guessService) Remove(id uint32) error {
	guesserIdMessage := rpc.GuesserId{Id: id}
	_, err := s.client.Remove(context.Background(), &guesserIdMessage)
	if err != nil {
		return err
	}

	return nil
}

func (s *guessService) Query(id uint32) (*rpc.QueryResponse, error) {
	guesserIdMessage := rpc.GuesserId{Id: id}
	res, err := s.client.Query(context.Background(), &guesserIdMessage)
	if err != nil {
		return nil, err
	}

	return res, nil
}
