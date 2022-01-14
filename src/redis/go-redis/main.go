package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	delete(ctx, rdb)
}

func tutorial(ctx context.Context, rdb *redis.Client) {
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "hogehoge").Result()
	if err == redis.Nil {
		fmt.Println("hogehoge not found")
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hogehoge", val2)
}

func delete(ctx context.Context, rdb *redis.Client) {
	if err := rdb.Del(ctx, "key"); err != nil {
		log.Fatal(err)
	}
}
