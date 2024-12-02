package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"ring/global"
)

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect Redis")
	}
	global.RedisClient = rdb
}
