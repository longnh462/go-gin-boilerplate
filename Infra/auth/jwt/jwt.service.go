package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey []byte
}

type Claims struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	Username string `json:"username"`
	Roles []string `json:"roles"`
	IsFromKeycloak bool `json:"isFromKeycloak"`
	jwt.RegisteredClaims
}

func NewJWTService(secretKey string) *JWTService {
    return &JWTService{
        SecretKey: []byte(secretKey),
    }
}