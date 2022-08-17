package interactor

import (
	"GoCleanMicroservice/abstract/domain/model"
)

type PingInteractor struct {
}

func (i *PingInteractor) Ping() (model.BaseResponse, error) {
	res := model.BaseResponse{Message: "PONG"}
	return res, nil
}
