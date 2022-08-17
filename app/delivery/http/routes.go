package http

import (
	"GoCleanMicroservice/app/delivery/http/route"
)

func (m *Server) addRoutes() {
	pingInteractor := m.Interactors.Interactor
	if pingInteractor == nil {
		panic(any("Ping interactor cannot be null."))
	}
	m.Engine.GET("/ping", route.Ping)
	route.InitPingInteractor(m.Interactors.Interactor)

	healthInteractor := m.Interactors.Health
	if healthInteractor == nil {
		panic(any("Health interactor cannot be null."))
	}
	m.Engine.GET("/health", route.Health)
	route.InitHealthInteractor(m.Interactors.Health)
}
