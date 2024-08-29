package config

import (
	"os"
	"testing"
)

// TestLoadConfig tests the LoadConfig function to ensure it loads configuration correctly from environment variables.
func TestLoadConfig(t *testing.T) {
	// Set up test environment variables
	os.Setenv("DB_CONNECTION_URL", "postgres://user:password@localhost:5432/testdb")
	os.Setenv("ETH_NODE_URL", "http://localhost:8545")
	os.Setenv("API_PORT", "9090")
	os.Setenv("JWT_SECRET", "testsecret")

	// Load the configuration
	cfg := LoadConfig()

	// Test if the configuration values match the expected values
	if cfg.DatabaseDSN != "postgres://user:password@localhost:5432/testdb" {
		t.Errorf("Expected DatabaseDSN to be 'postgres://user:password@localhost:5432/testdb', got '%s'", cfg.DatabaseDSN)
	}

	if cfg.EthereumNodeURL != "http://localhost:8545" {
		t.Errorf("Expected EthereumNodeURL to be 'http://localhost:8545', got '%s'", cfg.EthereumNodeURL)
	}

	if cfg.APIPort != "9090" {
		t.Errorf("Expected APIPort to be '9090', got '%s'", cfg.APIPort)
	}

	if cfg.JWTSecret != "testsecret" {
		t.Errorf("Expected JWTSecret to be 'testsecret', got '%s'", cfg.JWTSecret)
	}

	// Test with missing optional environment variables
	os.Unsetenv("API_PORT")
	os.Unsetenv("JWT_SECRET")

	cfg = LoadConfig()

	if cfg.APIPort != "8080" {
		t.Errorf("Expected default APIPort to be '8080', got '%s'", cfg.APIPort)
	}

	if cfg.JWTSecret != "defaultsecret" {
		t.Errorf("Expected default JWTSecret to be 'defaultsecret', got '%s'", cfg.JWTSecret)
	}

	// Test with missing required environment variables
	os.Unsetenv("DB_CONNECTION_URL")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to missing DB_CONNECTION_URL environment variable")
		}
	}()
	LoadConfig() // This should cause a log.Fatal call, which will exit the test
}
