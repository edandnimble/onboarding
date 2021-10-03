package tasks

import (
	"fmt"
	"sync"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

var server *machinery.Server
var machineryOnce sync.Once

func NewMachineryServer(startWorker bool) {
	machineryOnce.Do(func() {
		var err error
		server, err = getMachineryServer()
		if err != nil {
			fmt.Println("error creating machinery server: ", err.Error())
			return
		}
		if startWorker {
			err = launchWorker(server)
			if err != nil {
				fmt.Println("error launching machinery worker: ", err.Error())
				server = nil
			}
		}
	})
}

func StartFindPrimeTask(num, guesserId uint32, foundAt time.Time) error {
	task := tasks.Signature{
		Name: "prime",
		Args: []tasks.Arg{
			{Type: "uint32", Value: num},
			{Type: "uint32", Value: guesserId},
			{Type: "int64", Value: foundAt.Unix()},
		},
	}
	_, err := server.SendTask(&task)
	return err
}

func launchWorker(server *machinery.Server) error {
	worker := server.NewWorker("prime_worker", 10)
	return worker.Launch()
}

func getMachineryServer() (*machinery.Server, error) {
	taskserver, err := machinery.NewServer(&config.Config{
		Broker:        "redis://:6379",
		ResultBackend: "redis://:6379",
	})
	if err != nil {
		return nil, err
	}

	taskserver.RegisterTasks(map[string]interface{}{
		"prime": FindPrime,
	})

	return taskserver, nil
}
