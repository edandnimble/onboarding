package service

import (
	"context"
	"fmt"

	rpc "onboarding.com/number/grpcmodules"

	"google.golang.org/grpc"
)

type numService struct {
	client rpc.NumberRpcClient
}

func NewNumService() (*numService, error) {
	conn, err := grpc.Dial("localhost:1001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	client := rpc.NewNumberRpcClient(conn)
	return &numService{client: client}, nil
}

func (s *numService) Add(num uint32) error {
	numMessage := rpc.Number{Num: num}
	_, err := s.client.AddNumber(context.Background(), &numMessage)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *numService) Remove(num uint32) error {
	numMessage := rpc.Number{Num: num}
	_, err := s.client.RemoveNumber(context.Background(), &numMessage)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
