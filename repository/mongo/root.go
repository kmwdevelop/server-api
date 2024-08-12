package mongo

import "server-api/config"

type MongoClient struct {
	cfg *config.Config
}

func NewMongo(cfg *config.Config) (*MongoClient, error) {
	m := &MongoClient{cfg: cfg}

	return m, nil
}
