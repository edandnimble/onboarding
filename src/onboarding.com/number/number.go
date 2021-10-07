package number

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	rpc "onboarding.com/number/grpcmodules"
	"onboarding.com/tasks"
	"onboarding.com/utils"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type rpcServer struct {
	rpc.UnimplementedNumberRpcServer
}

func (s *rpcServer) Add(ctx context.Context, req *rpc.Number) (*rpc.ResponseStatus, error) {
	mongoClient := utils.GetNumberClient()
	err := mongoClient.Add(req.GetNum())

	if err != nil {
		return &rpc.ResponseStatus{Ok: false, ErrCode: 1}, err
	}

	return &rpc.ResponseStatus{Ok: true, ErrCode: 0}, nil
}

func (s *rpcServer) Remove(ctx context.Context, req *rpc.Number) (*rpc.ResponseStatus, error) {
	mongoClient := utils.GetNumberClient()
	err := mongoClient.Remove(req.GetNum())
	if err != nil {
		return &rpc.ResponseStatus{Ok: false, ErrCode: 1}, err
	}

	return &rpc.ResponseStatus{Ok: true, ErrCode: 0}, nil
}

func (s *rpcServer) Query(ctx context.Context, req *rpc.Number) (*rpc.QueryResponse, error) {
	mongoClient := utils.GetNumberClient()
	resp, err := mongoClient.Query(req.GetNum())
	if err != nil {
		return &rpc.QueryResponse{Status: &rpc.ResponseStatus{Ok: false, ErrCode: 1}}, err
	}

	var guesses []*rpc.NumberInfo_Guesses
	for _, r := range resp.Found {
		guesses = append(guesses,
			&rpc.NumberInfo_Guesses{
				GuesserID: r.GuesserId,
				FoundAt:   timestamppb.New(r.FoundAt),
				Attempt:   r.Attempt})
	}

	return &rpc.QueryResponse{
		Status: &rpc.ResponseStatus{Ok: true, ErrCode: 0},
		Info:   &rpc.NumberInfo{Num: req.GetNum(), IsActive: resp.IsActive, Guesses: guesses}}, nil
}

func (s *rpcServer) IsExist(stream rpc.NumberRpc_IsExistServer) error {
	redisClient := utils.GetRedisClient()
	mongoClient := utils.GetNumberClient()
	for {
		guess, err := stream.Recv()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Printf("guesser %d num %d\n", guess.GetId(), guess.GetNum())

		attemptCounter, err := redisClient.IncreaseGuess(guess.GetId())
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		exists, err := mongoClient.IsExist(guess.GetNum())
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if exists {

			now := time.Now()
			fmt.Printf("guesser %d found %d at %v", guess.GetId(), guess.GetNum(), now)

			mongoNum := &utils.MongoFoundNumber{GuesserId: guess.GetId(), Attempt: attemptCounter, FoundAt: now}
			guessMongo := &utils.MongoFoundGuesser{Num: guess.GetNum(), Attempt: attemptCounter, FoundAt: now}
			err := utils.UpdateCorrectGuessTransaction(guess.GetNum(), mongoNum, guess.GetId(), guessMongo)
			if err != nil {
				fmt.Println("Mongo transaction error ", err.Error())
				return err
			}

			err = tasks.StartFindPrimeTask(guess.GetNum(), guess.GetId(), now)
			if err != nil {
				fmt.Println("async task error ", err.Error())
				return err

			}
		}

		stream.Send(&rpc.NumberExistResponse{Status: &rpc.ResponseStatus{Ok: true, ErrCode: 0}, Exist: exists})

	}
}

func NewRpcServer() {
	initAsyncTasksServer()

	port := os.Getenv("NUMBER_GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := grpc.NewServer()
	rpc.RegisterNumberRpcServer(s, &rpcServer{})
	s.Serve(lis)
}

func initAsyncTasksServer() {
	tasks.NewMachineryServer(false)
}
