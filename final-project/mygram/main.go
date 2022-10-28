package main

import (
	"log"
	"mygram/config"
	"mygram/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.StartDB()
	r := router.StartApp()
	r.Run(":" + os.Getenv("PORT"))
}
