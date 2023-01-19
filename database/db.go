package database

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dint int) *redis.Client {

	redisDatabase := redis.NewClient(&redis.Options{
		DB:       dint,
		Addr:     "db:6543",
		Password: "",
	})

	return redisDatabase

}
