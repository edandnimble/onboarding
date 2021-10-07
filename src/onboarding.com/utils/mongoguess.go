package utils

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoGuessClient struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func GetGuessClient() *mongoGuessClient {
	client, db := GetMongoClientAndDb()
	collection := db.Collection("guess")
	collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{primitive.E{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true)},
	)
	return &mongoGuessClient{
		client:     client,
		db:         db,
		collection: collection}
}

func (c *mongoGuessClient) Add(guesser *MongoGuesser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.collection.InsertOne(ctx, guesser)
	return err
}

func (c *mongoGuessClient) Remove(id uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := c.Query(id)
	if err != nil {
		return err
	}

	if data.IsActive == false {
		return fmt.Errorf("Guesser %v is not active", id)
	}

	data.IsActive = false
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	_, err = c.collection.ReplaceOne(ctx, filter, data, options.Replace().SetUpsert(true))

	return err
}

func (c *mongoGuessClient) Query(id uint32) (*MongoGuesser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var res MongoGuesser
	err := c.collection.FindOne(
		ctx,
		bson.D{primitive.E{Key: "id", Value: id}}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, err
}
