package redis

import (
	"github.com/redis/go-redis/v9"
	"social_network_for_programmers/internal/config"
)

func NewRedisClient(cfg *config.RS) *redis.Client {
	options := redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
	}

	return redis.NewClient(&options)
}
