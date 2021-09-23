package utils

import (
	"context"
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
			Keys:    bson.D{primitive.E{Key: "id", Value: 0}},
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
	opts := (&options.UpdateOptions{}).SetUpsert(true)
	data := MongoGuesser{IsActive: false}
	_, err := c.collection.UpdateOne(
		ctx,
		bson.D{primitive.E{Key: "id", Value: id}},
		bson.D{primitive.E{Key: "$set", Value: data}},
		opts)
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
