package utils

import (
	"time"
)

type MongoFoundNumber struct {
	GuesserId uint32    `bson:"guesser_id,omitempty" json:"guesser_id,omitempty"`
	Attempt   int64     `bson:"attempt,omitempty" json:"attempt,omitempty"`
	FoundAt   time.Time `bson:"found_at,omitempty" json:"found_at,omitempty"`
}

type MongoNumber struct {
	Num      uint32             `bson:"num,required" json:"num,required"`
	IsActive bool               `bson:"is_active,omitempty" json:"is_active,omitempty"`
	Found    []MongoFoundNumber `bson:"found,omitempty" json:"guesses,omitempty"`
}

type MongoFoundGuesser struct {
	Num     uint32    `bson:"num,omitempty" json:"num,omitempty"`
	Attempt int64     `bson:"attempt,omitempty" json:"attempt,omitempty"`
	FoundAt time.Time `bson:"found_at,omitempty" json:"found_at,omitempty"`
}

type MongoGuesser struct {
	Id            uint32              `bson:"id,omitempty" json:"id,omitempty"`
	IsActive      bool                `bson:"is_active,omitempty" json:"is_active,omitempty"`
	BeginAt       uint32              `bson:"begin_at,omitempty" json:"begin_at,omitempty"`
	IncrementBy   uint32              `bson:"increment_by,omitempty" json:"increment_by,omitempty"`
	SleepInterval uint32              `bson:"sleep_interval,omitempty" json:"sleep_interval,omitempty"`
	Found         []MongoFoundGuesser `bson:"found,omitempty" json:"guesses,omitempty"`
}

type MongoPrime struct {
	Prime     uint32    `bson:"prime,required" json:"prime,required"`
	GuesserId uint32    `bson:"guesser_id,required" json:"guesser_id,required"`
	Num       uint32    `bson:"num,required" json:"num,required"`
	FoundAt   time.Time `bson:"found_at,required" json:"found_at,required"`
}
