package grpc

import (
	"GoCleanMicroservice/app/delivery/grpc/service"
)

func (m *Server) addServices() {
	//ping
	service.InitPingInteractor(m.Interactors.Interactor)
	service.RegisterPingerServer(m.Engine, &service.PingServer{})
	//health
	service.InitHealthInteractor(m.Interactors.Health)
	service.RegisterHealthCheckServer(m.Engine, &service.HealthServer{})
}
