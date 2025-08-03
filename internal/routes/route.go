package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/longnh462/go-gin-boilerplate/internal/api/authentication"
	"github.com/longnh462/go-gin-boilerplate/middlewares"
)

func SetupAuthRoutes(router *gin.RouterGroup, db *sql.DB) {
	// Initialize services
	authRepo := authentication.NewAuthRepository(db)
	authService := authentication.NewAuthService(authRepo)
	authHandler := authentication.NewAuthHandler(authService)

	// Auth routes (no middleware required)
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/refresh", authHandler.RefreshToken)

		// Protected auth routes
		authProtected := authRoutes.Group("/")
		authProtected.Use(middlewares.AuthMiddleware(authService))
		{
			authProtected.POST("/logout", authHandler.Logout)
			authProtected.GET("/profile", authHandler.GetProfile)
		}
	}
}

func SetupProtectedRoutes(router *gin.RouterGroup, db *sql.DB) {
	// Initialize services
	authRepo := authentication.NewAuthRepository(db)
	authService := authentication.NewAuthService(authRepo)

	// All protected routes
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware(authService))
	{
		// User routes
		userRoutes := protected.Group("/users")
		{
			userRoutes.GET("/", getUsersList)   // Get all users
			userRoutes.GET("/:id", getUserById) // Get user by ID
			userRoutes.PUT("/:id", updateUser)  // Update user

			// Admin only routes
			adminRoutes := userRoutes.Group("/")
			adminRoutes.Use(middlewares.RequireRole("admin"))
			{
				adminRoutes.DELETE("/:id", deleteUser) // Delete user (admin only)
				adminRoutes.POST("/", createUser)      // Create user (admin only)
			}
		}

		// Role management routes (admin only)
		roleRoutes := protected.Group("/roles")
		roleRoutes.Use(middlewares.RequireRole("admin"))
		{
			roleRoutes.GET("/", getRolesList)
			roleRoutes.POST("/", createRole)
			roleRoutes.PUT("/:id", updateRole)
			roleRoutes.DELETE("/:id", deleteRole)
		}
	}
}

func getUsersList(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get users list"})
}

func getUserById(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get user by ID"})
}

func updateUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update user"})
}

func deleteUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Delete user"})
}

func createUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Create user"})
}

func getRolesList(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get roles list"})
}

func createRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Create role"})
}

func updateRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update role"})
}

func deleteRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Delete role"})
}
