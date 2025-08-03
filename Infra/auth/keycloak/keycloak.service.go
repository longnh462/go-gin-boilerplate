package keycloak

import (
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
