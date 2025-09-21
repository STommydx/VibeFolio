package services

import (
	"os"
	"strconv"

	"github.com/STommydx/VibeFolio/src/config"
	"github.com/STommydx/VibeFolio/src/models"
)

// ConfigService handles configuration loading and parsing
type ConfigService struct {
	config *models.Configuration
}

// NewConfigService creates a new ConfigService instance
func NewConfigService() *ConfigService {
	return &ConfigService{
		config: &models.Configuration{
			Port:         9090,
			Host:         "localhost",
			ReadTimeout:  30,
			WriteTimeout: 30,
		},
	}
}

// LoadConfiguration loads configuration from various sources
func (c *ConfigService) LoadConfiguration() error {
	// Load from environment variables
	if portStr := os.Getenv("PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			c.config.Port = port
		}
	}

	if host := os.Getenv("HOST"); host != "" {
		c.config.Host = host
	}

	// In a real implementation, we would also load from:
	// 1. HCL configuration files
	// 2. Command line flags
	// 3. Default values

	return nil
}

// GetConfiguration returns the current configuration
func (c *ConfigService) GetConfiguration() *models.Configuration {
	return c.config
}

// ValidateConfiguration validates the current configuration
func (c *ConfigService) ValidateConfiguration() error {
	cfg := c.config

	if err := config.ValidatePort(cfg.Port); err != nil {
		return err
	}

	if err := config.ValidateTimeout(cfg.ReadTimeout); err != nil {
		return err
	}

	if err := config.ValidateTimeout(cfg.WriteTimeout); err != nil {
		return err
	}

	if cfg.Host != "" {
		if err := config.ValidateHost(cfg.Host); err != nil {
			return err
		}
	}

	return nil
}
