package service

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	rpc "onboarding.com/number/grpcmodules"
	"onboarding.com/utils"
)

type numService struct {
	client rpc.NumberRpcClient
}

func NewNumService() (*numService, error) {
	ip, port, err := utils.GetServiceDNS("number")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("connecting to number: " + ip + ":" + port)
	conn, _ := grpc.Dial(ip+":"+port, grpc.WithInsecure())

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
