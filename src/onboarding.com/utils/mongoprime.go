package utils

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPrimeClient struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func GetPrimeClient() *MongoPrimeClient {
	client, db := GetMongoClientAndDb()
	collection := db.Collection("primes")
	return &MongoPrimeClient{
		client:     client,
		db:         db,
		collection: collection}
}

func (c *MongoPrimeClient) Add(prime *MongoPrime) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.collection.InsertOne(ctx, prime)
	return err
}

func (c *MongoPrimeClient) Query() ([]*MongoPrime, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := c.collection.Find(ctx, primitive.D{})
	if err != nil {
		return nil, err
	}

	var res []*MongoPrime
	err = cursor.All(ctx, &res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return res, nil
}
