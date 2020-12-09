package redis

import "github.com/go-redis/redis/v8"

var RClient *redis.Client

func StartRedis() {
	RClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
