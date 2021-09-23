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

	// server.GET("/guessers/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.JSON(http.StatusOK)
	// })

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
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	})

	// server.GET("/nums/:num"

	server.POST("/nums", func(c *gin.Context) {
		var body NumBody
		c.BindJSON(&body)

		err := numService.Add(body.Num)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		c.Status(http.StatusOK)
	})

	server.DELETE("/nums/:num", func(c *gin.Context) {
		var body NumBody
		c.BindJSON(&body)

		err := numService.Remove(body.Num)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		c.Status(http.StatusOK)
	})

	// server.GET("/primes"

	// Listen and serve on 0.0.0.0:8080
	server.Run(":8080")
}
