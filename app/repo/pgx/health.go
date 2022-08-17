package pgx

import (
	"GoCleanMicroservice/abstract/domain/model"
	"GoCleanMicroservice/common"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type HealthRepo struct {
	Pool *pgxpool.Pool
}

func (r *HealthRepo) Check() model.HealthResponse {
	var err error
	var elapsedTime = common.MeasureTime(func() {
		var query, e = r.Pool.Query(context.Background(), "select 1;")
		err = e
		for query.Next() {
		}
	})
	return model.HealthResponse{
		IsConnected: err == nil,
		Time:        elapsedTime.Milliseconds(),
	}
}
