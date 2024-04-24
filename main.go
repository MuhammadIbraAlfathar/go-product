package main

import (
	"fmt"
	"github.com/MuhammadIbraAlfathar/go-product/config"
	"github.com/MuhammadIbraAlfathar/go-product/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route.UserRouter(api)

	err := r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
	if err != nil {
		log.Fatal("Error connection to port")
	}
}
