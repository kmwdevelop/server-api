package repository

import (
	"server-api/config"
	"server-api/repository/mongo"
	"server-api/repository/redis"
)

type Repository struct {
	cfg   *config.Config
	mongo *mongo.MongoClient
	redis *redis.RedisClient
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{cfg: cfg}
	var err error
	if r.mongo, err = mongo.NewMongo(cfg); err != nil {
		return nil, err
	} else if r.redis, err = redis.NewRedis(cfg); err != nil {
		return nil, err
	}

	return r, nil
}
