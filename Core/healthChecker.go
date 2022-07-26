package Core

import (
	"GoCleanMicroservice/Domain/Model"
	"GoCleanMicroservice/Repo"
)

type HealthInteractor struct {
	HealthRepo *Repo.HealthRepo
}

func (i *HealthInteractor) Health() Model.HealthResponse {
	return (*(i.HealthRepo)).Check()
}
