package interactor

import (
	"GoCleanMicroservice/abstract/domain/model"
)

type HealthInteractor interface {
	Health() model.HealthResponse
}
