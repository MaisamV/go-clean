package interactor

import (
	"GoCleanMicroservice/abstract/domain/model"
	"GoCleanMicroservice/abstract/repo"
)

type HealthInteractor struct {
	HealthRepo *repo.HealthRepo
}

func (i *HealthInteractor) Health() model.HealthResponse {
	return (*(i.HealthRepo)).Check()
}
