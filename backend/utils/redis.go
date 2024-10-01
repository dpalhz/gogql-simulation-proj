package utils

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis(addr string, password string) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	// Test the connection
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	// Reset database
	// err = rdb.FlushDB(ctx).Err()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to reset Redis database: %v", err)
	// }

	return rdb, nil
}
