package databases

import (
	"os"

	"github.com/go-redis/redis"
)

type Redis struct {
}

func NewRedis() *Redis {
	return &Redis{}
}

func (redises *Redis) Open() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return redisClient
}
