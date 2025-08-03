package keycloak

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	"github.com/longnh462/go-gin-boilerplate/internal/configs"
)

type KeycloakService struct {
	client *gocloak.GoCloak
	config *configs.KeycloakConfig
}

func NewKeycloakService() *KeycloakService {
	config := configs.GetKeycloakConfig()
	client := gocloak.NewClient((config.Host))

	return &KeycloakService{
		client: client,
		config: config,
	}
}

func (ks *KeycloakService) LoginUser(username, password string) (*gocloak.JWT, error) {
	ctx := context.Background()

	token, err := ks.client.Login(ctx, ks.config.ClientID, ks.config.ClientSecret, ks.config.Realm, username, password)
	if err != nil {
		return nil, fmt.Errorf("failed to login user: %w", err)
	}
	return token, nil
}

func (ks *KeycloakService) ValidateToken(accessToken string) (*gocloak.JWT, error) {
	ctx := context.Background()

	// Validate the token
	result, err := ks.client.RetrospectToken(ctx, accessToken, ks.config.ClientID, ks.config.ClientSecret, ks.config.Realm)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}

	if !*result.Active {
		return nil, fmt.Errorf("token is not active")
	}

	return &gocloak.JWT{AccessToken: accessToken}, nil
}

func (ks *KeycloakService) RefreshToken(refreshToken string) (*gocloak.JWT, error) {
	ctx := context.Background()

	token, err := ks.client.RefreshToken(ctx, refreshToken, ks.config.ClientID, ks.config.ClientSecret, ks.config.Realm)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	return token, nil
}
