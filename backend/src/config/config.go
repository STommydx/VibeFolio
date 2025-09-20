package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"golang.org/x/sys/unix"
)

// Config represents the application configuration
type Config struct {
	Server   ServerConfig   `hcl:"server,block"`
	Database DatabaseConfig `hcl:"database,block"`
	NATS     NATSConfig     `hcl:"nats,block"`
	OAuth    OAuthConfig    `hcl:"oauth,block"`
	Logging  LoggingConfig  `hcl:"logging,block"`
}

// ServerConfig represents the server configuration
type ServerConfig struct {
	Port int    `hcl:"port"`
	Host string `hcl:"host"`
}

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	Driver string `hcl:"driver"`
	DSN    string `hcl:"dsn"`
}

// NATSConfig represents the NATS configuration
type NATSConfig struct {
	URL     string `hcl:"url"`
	Embedded bool   `hcl:"embedded"`
}

// OAuthProvider represents an OAuth provider configuration
type OAuthProvider struct {
	Name        string `hcl:"name"`
	ClientID    string `hcl:"client_id"`
	ClientSecret string `hcl:"client_secret"`
	RedirectURL  string `hcl:"redirect_url"`
}

// OAuthConfig represents the OAuth configuration
type OAuthConfig struct {
	Providers []OAuthProvider `hcl:"providers,block"`
}

// LoggingConfig represents the logging configuration
type LoggingConfig struct {
	Level  string `hcl:"level"`
	Format string `hcl:"format"`
}

// Load loads the configuration from the default location
func Load() (*Config, error) {
	// Try to find config file in XDG config directory
	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get config path: %w", err)
	}

	// If config file doesn't exist, create a default one
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createDefaultConfig(configPath); err != nil {
			return nil, fmt.Errorf("failed to create default config: %w", err)
		}
	}

	// Load the configuration
	var config Config
	if err := hclsimple.DecodeFile(configPath, nil, &config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &config, nil
}

// getConfigPath returns the path to the configuration file
func getConfigPath() (string, error) {
	// Get XDG config directory
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		// Default to ~/.config
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		xdgConfigHome = filepath.Join(homeDir, ".config")
	}

	// Create the application config directory if it doesn't exist
	appConfigDir := filepath.Join(xdgConfigHome, "vibefolio")
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create config directory: %w", err)
	}

	return filepath.Join(appConfigDir, "config.hcl"), nil
}

// createDefaultConfig creates a default configuration file
func createDefaultConfig(path string) error {
	defaultConfig := `server {
  port = 8080
  host = "localhost"
}

database {
  driver = "sqlite"
  dsn = "./data/vibefolio.db"
}

nats {
  url = "nats://localhost:4222"
  embedded = true
}

oauth {
  providers = [
    {
      name = "google"
      client_id = "your-google-client-id"
      client_secret = "your-google-client-secret"
      redirect_url = "http://localhost:8080/auth/callback"
    },
    {
      name = "github"
      client_id = "your-github-client-id"
      client_secret = "your-github-client-secret"
      redirect_url = "http://localhost:8080/auth/callback"
    }
  ]
}

logging {
  level = "info"
  format = "json"
}
`

	return os.WriteFile(path, []byte(defaultConfig), 0644)
}