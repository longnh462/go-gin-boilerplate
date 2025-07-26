package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/longnh462/go-gin-boilerplate/docs"
	"github.com/longnh462/go-gin-boilerplate/internal/configs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Gin Boilerplate APIs
// @version         1.0.0
// @description     A Go Gin boilerplate server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://github.com/longnh462/go-gin-boilerplate
// @contact.email  longnh.uit@gmail.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:7070
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Longnh (Drake)

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

	// API v1 route group
    v1 := router.Group("/api/v1")
    {
        v1.GET("/swagger", func(context *gin.Context) {
            context.JSON(200, gin.H{
                "message": "Welcome to Go Gin Boilerplate API v1",
            })
        })
    }

	router.GET("/swaggo/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "7070"
	}

	fmt.Printf("Server is running on localhost:%s: 🔥🔥🔥🔥🔥 \n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
