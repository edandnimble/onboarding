package number

import (
	"context"
	"fmt"
	"net"
	rpc "onboarding.com/number/grpcmodules"
	"onboarding.com/utils"

	"google.golang.org/grpc"
)

type rpcServer struct{}

func (s *rpcServer) AddNumber(ctx context.Context, req *rpc.Number) (*rpc.ResponseStatus, error) {
	mongoClient := utils.GetNumberClient()
	err := mongoClient.Add(req.GetNum())

	//redisClient := utils.GetRedisClient()
	//err := redisClient.AddNumber(req.GetNum())
	if err != nil {
		return &rpc.ResponseStatus{Ok: false, ErrCode: 1}, err
	}

	return &rpc.ResponseStatus{Ok: true, ErrCode: 0}, nil
}

func (s *rpcServer) RemoveNumber(ctx context.Context, req *rpc.Number) (*rpc.ResponseStatus, error) {
	mongoClient := utils.GetNumberClient()
	err := mongoClient.Remove(req.GetNum())
	if err != nil {
		return &rpc.ResponseStatus{Ok: false, ErrCode: 1}, err
	}

	return &rpc.ResponseStatus{Ok: true, ErrCode: 0}, nil
}

func (s *rpcServer) QueryNumber(ctx context.Context, req *rpc.Number) (*rpc.QueryResponse, error) {
	// use mongo
	redisClient := utils.GetRedisClient()
	// TODO: get number details (and save)
	err := redisClient.RemoveNumber(req.GetNum())
	if err != nil {
		return &rpc.QueryResponse{Status: &rpc.ResponseStatus{Ok: false, ErrCode: 1}}, err
	}

	return &rpc.QueryResponse{
		Status: &rpc.ResponseStatus{Ok: true, ErrCode: 0},
		Info: &rpc.NumberInfo{Num: , IsActive: , GuesserID: , Found:},}, nil
}

func (s *NumberServer) IsNumberExist(stream rpc.NumberManager_IsNumberExistServer) error {
	redisClient := utils.GetRedisClient()
	for {
		guess, err := stream.Recv()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		err := redisClient.IncreaseGuess(guess.GetId())
		exists, err := redisClient.IsNumberExist(guess.GetNum())
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		stream.Send(&rpc.NumberExistResponse{Status: &rpc.ResponseStatus{Ok: true, ErrCode: 0}, Exist: exists})

		if exists {
			// TODO run machinary & update monog guesser, num in transaction
		}
	}
}

func NewRpcServer() {
	lis, err := net.Listen("tcp", "localhost:1001")
	if err != nil {
		fmt.Println(err.Error())
	}

	s := grpc.NewServer()
	rpc.RegisterNumberManagerServer(s, &NumberRpcServer{})
	s.Serve(lis)
}
