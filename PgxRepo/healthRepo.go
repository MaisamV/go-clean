package PgxRepo

import (
	"GoCleanMicroservice/Common"
	"GoCleanMicroservice/Domain/Model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type HealthRepo struct {
	Pool *pgxpool.Pool
}

func (r *HealthRepo) Check() Model.HealthResponse {
	var err error
	var elapsedTime = Common.MeasureTime(func() {
		var query, e = r.Pool.Query(context.Background(), "select 1;")
		err = e
		for query.Next() {
		}
	})
	return Model.HealthResponse{
		IsConnected: err == nil,
		Time:        elapsedTime.Milliseconds(),
	}
}
