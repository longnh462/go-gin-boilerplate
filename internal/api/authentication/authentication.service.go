package authentication

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/longnh462/go-gin-boilerplate/Infra/auth/jwt"
	"github.com/longnh462/go-gin-boilerplate/Infra/auth/keycloak"
	"github.com/longnh462/go-gin-boilerplate/internal/api/authentication/dto"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	keycloakService *keycloak.KeycloakService
	jwtService      *jwt.JWTService
	authRepo        *AuthRepository
}

func NewAuthService(authRepo *AuthRepository) *AuthService {
	// Lấy JWT secret từ environment variable
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		jwtSecret = "your-default-secret-key" // Fallback, nên cảnh báo trong production
	}

	return &AuthService{
		keycloakService: keycloak.NewKeycloakService(),
		jwtService:      jwt.NewJWTService(jwtSecret),
		authRepo:        authRepo,
	}
}

func (as *AuthService) LoginUser(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1. Thử login với Keycloak trước
	keycloakToken, err := as.keycloakService.LoginUser(req.Email, req.Password)
	if err == nil {
		// Login thành công với Keycloak
		// Parse token để lấy thông tin user từ Keycloak
		userInfo, err := as.getUserInfoFromKeycloakToken(keycloakToken.AccessToken)
		if err != nil {
			return nil, fmt.Errorf("failed to get user info from keycloak: %w", err)
		}

		// Tạo JWT token local với thông tin từ Keycloak
		localToken, err := as.jwtService.GenerateToken(
			userInfo.UserID,
			userInfo.Email,
			userInfo.Username,
			userInfo.Roles,
			true, // isFromKeycloak = true
		)
		if err != nil {
			return nil, fmt.Errorf("failed to generate local token: %w", err)
		}

		return &dto.LoginResponse{
			Token:        localToken,
			RefreshToken: keycloakToken.RefreshToken,
			Type:         "keycloak",
			ExpiresIn:    3600, // 1 hour
		}, nil
	}

	// 2. Nếu Keycloak fail, thử login với local database
	user, err := as.authRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Verify password
	if !as.verifyPassword(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Lấy roles của user
	roles, err := as.authRepo.GetUserRoles(user.UserId)
	if err != nil {
		roles = []string{"user"} // Default role nếu không lấy được roles
	}

	// Tạo JWT token local
	localToken, err := as.jwtService.GenerateToken(
		user.UserId,
		user.Email,
		user.Username,
		roles,
		false, // isFromKeycloak = false
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.LoginResponse{
		Token:     localToken,
		Type:      "local",
		ExpiresIn: 86400, // 24 hours
	}, nil
}

func (as *AuthService) ValidateToken(tokenString string) (*jwt.Claims, error) {
	// Validate JWT token
	claims, err := as.jwtService.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Nếu token từ Keycloak, có thể validate thêm với Keycloak
	// (Optional - có thể bỏ qua để giảm API calls)
	// if claims.IsFromKeycloak {
		// Uncomment nếu muốn validate với Keycloak mỗi request
		// _, err := as.keycloakService.ValidateToken(tokenString)
		// if err != nil {
		//     return nil, fmt.Errorf("keycloak token validation failed: %w", err)
		// }
	// }

	return claims, nil
}

func (as *AuthService) RefreshToken(refreshToken string) (*dto.LoginResponse, error) {
	// Thử refresh với Keycloak
	newToken, err := as.keycloakService.RefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	// Parse token để lấy thông tin user
	userInfo, err := as.getUserInfoFromKeycloakToken(newToken.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info from refreshed token: %w", err)
	}

	// Tạo JWT token local mới
	localToken, err := as.jwtService.GenerateToken(
		userInfo.UserID,
		userInfo.Email,
		userInfo.Username,
		userInfo.Roles,
		true,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new local token: %w", err)
	}

	return &dto.LoginResponse{
		Token:        localToken,
		RefreshToken: newToken.RefreshToken,
		Type:         "keycloak",
		ExpiresIn:    3600,
	}, nil
}

// Helper functions
func (as *AuthService) verifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// func (as *AuthService) hashPassword(password string) (string, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashedPassword), nil
// }

type UserInfo struct {
	UserID   uuid.UUID
	Username string
	Email    string
	Roles    []string
}

func (as *AuthService) getUserInfoFromKeycloakToken(accessToken string) (*UserInfo, error) {
	// Simplified version - parse JWT claims directly
	// In production, you might want to call Keycloak's userinfo endpoint
	claims, err := as.jwtService.ValidateToken(accessToken)
	if err != nil {
		// Fallback: parse basic info from token
		// This is a simplified approach
		return &UserInfo{
			UserID:   uuid.New(), // Replace with actual parsing logic
			Username: "keycloak-user",
			Email:    "user@keycloak.local",
			Roles:    []string{"user"},
		}, nil
	}

	return &UserInfo{
		UserID:   claims.UserId,
		Username: claims.Username,
		Email:    claims.Email,
		Roles:    claims.Roles,
	}, nil
}
