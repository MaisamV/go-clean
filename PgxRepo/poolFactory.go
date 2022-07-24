package PgxRepo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type PoolFactory struct {
}

func (p *PoolFactory) Create() (*pgxpool.Pool, error) {
	return pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}
