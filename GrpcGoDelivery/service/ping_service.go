package service

import (
	"GoCleanMicroservice/Domain"
	"golang.org/x/net/context"
)

var pingInteractor *Domain.PingInteractor

type PingServer struct {
	UnimplementedPingerServer
}

func InitPingInteractor(p *Domain.PingInteractor) {
	if pingInteractor == nil {
		pingInteractor = p
	}
}

func (s *PingServer) Ping(ctx context.Context, request *PingRequest) (*PingReply, error) {
	res, err := (*pingInteractor).Ping()
	return &PingReply{Message: res.Message}, err
}
