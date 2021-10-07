package utils

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func GetMongoClientAndDb() (*mongo.Client, *mongo.Database) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ip, port, err := GetServiceDNS("mongo")
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	opts := options.Client().ApplyURI("mongodb://" + ip + ":" + port)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	db := client.Database("onboarding")

	return client, db
}

func UpdateCorrectGuessTransaction(num uint32, mongoNum *MongoFoundNumber, guess uint32, mongoGuess *MongoFoundGuesser) error {
	client, db := GetMongoClientAndDb()
	numCollection := db.Collection("number")
	guessCollection := db.Collection("guess")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	logic := func(sessCtx mongo.SessionContext) error {
		opts := &options.UpdateOptions{}
		opts = opts.SetUpsert(true)
		filter := bson.M{"num": num}
		update := bson.M{"$push": bson.M{"found": *mongoNum}}
		res, err := numCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
		if res.ModifiedCount != 1 {
			return fmt.Errorf("Number %d not found", num)
		}

		filter = bson.M{"id": guess}
		update = bson.M{"$push": bson.M{"found": *mongoGuess}}
		res, err = guessCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
		if res.ModifiedCount != 1 {
			return fmt.Errorf("Guess %d not found", guess)
		}

		return nil
	}

	return RunInTransaction(ctx, session, logic)
}

func RunInTransaction(ctx context.Context, sess mongo.Session, f func(sessCtx mongo.SessionContext) error) error {
	opts := &options.UpdateOptions{}
	opts = opts.SetUpsert(false)
	sessCtx := mongo.NewSessionContext(ctx, sess)
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	err := sess.StartTransaction(txnOpts)
	if err != nil {
		return err
	}
	err = f(sessCtx)
	if err != nil {
		if abortErr := sess.AbortTransaction(sessCtx); abortErr != nil {
			fmt.Printf("error trying to abort failed transaction: %s", abortErr)
		}
		return err
	}
	err = sess.CommitTransaction(sessCtx)
	if err != nil {
		if abortErr := sess.AbortTransaction(sessCtx); abortErr != nil {
			fmt.Printf("error trying to abort failed transaction: %s", abortErr)
		}
	}
	return err
}
