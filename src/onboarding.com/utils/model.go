package utils

import (
	"time"
)

type MongoGuess struct {
	GuesserId uint32    `bson:"guesser_id,omitempty"`
	Time      time.Time `bson:"time,omitempty"`
}

type MongoNumber struct {
	Num      uint32       `bson:"num,required"`
	IsActive bool         `bson:"is_active,omitempty"`
	Guesses  []MongoGuess `bson:"guesses,omitempty"`
}
