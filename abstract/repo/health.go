package repo

import (
	"GoCleanMicroservice/abstract/domain/model"
)

type HealthRepo interface {
	Check() model.HealthResponse
}
