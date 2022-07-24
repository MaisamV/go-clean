package PgxRepo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type HealthRepo struct {
	pool *pgxpool.Pool
}

func (r *HealthRepo) Check() bool {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	_, err := r.pool.Query(context.Background(), "select 1")
	return err == nil
}
