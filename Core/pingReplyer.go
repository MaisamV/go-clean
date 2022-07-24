package Core

import (
	"GoCleanMicroservice/Domain/Model"
)

type PingInteractor struct {
}

func (i *PingInteractor) Ping() (Model.BaseResponse, error) {
	res := Model.BaseResponse{Message: "PONG"}
	return res, nil
}
