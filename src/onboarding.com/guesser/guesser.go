package guesser

import (
	"context"
	"fmt"

	apirpc "onboarding.com/api/grpcmodules"
	rpc "onboarding.com/guesser/grpcmodules"
	numrpc "onboarding.com/number/grpcmodules"

	"google.golang.org/grpc"
)

type rpcServer struct {
	conn     *grpc.ClientConn
	client   rpc.GuesserManagerClient
	idToChan map[uint32](chan bool)
	id       uint32
}

func guesserRoutine(begin, incrementBy, sleepInterval, id uint32, conn *grpc.ClientConn, done <-chan bool) {
	i := 0
	client := apirpc.NewApiManagerClient(conn)
	for {
		select {
		case <-done:
			return
		default:
			client.GuessNumber(numrpc.GuessNumber{Num: begin + (incrementBy * uint32(i)), Id: id})
			i += 1
		}
	}
}

func (s *rpcServer) AddGuesser(ctx context.Context, req *rpc.Guesser) (*rpc.AddGuesserResponse, error) {
	done := make(chan bool)
	s.id += 1
	s.idToChan[s.id] = done

	go guesserRoutine(req.GetBeginAt(), req.GetIncrementBy(), req.GetSleepInterval(), s.id, s.conn, done)

	return &rpc.AddGuesserResponse{Status: &rpc.ResponseStatus{Ok: true, ErrCode: 0}, Id: s.id}, nil
}

func (s *rpcServer) RemoveGuesser(ctx context.Context, req *rpc.GuesserId) (*rpc.ResponseStatus, error) {
	done, ok := s.idToChan[req.GetId()]
	if !ok {
		return &rpc.ResponseStatus{Ok: false, ErrCode: 404}, nil
	}
	done <- true
	close(done)
	delete(s.idToChan, req.GetId())

	return &rpc.ResponseStatus{Ok: true, ErrCode: 0}, nil
}

func (s *rpcServer) QueryGuesser(ctx context.Context, req *rpc.GuesserId) (*rpc.QueryResponse, error) {
	// redisClient := utils.GetRedisClient()
	// err := redisClient.AddNumber(req.GetNum())
	// if err != nil {

	// 	return &rpc.QueryResponse{Status: &rpc.ResponseStatus{Ok: false, ErrCode: 1}, Active: 0, GuessInfo: }, err
	// }

	return &rpc.QueryResponse{Status: &rpc.ResponseStatus{Ok: true, ErrCode: 0}, Active: false, Info: nil}, nil
}

// func (s *guessServer) IsExists(id uint32) (bool) {
// 	_, ok := s.idToChan[id]
// 	return ok
// }

func NewRpcServer() (*rpcServer, error) {
	// api client
	conn, err := grpc.Dial("localhost:1002", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	client := apirpc.NewApiManagerClient(conn)
	idToChan := make(map[uint32](chan bool))
	return &rpcServer{client: client, conn: conn, idToChan: idToChan}, nil
}
