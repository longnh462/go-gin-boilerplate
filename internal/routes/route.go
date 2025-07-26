package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/longnh462/go-gin-boilerplate/internal/handlers"
)

func SetupUserRoutes(router *gin.RouterGroup) {
    userRoutes := router.Group("/users")
    {
        userRoutes.GET("/", handlers.ListUsers)
        userRoutes.GET("/:id", handlers.GetUser)
    }
}