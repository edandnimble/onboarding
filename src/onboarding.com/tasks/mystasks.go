package tasks

import (
	"math"
	"time"

	"onboarding.com/api/service"
)

func FindPrime(num, guesserId uint32, foundAt int64) error {
	searchValue := num
	for isPrime(searchValue) == false {
		searchValue++
	}

	err := service.AddPrime(searchValue, num, guesserId, time.Unix(foundAt, 0))
	if err != nil {
		return err
	}

	return nil
}

func isPrime(value uint32) bool {
	for i := uint32(2); i <= uint32(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
