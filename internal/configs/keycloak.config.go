package configs

type KeycloakConfig struct {
	Host         string
	ClientID     string
	ClientSecret string
	Realm        string
}

func GetKeycloakConfig() *KeycloakConfig {
	return &KeycloakConfig{
		Host:         getEnv("KEYCLOAK_HOST", "http://localhost:8080"),
		ClientID:     getEnv("KEYCLOAK_CLIENT_ID", "go-gin-client"),
		ClientSecret: getEnv("KEYCLOAK_CLIENT_SECRET", "go-gin-secret"),
		Realm:        getEnv("KEYCLOAK_REALM", "go-gin-realm"),
	}
}

