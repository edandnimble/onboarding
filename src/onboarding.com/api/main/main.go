package main

import "onboarding.com/api"

func main() {
	go api.NewRpcServer()
	api.RunServer()
}
