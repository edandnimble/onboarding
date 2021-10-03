package main

import (
	"onboarding.com/api"
	"onboarding.com/guesser"
	"onboarding.com/number"
	"onboarding.com/tasks"
)

func main() {
	go guesser.NewRpcServer()
	go number.NewRpcServer()
	go api.NewRpcServer()
	go tasks.NewMachineryServer(true)
	api.RunServer()
}
