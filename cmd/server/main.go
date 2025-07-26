package main

import (
	"fmt"
	"log"

	"github.com/longnh462/go-gin-boilerplate/internal/configs"
)

func main() {
	fmt.Println("Starting Go Gin Application.... 🚀🚀🚀🚀🚀")

	// Connect to database
	db, err := configs.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to create db instance 🚫🚫🚫🚫🚫 %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to create sqlDB instance 🚫🚫🚫🚫🚫 %v", err)
	}
	defer sqlDB.Close()

	fmt.Println("Successfully connected to database 🆙🆙🆙🆙🆙")
}
