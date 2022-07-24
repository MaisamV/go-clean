package Core

import (
	"GoCleanMicroservice/Domain/Model"
	"GoCleanMicroservice/Repo"
)

type HealthInteractor struct {
	repo *Repo.HealthRepo
}

func (i *HealthInteractor) Health() Model.HealthResponse {
	return (*i.repo).Check()
}
