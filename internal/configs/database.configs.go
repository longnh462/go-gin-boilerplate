package configs

import (
	"fmt"
	"os"

	// "github.com/longnh462/go-gin-boilerplate/Infra/database/postgres/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	Sslmode  string
	Schema   string
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getEnv("POSTGRES_HOST", "localhost"),
		User:     getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", "postgres"),
		DBName:   getEnv("POSTGRES_DB", "go-gin-boilerplate"),
		Port:     getEnv("POSTGRES_PORT", "5432"),
		Sslmode:  getEnv("POSTGRES_SSL", "disable"),
		Schema:   getEnv("POSTGRES_SCHEMA", "ggb"),
	}
}

func (config *DatabaseConfig) getDbConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s options=-csearch_path=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.Sslmode,
		config.Schema,
	)
}

func ConnectDb() (*gorm.DB, error) {
	config := GetDatabaseConfig()
	dbConnectionString := config.getDbConnectionString()
	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// if err:= migrations.AutoMigrate(db); err != nil {
	// 	return nil, fmt.Errorf("failed to run migrations: %w", err)
	// }

	return db, nil
}
