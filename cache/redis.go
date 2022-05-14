package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

func connectRedis(ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
	return client, err
}

// Set lifetime = 0 means forever
func Set(key, val string, lifetime int) {
	var ctx = context.Background()
	redisClient, err := connectRedis(ctx)
	expire := time.Duration(lifetime) * time.Minute
	redisClient.Set(ctx, key, val, expire).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func Get(key string) string {
	var ctx = context.Background()
	redisClient, err := connectRedis(ctx)
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val
}
