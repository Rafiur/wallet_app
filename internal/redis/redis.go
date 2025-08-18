package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Client struct {
	RC *redis.Client
}

func NewClient(url string) (*Client, error) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	r := redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = r.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return &Client{
		RC: r,
	}, nil
}
