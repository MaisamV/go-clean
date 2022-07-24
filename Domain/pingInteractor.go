package Domain

import "GoCleanMicroservice/Domain/Model"

type PingInteractor interface {
	Ping() (Model.BaseResponse, error)
}
