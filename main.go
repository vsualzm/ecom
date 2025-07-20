package main

import (
	"ecom/config"
	"ecom/routes"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	r := routes.SetupRoutes()
	r.Run(":8080")
}
