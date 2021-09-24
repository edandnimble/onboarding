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
	conn, err := grpc.Dial(":50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	client := rpc.NewNumberRpcClient(conn)
	return &numService{client: client}, nil
}

func (s *numService) Add(num uint32) error {
	numMessage := rpc.Number{Num: num}
	_, err := s.client.Add(context.Background(), &numMessage)
	if err != nil {
		return err
	}

	return nil
}

func (s *numService) Remove(num uint32) error {
	numMessage := rpc.Number{Num: num}
	_, err := s.client.Remove(context.Background(), &numMessage)
	if err != nil {
		return err
	}

	return nil
}

func (s *numService) Query(num uint32) (*rpc.QueryResponse, error) {
	numberMessage := rpc.Number{Num: num}
	res, err := s.client.Query(context.Background(), &numberMessage)
	if err != nil {
		return nil, err
	}

	return res, nil
}
