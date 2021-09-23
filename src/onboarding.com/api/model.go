package api

import "time"

type NumBody struct {
	Num uint32 `json:"num,required"`
}

type Guess struct {
	GuesserId uint32    `json:"guesser_id,required"`
	Found     time.Time `json:"found_time,required"`
	Attempt   uint32    `json:"attempt_num,required"`
}

type NumGetResp struct {
	Num      uint32  `json:"num,required"`
	IsActive bool    `json:"is_active,required"`
	Guesses  []Guess `json:"guesses,required"`
}

type GuesserBody struct {
	Begin         uint32 `json:"begin,required"`
	IncrementBy   uint32 `json:"increment_by,required"`
	SleepInterval uint32 `json:"sleep_interval,required"`
}
