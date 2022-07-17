package main

import (
	"GoCleanMicroservice/Delivery"
	"GoCleanMicroservice/GinDelivery"
)

func main() {
	var server = createServer()
	err := server.Init()
	if err != nil {
		return
	}
	err = server.Serv()
	if err != nil {
		return
	}
}

func createServer() Delivery.Server {
	return &GinDelivery.Server{}
}
