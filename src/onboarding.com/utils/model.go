package utils

import (
	"time"
)

type MongoFoundNumber struct {
	GuesserId uint32    `bson:"guesser_id,omitempty" json:"guesser_id,omitempty"`
	Attempt   int64     `bson:"attempt,omitempty" json:"attempt,omitempty"`
	Time      time.Time `bson:"time,omitempty" json:"time,omitempty"`
}

type MongoNumber struct {
	Num      uint32             `bson:"num,required" json:"num,required"`
	IsActive bool               `bson:"is_active,omitempty" json:"is_active,omitempty"`
	Found    []MongoFoundNumber `bson:"found,omitempty" json:"guesses,omitempty"`
}

type MongoFoundGuesser struct {
	Num     uint32    `bson:"num,omitempty" json:"num,omitempty"`
	Attempt int64     `bson:"attempt,omitempty" json:"attempt,omitempty"`
	Time    time.Time `bson:"time,omitempty" json:"time,omitempty"`
}

type MongoGuesser struct {
	Id            uint32              `bson:"id,omitempty" json:"id,omitempty"`
	IsActive      bool                `bson:"is_active,omitempty" json:"is_active,omitempty"`
	BeginAt       uint32              `bson:"begin_at,omitempty" json:"begin_at,omitempty"`
	IncrementBy   uint32              `bson:"increment_by,omitempty" json:"increment_by,omitempty"`
	SleepInterval uint32              `bson:"sleep_interval,omitempty" json:"sleep_interval,omitempty"`
	Found         []MongoFoundGuesser `bson:"found,omitempty" json:"guesses,omitempty"`
}
