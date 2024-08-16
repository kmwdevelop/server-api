package repository

import (
	"server-api/config"
	"server-api/repository/mongo"
	"server-api/repository/redis"
)

type Repository struct {
	cfg   *config.Config
	Mongo *mongo.MongoClient
	Redis *redis.RedisClient
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{cfg: cfg}
	var err error
	if r.Mongo, err = mongo.NewMongo(cfg); err != nil {
		return nil, err
	} else if r.Redis, err = redis.NewRedis(cfg); err != nil {
		return nil, err
	}

	return r, nil
}
