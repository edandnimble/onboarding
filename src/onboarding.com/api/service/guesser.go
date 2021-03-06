package service

import (
	"context"
	"fmt"

	rpc "onboarding.com/guesser/grpcmodules"
	"onboarding.com/utils"

	"google.golang.org/grpc"
)

type GuessService struct {
	client rpc.GuesserRpcClient
}

func NewGuessService() (*GuessService, error) {
	ip, port, err := utils.GetServiceDNS("guesser")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println("connecting to guesser: " + ip + ":" + port)
	conn, err := grpc.Dial(ip+":"+port, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	rpcClient := rpc.NewGuesserRpcClient(conn)
	return &GuessService{client: rpcClient}, nil
}

func (s *GuessService) Add(beginAt, incrementBy, sleepInterval uint32) error {
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

func (s *GuessService) Remove(id uint32) error {
	guesserIdMessage := rpc.GuesserId{Id: id}
	_, err := s.client.Remove(context.Background(), &guesserIdMessage)
	if err != nil {
		return err
	}

	return nil
}

func (s *GuessService) Query(id uint32) (*rpc.QueryResponse, error) {
	guesserIdMessage := rpc.GuesserId{Id: id}
	res, err := s.client.Query(context.Background(), &guesserIdMessage)
	if err != nil {
		return nil, err
	}

	return res, nil
}
