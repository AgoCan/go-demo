package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClient *redis.Client

func connect() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	r, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

}

func main() {
	connect()
	err := redisClient.Set(ctx, "key1", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := redisClient.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
