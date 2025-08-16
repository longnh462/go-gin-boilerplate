package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	// Create Gin Router
	router := gin.Default()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "7070"
	}

	fmt.Printf("Server is running on localhost:%s: 🔥🔥🔥🔥🔥 \n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
