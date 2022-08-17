package domain

import "GoCleanMicroservice/abstract/domain/interactor"

type InteractorPackage struct {
	Interactor *interactor.PingInteractor
	Health     *interactor.HealthInteractor
}
