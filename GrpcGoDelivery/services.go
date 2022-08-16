package GrpcGoDelivery

import (
	"GoCleanMicroservice/GrpcGoDelivery/service"
	pb "GoCleanMicroservice/GrpcGoDelivery/service"
)

func (m *Server) setInteractors() {
	service.InitPingInteractor(m.Interactors.Interactor)
	service.InitHealthInteractor(m.Interactors.Health)
}

func (m *Server) addServices() {
	pb.RegisterPingerServer(m.Engine, &service.PingServer{})
	pb.RegisterHealthCheckServer(m.Engine, &service.HealthServer{})
}
