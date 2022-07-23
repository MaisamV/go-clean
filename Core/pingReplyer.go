package Core

import "GoCleanMicroservice/Domain"

type PingInteractor struct {
}

func (i *PingInteractor) Ping() (Domain.BaseResponse, error) {
	res := Domain.BaseResponse{Message: "PONG"}
	return res, nil
}
