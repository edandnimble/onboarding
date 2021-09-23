package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoClientAndDb() (*mongo.Client, *mongo.Database) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client := mongo.Connect(ctx, opts)
	db := client.Database("onboarding")

	return &client, &db
}

type mongoNumberClient struct {
	client     mongo.Client
	db         mongo.Database
	collection *mongo.Collection
}

type mongoGuessClient struct {
	client     mongo.Client
	db         mongo.Database
	collection *mongo.Collection
}

func GetNumberClient() *mongoNumberClient {
	client, db := getMongoClientAndDb()
	collection := db.Collection("number")
	// collection.Indexes().CreateOne(
	// 	context.Background(),
	// 	mongo.IndexModel{
	// 		Keys:    bson.D{primitive.E{Key: "Num", Value: 0}},
	// 		Options: nil},
	// )
	return &mongoNumberClient{
		client:     client,
		db:         db,
		collection: collection}
}

func GetGuessClient() *mongoGuessClient {
	client, db := getMongoClientAndDb()
	collection := db.Collection("guess")
	return &mongoGuessClient{
		client:     client,
		db:         db,
		collection: collection}
}

func (c *mongoNumberClient) Add(num uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.UpdateOptions{}.SetUpsert(true)
	data := MongoNumber{IsActive: true}
	_, err := c.collection.UpdateOne(
		ctx,
		bson.D{primitive.E{Key: "Num", Value: num}},
		bson.D{primitive.E{Key: "$set", Value: data}},
		opts)
	return err
}

func (c *mongoNumberClient) Remove(num uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.UpdateOptions{}.SetUpsert(true)
	data := MongoNumber{IsActive: false}
	_, err := c.collection.UpdateOne(
		ctx,
		bson.D{primitive.E{Key: "Num", Value: num}},
		bson.D{primitive.E{Key: "$set", Value: data}},
		opts)
	return err
}

func (c *mongoNumberClient) Query(num uint32) (*MongoNumber, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var res MongoNumber
	err := c.collection.FindOne(
		ctx,
		bson.D{primitive.E{Key: "Num", Value: num}}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, err
}
