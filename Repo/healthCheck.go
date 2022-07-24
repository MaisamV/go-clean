package Repo

import "GoCleanMicroservice/Domain/Model"

type HealthRepo interface {
	Check() Model.HealthResponse
}
