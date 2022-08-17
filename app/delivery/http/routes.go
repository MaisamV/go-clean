package http

import (
	"GoCleanMicroservice/app/delivery/http/route"
)

func (m *Server) addRoutes() {
	m.Engine.GET("/ping", route.Ping)
	m.Engine.GET("/health", route.Health)
}

func (m *Server) setInteractors() {
	route.InitPingInteractor(m.Interactors.Interactor)
	route.InitHealthInteractor(m.Interactors.Health)
}
