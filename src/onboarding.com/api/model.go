package api

type NumBody struct {
	Num uint32 `json:"num"`
}

type GuesserBody struct {
	Begin         uint32 `json:"begin"`
	IncrementBy   uint32 `json:"incrementBy"`
	SleepInterval uint32 `json:"sleepInterval"`
}
