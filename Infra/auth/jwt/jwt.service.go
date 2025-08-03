package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	SecretKey []byte
}

type Claims struct {
	UserId         uuid.UUID   `json:"userId"`
	Email          string   `json:"email"`
	Username       string   `json:"username"`
	Roles          []string `json:"roles"`
	IsFromKeycloak bool     `json:"isFromKeycloak"`
	jwt.RegisteredClaims
}

func NewJWTService(secretKey string) *JWTService {
	return &JWTService{
		SecretKey: []byte(secretKey),
	}
}

func (j *JWTService) GenerateToken(userId uuid.UUID, email, username string, roles []string, isFromKeycloak bool) (string, error) {
	claims := Claims{
		UserId:         userId,
		Email:          email,
		Username:       username,
		Roles:          roles,
		IsFromKeycloak: isFromKeycloak,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-gin-boilerplate",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SecretKey)
}

func (j *JWTService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")

}
