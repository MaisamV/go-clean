package grpc

import (
	"GoCleanMicroservice/app/delivery/grpc/service"
)

func (m *Server) addServices() {
	m.addPing()
	m.addHealth()
}

func (m *Server) addHealth() {
	healthInteractor := m.Interactors.Health
	panicOnNull("Health", healthInteractor)
	service.InitHealthInteractor(m.Interactors.Health)
	service.RegisterHealthCheckServer(m.Engine, &service.HealthServer{})
}

func (m *Server) addPing() {
	pingInteractor := m.Interactors.Ping
	panicOnNull("Ping", pingInteractor)
	service.InitPingInteractor(pingInteractor)
	service.RegisterPingerServer(m.Engine, &service.PingServer{})
}

func panicOnNull(interactorName string, interactor interface{}) {
	if interactor == nil {
		panic(any(interactorName + " interactor cannot be null."))
	}
}
