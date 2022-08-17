package service

import (
	"GoCleanMicroservice/abstract/domain/interactor"
	"golang.org/x/net/context"
)

var pingInteractor *interactor.PingInteractor

type PingServer struct {
	UnimplementedPingerServer
}

func InitPingInteractor(p *interactor.PingInteractor) {
	pingInteractor = p
}

func (s *PingServer) Ping(ctx context.Context, request *PingRequest) (*PingReply, error) {
	res, err := (*pingInteractor).Ping()
	return &PingReply{Message: res.Message}, err
}
