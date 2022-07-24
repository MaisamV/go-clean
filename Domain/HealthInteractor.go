package Domain

import "GoCleanMicroservice/Domain/Model"

type HealthInteractor interface {
	Health() Model.HealthResponse
}
