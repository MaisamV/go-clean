package http

import (
	"GoCleanMicroservice/app/delivery/http/route"
)

func (m *Server) addRoutes() {
	m.addPing()
	m.addHealth()
}

func (m *Server) addHealth() {
	healthInteractor := m.Interactors.Health
	panicOnNull("Health", healthInteractor)
	m.Engine.GET("/health", route.Health)
	route.InitHealthInteractor(healthInteractor)
}

func (m *Server) addPing() {
	pingInteractor := m.Interactors.Ping
	panicOnNull("Ping", pingInteractor)
	m.Engine.GET("/ping", route.Ping)
	route.InitPingInteractor(pingInteractor)
}

func panicOnNull(interactorName string, interactor interface{}) {
	if interactor == nil {
		panic(any(interactorName + " interactor cannot be null."))
	}
}
