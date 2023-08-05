package pgx

import (
	"GoCleanMicroservice/abstract/domain/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/maisamv/monitor"
)

type HealthRepo struct {
	Pool *pgxpool.Pool
}

func (r *HealthRepo) Check() model.HealthResponse {
	var err error
	var elapsedTime = monitor.MeasureTime(func() {
		var query, e = r.Pool.Query(context.Background(), "select 1;")
		defer query.Close()
		err = e
	})
	return model.HealthResponse{
		IsConnected: err == nil,
		Time:        elapsedTime.Milliseconds(),
	}
}
