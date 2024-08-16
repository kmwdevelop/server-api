package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server-api/config"
)

type MongoClient struct {
	cfg    *config.Config
	client *mongo.Client
	db     *mongo.Database

	like *mongo.Collection
}

func NewMongo(cfg *config.Config) (*MongoClient, error) {
	m := &MongoClient{cfg: cfg}

	var err error
	ctx := context.Background()

	if m.client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.URL)); err != nil {
		return nil, err
	} else if m.client.Ping(ctx, nil); err != nil {
		return nil, err
	} else {
		m.db = m.client.Database(cfg.Mongo.DB)
		m.like = m.db.Collection(cfg.Mongo.Like)
		log.Println("Connected to MongoDB")
		return m, nil
	}

	return m, nil
}

//gkmwdev
//6cHdTA80zuB1MrgZ
