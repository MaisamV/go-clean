package infrastructure

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DatabaseChecker implements the DatabaseChecker port
type DatabaseChecker struct {
	db *pgxpool.Pool
}

// NewDatabaseChecker creates a new database checker
func NewDatabaseChecker(db *pgxpool.Pool) *DatabaseChecker {
	return &DatabaseChecker{
		db: db,
	}
}

// CheckDatabase checks the database connectivity and response time
func (dc *DatabaseChecker) CheckDatabase(ctx context.Context) (bool, time.Duration, error) {
	start := time.Now()

	// Create a context with timeout for the health check
	checkCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Simple ping to check database connectivity
	err := dc.db.Ping(checkCtx)
	duration := time.Since(start)

	if err != nil {
		return false, duration, err
	}

	return true, duration, nil
}
