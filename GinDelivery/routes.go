package GinDelivery

import (
	"GoCleanMicroservice/GinDelivery/Route"
)

func (m *Server) addRoutes() {
	m.Engine.GET("/ping", Route.Ping)
	m.Engine.GET("/health", Route.Health)
}

func (m *Server) setInteractors() {
	Route.InitPingInteractor(m.Interactors.Interactor)
	Route.InitHealthInteractor(m.Interactors.Health)
}
