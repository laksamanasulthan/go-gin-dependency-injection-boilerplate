package main

import (
	"go-gin-dependency-injection-boilerplate/injector"
	"go-gin-dependency-injection-boilerplate/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	inj, err := injector.NewInjector()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.RegisterRoutes(router, inj)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := router.Run(":7890"); err != nil {
		log.Fatal(err)
	}
}
