package config

import (
	"log"
	"os"
)

// Config struct holds the application configuration settings, such as database credentials and Ethereum node URL.
type Config struct {
	DatabaseDSN     string // DSN (Data Source Name) for PostgreSQL connection
	EthereumNodeURL string // URL of the Ethereum node to interact with
	APIPort         string // Port for the API server to listen on
	JWTSecret       string // JWT secret for authentication
}

// LoadConfig loads configuration settings from environment variables.
// It ensures that required configuration values are set and provides default values where necessary.
func LoadConfig() *Config {
	dsn := os.Getenv("DB_CONNECTION_URL")
	if dsn == "" {
		log.Fatal("DB_CONNECTION_URL environment variable is required")
	}

	ethereumNodeURL := os.Getenv("ETH_NODE_URL")
	if ethereumNodeURL == "" {
		log.Fatal("ETH_NODE_URL environment variable is required")
	}

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = "8080" // Default port if not provided
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "defaultsecret" // Default secret if not provided (for development/testing purposes only)
	}

	return &Config{
		DatabaseDSN:     dsn,
		EthereumNodeURL: ethereumNodeURL,
		APIPort:         apiPort,
		JWTSecret:       jwtSecret,
	}
}
