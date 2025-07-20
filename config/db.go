package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {

	// Connection to database
	fmt.Println("Conecction to database......")
	connectionDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	DB, err := sql.Open("postgres", connectionDB)

	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database PING ERROR!!!")
	}

	fmt.Println("Connected Sucessfully to database")

}
