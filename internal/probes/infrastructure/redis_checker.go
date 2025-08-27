package infrastructure

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisChecker implements the RedisChecker port
type RedisChecker struct {
	client *redis.Client
}

// NewRedisChecker creates a new Redis checker
func NewRedisChecker(client *redis.Client) *RedisChecker {
	return &RedisChecker{
		client: client,
	}
}

// CheckRedis checks the Redis connectivity and response time
func (rc *RedisChecker) CheckRedis(ctx context.Context) (bool, time.Duration, error) {
	start := time.Now()
	
	// Create a context with timeout for the health check
	checkCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	// Simple ping to check Redis connectivity
	result := rc.client.Ping(checkCtx)
	duration := time.Since(start)
	
	if err := result.Err(); err != nil {
		return false, duration, err
	}
	
	return true, duration, nil
}