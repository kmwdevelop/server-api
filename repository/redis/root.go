package redis

import (
	"github.com/go-redis/redis/v7"
	"log"
	"server-api/config"
)

type RedisClient struct {
	cfg    *config.Config
	client *redis.Client
}

func NewRedis(cfg *config.Config) (*RedisClient, error) {
	r := &RedisClient{cfg: cfg}
	r.client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.URL,
		Password: cfg.Redis.Password,
		Username: cfg.Redis.UserName,
		DB:       cfg.Redis.DB,
	})

	if err := r.client.Ping().Err(); err != nil {
		return nil, err
	} else {
		log.Println("Redis connect success")
	}
	return r, nil
}
