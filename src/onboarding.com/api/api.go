package api

import (
	"fmt"
	"net/http"
	"strconv"

	"onboarding.com/api/service"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	server := gin.Default()
	numService, err := service.NewNumService()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	guessService, err := service.NewGuessService()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	server.GET("/guessers/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}
		fmt.Printf("Get guesser %d\n", id)

		res, err := guessService.Query(uint32(id))
		if err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}
		fmt.Printf("Guesser: %v\n", res)

		c.JSON(http.StatusOK, res)
	})

	server.POST("/guessers", func(c *gin.Context) {
		var body GuesserBody
		c.BindJSON(&body)

		err := guessService.Add(body.Begin, body.IncrementBy, body.SleepInterval)
		if err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	})

	server.DELETE("/guessers/:id", func(c *gin.Context) {
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
	})

	server.GET("/nums/:num", func(c *gin.Context) {
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
	})

	server.POST("/nums", func(c *gin.Context) {
		var body NumBody
		c.BindJSON(&body)

		err := numService.Add(body.Num)
		if err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	})

	server.DELETE("/nums/:num", func(c *gin.Context) {
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
	})

	// server.GET("/primes"

	// Listen and serve on 0.0.0.0:8080
	server.Run(":8080")
}
