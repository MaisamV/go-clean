package pgx

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type PoolFactory struct {
}

func (p *PoolFactory) Create() (*pgxpool.Pool, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name?pool_min_conns=10&pool_max_conns=100"
	return pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}
