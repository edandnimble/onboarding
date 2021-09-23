package service

import (
	"context"
	"fmt"

	rpc "onboarding.com/guesser/grpcmodules"

	"google.golang.org/grpc"
)

type guessService struct {
	client rpc.GuesserRpcClient
}

func NewGuessService() (*guessService, error) {
	conn, err := grpc.Dial("localhost:1002", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	client := rpc.NewGuesserRpcClient(conn)
	return &guessService{client: client}, nil
}

func (s *guessService) Add(beginAt, incrementBy, sleepInterval uint32) error {
	gusserMessage := rpc.Guesser{
		BeginAt:       beginAt,
		IncrementBy:   incrementBy,
		SleepInterval: sleepInterval}
	_, err := s.client.AddGuesser(context.Background(), &gusserMessage)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *guessService) Remove(id uint32) error {
	guesserIdMessage := rpc.GuesserId{Id: id}
	_, err := s.client.RemoveGuesser(context.Background(), &guesserIdMessage)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
