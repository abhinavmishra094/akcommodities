package main

import (
	"akcomodities/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/api/exchange/", routes.Exchange())

	println(os.Getenv("PORT"))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)
}
