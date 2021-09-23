package main

import (
	"onboarding.com/api"
	"onboarding.com/guesser"
	"onboarding.com/number"
)

func main() {
	go guesser.NewRpcServer()
	go number.NewRpcServer()
	go api.NewRpcServer()
	api.RunServer()
}
