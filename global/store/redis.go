package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"ring/global"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect Redis")
	}
	global.RedisClient = rdb
}
