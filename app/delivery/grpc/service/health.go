package service

import (
	"GoCleanMicroservice/abstract/domain/interactor"
	"golang.org/x/net/context"
)

var healthInteractor *interactor.HealthInteractor

type HealthServer struct {
	UnimplementedHealthCheckServer
}

func InitHealthInteractor(h *interactor.HealthInteractor) {
	if healthInteractor == nil {
		healthInteractor = h
	}
}

func (s *HealthServer) Health(ctx context.Context, request *HealthRequest) (*HealthReply, error) {
	res := (*healthInteractor).Health()
	return &HealthReply{IsConnected: res.IsConnected, Time: res.Time}, nil
}
