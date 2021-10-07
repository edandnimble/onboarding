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

type mongoNumberClient struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func GetNumberClient() *mongoNumberClient {
	client, db := GetMongoClientAndDb()
	collection := db.Collection("number")
	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{primitive.E{Key: "num", Value: 1}},
			Options: options.Index().SetUnique(true)},
	)
	if err != nil {
		fmt.Println("Error creating index ", err.Error())
		return nil
	}
	return &mongoNumberClient{
		client:     client,
		db:         db,
		collection: collection}
}

func (c *mongoNumberClient) Add(num uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := c.Query(num)
	if err == mongo.ErrNoDocuments {
		data := MongoNumber{Num: num, IsActive: true}
		_, err = c.collection.InsertOne(ctx, data)
		return err
	}

	if err != nil {
		return err
	}

	fmt.Printf("Number %v already exists, activating\n", num)
	if err != nil {
		return err
	}

	if data.IsActive == true {
		return fmt.Errorf("Number %v is already active", num)
	}

	data.IsActive = true

	filter := bson.D{primitive.E{Key: "num", Value: num}}
	_, err = c.collection.ReplaceOne(ctx, filter, data, options.Replace().SetUpsert(true))

	return err

}

func (c *mongoNumberClient) Remove(num uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := c.Query(num)
	if err != nil {
		return err
	}

	if data.IsActive == false {
		return fmt.Errorf("Number %v is not active", num)
	}

	data.IsActive = false

	filter := bson.D{primitive.E{Key: "num", Value: num}}
	_, err = c.collection.ReplaceOne(ctx, filter, data, options.Replace().SetUpsert(true))

	return err
}

func (c *mongoNumberClient) Query(num uint32) (*MongoNumber, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var res MongoNumber
	err := c.collection.FindOne(
		ctx,
		bson.D{primitive.E{Key: "num", Value: num}}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *mongoNumberClient) IsExist(num uint32) (bool, error) {
	res, err := c.Query(num)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return res.IsActive, nil
}
