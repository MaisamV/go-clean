package grpc

import (
	"GoCleanMicroservice/app/delivery/grpc/service"
)

func (m *Server) addServices() {
	//ping
	pingInteractor := m.Interactors.Ping
	if pingInteractor == nil {
		panic(any("Ping interactor cannot be null."))
	}
	service.InitPingInteractor(pingInteractor)
	service.RegisterPingerServer(m.Engine, &service.PingServer{})
	//health
	healthInteractor := m.Interactors.Health
	if healthInteractor == nil {
		panic(any("Health interactor cannot be null."))
	}
	service.InitHealthInteractor(m.Interactors.Health)
	service.RegisterHealthCheckServer(m.Engine, &service.HealthServer{})
}
