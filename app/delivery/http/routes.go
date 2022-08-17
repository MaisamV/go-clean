package http

import (
	"GoCleanMicroservice/app/delivery/http/route"
)

func (m *Server) addRoutes() {
	m.Engine.GET("/ping", route.Ping)
	route.InitPingInteractor(m.Interactors.Interactor)

	m.Engine.GET("/health", route.Health)
	route.InitHealthInteractor(m.Interactors.Health)
}
