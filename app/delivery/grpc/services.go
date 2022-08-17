package grpc

import (
	service2 "GoCleanMicroservice/app/delivery/grpc/service"
)

func (m *Server) setInteractors() {
	service2.InitPingInteractor(m.Interactors.Interactor)
	service2.InitHealthInteractor(m.Interactors.Health)
}

func (m *Server) addServices() {
	service2.RegisterPingerServer(m.Engine, &service2.PingServer{})
	service2.RegisterHealthCheckServer(m.Engine, &service2.HealthServer{})
}
