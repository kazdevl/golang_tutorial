package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/go.net/context"
)

type Group struct {
	ID    string
	Users []Rankig
}

type Rankig struct {
	User  string
	Score int64
}

type RedisSorter struct {
	redisClient *redis.Client
}

func main() {

}

func NewRedisSorter() *RedisSorter {
	return &RedisSorter{redisClient: redis.NewClient(&redis.Options{
		Addr: "http://localhost:6380",
	})}
}

func (r *RedisSorter) Add(ctx context.Context, key string, data Rankig) error {
	cmd := r.redisClient.ZAdd(ctx, key, &redis.Z{Score: float64(data.Score), Member: data.User})
	return cmd.Err()
}

func (r *RedisSorter) GetRank(ctx context.Context, key, id string) (int64, error) {
	cmd := r.redisClient.ZRevRank(ctx, key, id)
	return cmd.Val() + 1, cmd.Err()
}

func (r *RedisSorter) GetTop(ctx context.Context, key, id string) (string, error) {
	cmd := r.redisClient.ZRevRange(ctx, key, 0, 1)
	fmt.Println(cmd.Val())
	return cmd.Val()[0], cmd.Err()
}

func (r *RedisSorter) GetHighTier(ctx context.Context, key, id string) (int64, error) {

}

func (r *RedisSorter) GetRankList(ctx context.Context, key, id string) (int64, error) {

}
