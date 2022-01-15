package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
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
	gset := []Group{
		{
			ID: "a",
			Users: []Rankig{
				{"Hoge1", 100},
				{"Hoge2", 10},
				{"Hoge3", 50},
			},
		},
		{
			ID: "b",
			Users: []Rankig{
				{"Hoge101", 300},
				{"Hoge201", 200},
				{"Hoge301", 500},
			},
		},
	}

	sorter := NewRedisSorter()
	for _, g := range gset {
		sorter.Add(context.Background(), g.ID, g.Users)
	}

	// get top of group "a"
	topUser, err := sorter.GetTop(context.Background(), gset[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(topUser)

	// get "Hoge101"'s Rank
	targetUserRank, err := sorter.GetRank(context.Background(), gset[1].ID, "Hoge101")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(targetUserRank)
}

func NewRedisSorter() *RedisSorter {
	return &RedisSorter{redisClient: redis.NewClient(&redis.Options{
		Addr: ":6380",
	})}
}

func (r *RedisSorter) Add(ctx context.Context, key string, data []Rankig) error {
	members := make([]*redis.Z, 0, len(data))
	for _, d := range data {
		members = append(members, &redis.Z{Score: float64(d.Score), Member: d.User})
	}
	cmd := r.redisClient.ZAdd(ctx, key, members...)
	return cmd.Err()
}

func (r *RedisSorter) GetRank(ctx context.Context, key, member string) (int64, error) {
	cmd := r.redisClient.ZRevRank(ctx, key, member)
	if err := cmd.Err(); err != nil {
		return 0, err
	}
	return cmd.Val() + 1, nil
}

func (r *RedisSorter) GetTop(ctx context.Context, key string) (string, error) {
	cmd := r.redisClient.ZRevRange(ctx, key, 0, 0)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	fmt.Println(cmd.Val())
	return cmd.Val()[0], nil
}
