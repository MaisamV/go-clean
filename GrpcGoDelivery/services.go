package GrpcGoDelivery

import (
	"GoCleanMicroservice/GrpcGoDelivery/service"
	pb "GoCleanMicroservice/GrpcGoDelivery/service"
)

func (m *Server) setInteractors() {
	service.InitPingInteractor(m.Interactors.Interactor)
}

func (m *Server) addServices() {
	pb.RegisterPingerServer(m.Engine, &service.PingServer{})
}
