package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type UserResponse struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

// GetUser godoc
// @Summary      Get a user by ID
// @Description  Returns a user based on the ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  UserResponse
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/{id} [get]
func GetUser(c *gin.Context) {
    // id := c.Param("id")

    // This would typically fetch from the database
    // Mock response for demonstration
    user := UserResponse{
        ID:       1,
        Username: "testuser",
        Email:    "test@example.com",
    }

    c.JSON(http.StatusOK, user)
}

// ListUsers godoc
// @Summary      List all users
// @Description  Returns a list of users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   UserResponse
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
func ListUsers(c *gin.Context) {
    // This would typically fetch from the database
    // Mock response for demonstration
    users := []UserResponse{
        {ID: 1, Username: "user1", Email: "user1@example.com"},
        {ID: 2, Username: "user2", Email: "user2@example.com"},
    }

    c.JSON(http.StatusOK, users)
}