package guesser

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	apirpc "onboarding.com/api/grpcmodules"
	rpc "onboarding.com/guesser/grpcmodules"
	"onboarding.com/utils"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type rpcServer struct {
	rpc.UnimplementedGuesserRpcServer
	conn     *grpc.ClientConn
	idToChan map[uint32](chan bool)
	id       uint32
}

func guesserRoutine(begin, incrementBy, sleepInterval, id uint32, conn *grpc.ClientConn, done <-chan bool) {
	var i uint32 = 0
	client := apirpc.NewApiRpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.GuessNumber(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		select {
		case <-done:
			return
		default:
			nextNum := begin + (incrementBy * i)
			//fmt.Printf("Guesser: %d guessed number: %d\n", id, nextNum)
			stream.Send(&apirpc.Guess{Num: nextNum, Id: id})
			i += 1
			time.Sleep(time.Duration(sleepInterval) * time.Millisecond)
		}
	}
}

func (s *rpcServer) Add(ctx context.Context, req *rpc.Guesser) (*rpc.AddGuesserResponse, error) {
	mongoClient := utils.GetGuessClient()
	newGuesser := &utils.MongoGuesser{
		Id:            s.id + 1,
		IsActive:      true,
		BeginAt:       req.GetBeginAt(),
		IncrementBy:   req.GetIncrementBy(),
		SleepInterval: req.GetSleepInterval()}
	fmt.Printf("Adding new guesser %v\n", newGuesser)

	err := mongoClient.Add(newGuesser)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	done := make(chan bool)
	s.id += 1
	s.idToChan[s.id] = done

	go guesserRoutine(req.GetBeginAt(), req.GetIncrementBy(), req.GetSleepInterval(), s.id, s.conn, done)

	return &rpc.AddGuesserResponse{Status: &rpc.ResponseStatus{Ok: true, ErrCode: 0}, Id: s.id}, nil
}

func (s *rpcServer) Remove(ctx context.Context, req *rpc.GuesserId) (*rpc.ResponseStatus, error) {
	mongoClient := utils.GetGuessClient()
	err := mongoClient.Remove(req.GetId())
	if err != nil {
		return &rpc.ResponseStatus{Ok: false, ErrCode: 500}, nil
	}
	done, ok := s.idToChan[req.GetId()]
	if !ok {
		return &rpc.ResponseStatus{Ok: false, ErrCode: 404}, nil
	}
	done <- true
	close(done)
	delete(s.idToChan, req.GetId())

	return &rpc.ResponseStatus{Ok: true, ErrCode: 0}, nil
}

func (s *rpcServer) Query(ctx context.Context, req *rpc.GuesserId) (*rpc.QueryResponse, error) {
	mongoClient := utils.GetGuessClient()
	res, err := mongoClient.Query(req.GetId())
	if err != nil {
		fmt.Printf("Error query mongo: %s\n", err.Error())
		return &rpc.QueryResponse{Status: &rpc.ResponseStatus{Ok: false, ErrCode: 500}}, nil
	}
	if res == nil {
		return &rpc.QueryResponse{Status: &rpc.ResponseStatus{Ok: false, ErrCode: 404}}, nil
	}

	var guesses []*rpc.GuessInfo
	for _, g := range res.Found {
		guesses = append(guesses, &rpc.GuessInfo{Num: g.Num, Attempt: g.Attempt, FoundAt: timestamppb.New(g.FoundAt)})
	}

	return &rpc.QueryResponse{
		Status:        &rpc.ResponseStatus{Ok: true, ErrCode: 0},
		Active:        res.IsActive,
		BeginAt:       res.BeginAt,
		IncrementBy:   res.IncrementBy,
		SleepInterval: res.SleepInterval,
		Guesses:       guesses}, nil
}

func NewRpcServer() {
	// api client
	ip, port, err := utils.GetServiceDNS("grpcapi")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("connecting to grpc api: " + ip + ":" + port)

	conn, err := grpc.Dial(ip+":"+port, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	// guess grpc server
	idToChan := make(map[uint32](chan bool))
	guessRpcServer := rpcServer{conn: conn, idToChan: idToChan}

	grpcPort := os.Getenv("GUESSER_GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := grpc.NewServer()
	rpc.RegisterGuesserRpcServer(s, &guessRpcServer)
	s.Serve(lis)
}
