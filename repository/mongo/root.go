package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server-api/config"
	"time"
)

type MongoClient struct {
	cfg    *config.Config
	client *mongo.Client
	db     *mongo.Database

	stock *mongo.Collection
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
		m.stock = m.db.Collection(cfg.Mongo.Stock)
		log.Println("Connected to MongoDB")
		return m, nil
	}

	return m, nil
}

func (m *MongoClient) AddStock(name string) error {

	opt := options.Update().SetUpsert(true)
	filter := bson.M{"name": name}
	update := bson.M{"$set": bson.M{
		"name":     name,
		"createAt": time.Now().Unix(),
	}}

	if _, err := m.stock.UpdateOne(context.Background(), filter, update, opt); err != nil {
		return err
	} else {
		return nil
	}
}
