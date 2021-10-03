package service

import (
	"time"

	"onboarding.com/utils"
)

var mongoPrime *utils.MongoPrimeClient

func init() {
	mongoPrime = utils.GetPrimeClient()
}

func AddPrime(prime, num, guesserId uint32, foundAt time.Time) error {
	primeData := utils.MongoPrime{
		Prime:     prime,
		Num:       num,
		GuesserId: guesserId,
		FoundAt:   foundAt,
	}
	err := mongoPrime.Add(&primeData)
	return err
}

func QueryPrime() ([]*utils.MongoPrime, error) {
	return mongoPrime.Query()
}
