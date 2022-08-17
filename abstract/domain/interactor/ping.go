package interactor

import (
	"GoCleanMicroservice/abstract/domain/model"
)

type PingInteractor interface {
	Ping() (model.BaseResponse, error)
}
