package PgxRepo

import (
	"GoCleanMicroservice/Common"
	"GoCleanMicroservice/Domain/Model"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type HealthRepo struct {
	pool *pgxpool.Pool
}

func (r *HealthRepo) Check() Model.HealthResponse {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	var query pgx.Rows
	var err error
	var elapsedTime = Common.MeasureTime(func() {
		query, err = r.pool.Query(context.Background(), "select 1")
	})
	return Model.HealthResponse{
		IsConnected: err == nil,
		Time:        elapsedTime.Milliseconds(),
	}
}
