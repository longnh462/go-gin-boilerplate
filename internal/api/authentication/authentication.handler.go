package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/longnh462/go-gin-boilerplate/internal/api/authentication/dto"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login godoc
// @Summary Login User
// @Description Login with email and password (supports both Keycloak and local authentication)
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginRequest body dto.LoginRequest true "Login Request"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (ah *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := ah.authService.LoginUser(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// RefreshToken godoc
// @Summary Refresh Token
// @Description Refresh access token using refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param refreshRequest body dto.RefreshTokenRequest true "Refresh Token Request"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/refresh [post]
func (ah *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ah.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Logout godoc
// @Summary Logout User
// @Description Logout current user
// @Tags Authentication
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string
// @Router /auth/logout [post]
func (ah *AuthHandler) Logout(c *gin.Context) {
	// For JWT tokens, logout is typically handled client-side
	// But you can implement token blacklisting here if needed
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// GetProfile godoc
// @Summary Get User Profile
// @Description Get current user profile information
// @Tags Authentication
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /auth/profile [get]
func (ah *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	username, _ := c.Get("username")
	email, _ := c.Get("email")
	roles, _ := c.Get("roles")
	isFromKeycloak, _ := c.Get("is_from_keycloak")

	profile := gin.H{
		"user_id":          userID,
		"username":         username,
		"email":            email,
		"roles":            roles,
		"is_from_keycloak": isFromKeycloak,
	}

	c.JSON(http.StatusOK, profile)
}
