package domain

import "GoCleanMicroservice/abstract/domain/interactor"

type InteractorPackage struct {
	Ping   *interactor.PingInteractor
	Health *interactor.HealthInteractor
}
