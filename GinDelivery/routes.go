package GinDelivery

import (
	"GoCleanMicroservice/Domain"
	"GoCleanMicroservice/GinDelivery/Route"
)

type InteractorPackage struct {
	Interactor *Domain.PingInteractor
}

func (m *Server) addRoutes() {
	m.Engine.GET("/ping", Route.Ping)
}

func (m *Server) setInteractors() {
	Route.InitPingInteractor(m.Interactors.Interactor)
}
