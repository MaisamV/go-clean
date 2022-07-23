package main

import (
	"GoCleanMicroservice/Core"
	"GoCleanMicroservice/Delivery"
	"GoCleanMicroservice/Domain"
	"GoCleanMicroservice/GinDelivery"
)

func main() {
	var pingInteractor = createInteractor()
	var server = createServer(&pingInteractor)
	err := server.Init()
	if err != nil {
		return
	}
	err = server.Serv()
	if err != nil {
		return
	}
}

func createInteractor() Domain.PingInteractor {
	return &Core.PingInteractor{}
}

func createServer(i *Domain.PingInteractor) Delivery.Server {
	return &GinDelivery.Server{Interactors: &GinDelivery.InteractorPackage{Interactor: i}}
}
