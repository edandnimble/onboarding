package api

import (
	"fmt"
	"net/http"
	"os"
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

		res, err := guessService.Query(uint32(id))
		if err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}

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

	server.GET("/primes", func(c *gin.Context) {
		fmt.Printf("Get primes")

		res, err := service.QueryPrime()
		if err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	// Listen and serve
	httpPort := os.Getenv("API_HTTP_PORT")
	server.Run(":" + httpPort)
}
