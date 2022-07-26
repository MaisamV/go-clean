package GinDelivery

import (
	"GoCleanMicroservice/Domain"
	"GoCleanMicroservice/GinDelivery/Route"
)

type InteractorPackage struct {
	Interactor *Domain.PingInteractor
	Health     *Domain.HealthInteractor
}

func (m *Server) addRoutes() {
	m.Engine.GET("/ping", Route.Ping)
	m.Engine.GET("/health", Route.Health)
}

func (m *Server) setInteractors() {
	Route.InitPingInteractor(m.Interactors.Interactor)
	Route.InitHealthInteractor(m.Interactors.Health)
}
