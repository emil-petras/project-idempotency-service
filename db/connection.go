package db

import (
	"github.com/go-redis/redis"
)

var Client *redis.Client

func Connect(addr string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	Client = client

	return nil
}
