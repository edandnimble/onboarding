package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"onboarding.com/api/service"

	"github.com/gin-gonic/gin"
)

var guessService *service.GuessService
var numService *service.NumService

func init() {
	var err error
	numService, err = service.NewNumService()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	guessService, err = service.NewGuessService()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func getGuesser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	res, err := guessService.Query(uint32(id))
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, res)
}

func createGuesser(c *gin.Context) {
	var body GuesserBody
	c.BindJSON(&body)

	err := guessService.Add(body.Begin, body.IncrementBy, body.SleepInterval)
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func deleteGuesser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	err = guessService.Remove(uint32(id))
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func getNumber(c *gin.Context) {
	num, err := strconv.ParseUint(c.Param("num"), 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	res, err := numService.Query(uint32(num))
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, res)
}

func deleteNumber(c *gin.Context) {
	num, err := strconv.ParseUint(c.Param("num"), 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Println("Removing num ", num)

	err = numService.Remove(uint32(num))
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func createNumber(c *gin.Context) {
	var body NumBody
	c.BindJSON(&body)

	err := numService.Add(body.Num)
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func getPrimes(c *gin.Context) {
	fmt.Printf("Get primes")

	res, err := service.QueryPrime()
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, res)
}

func RunServer() {
	server := gin.Default()

	server.GET("/guessers/:id", getGuesser)

	server.POST("/guessers", createGuesser)

	server.DELETE("/guessers/:id", deleteGuesser)

	server.GET("/nums/:num", getNumber)

	server.POST("/nums", createNumber)

	server.DELETE("/nums/:num", deleteNumber)

	server.GET("/primes", getPrimes)

	// Listen and serve
	httpPort := os.Getenv("API_HTTP_PORT")
	server.Run(":" + httpPort)
}
