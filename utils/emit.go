package utils

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient() (*redis.Client, func()) {
	rdb := redis.NewClient(&redis.Options{
		Username: GetEnv("REDIS_USERNAME", "default"),
		Addr:     GetEnv("REDIS_ADDRESS", "localhost"),
		Password: GetEnv("REDIS_PASSWORD", "secret"),
		DB:       0,
	})

	return rdb, func() {
		rdb.Close()
	}
}

func Listen(rdb *redis.Client, callback func(string)) {
	var ctx = context.Background()

	pubsub := rdb.Subscribe(ctx, "event")
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		callback(msg.Payload)
	}
}
