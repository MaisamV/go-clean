package service

import (
	"GoCleanMicroservice/Domain"
	"golang.org/x/net/context"
)

var healthInteractor *Domain.HealthInteractor

type HealthServer struct {
	UnimplementedHealthCheckServer
}

func InitHealthInteractor(h *Domain.HealthInteractor) {
	if healthInteractor == nil {
		healthInteractor = h
	}
}

func (s *HealthServer) Health(ctx context.Context, request *HealthRequest) (*HealthReply, error) {
	res := (*healthInteractor).Health()
	return &HealthReply{IsConnected: res.IsConnected, Time: res.Time}, nil
}
