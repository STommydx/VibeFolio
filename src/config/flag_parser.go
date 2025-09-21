package config

import (
	"flag"
	"fmt"

	"github.com/STommydx/VibeFolio/src/models"
)

// FlagParser parses configuration from command line flags
type FlagParser struct {
	fs *flag.FlagSet
}

// NewFlagParser creates a new command line flag parser
func NewFlagParser() *FlagParser {
	fs := flag.NewFlagSet("config", flag.ContinueOnError)
	return &FlagParser{
		fs: fs,
	}
}

// ParseConfig parses configuration from command line flags and returns a Configuration model
func (p *FlagParser) ParseConfig(args []string) (*models.Configuration, error) {
	config := &models.Configuration{
		Port:         9090, // Default port
		Host:         "localhost",
		ReadTimeout:  30,
		WriteTimeout: 30,
	}

	// Define flags
	port := p.fs.Int("port", config.Port, "Port to listen on")
	host := p.fs.String("host", config.Host, "Host to bind to")
	readTimeout := p.fs.Int("read-timeout", config.ReadTimeout, "Read timeout in seconds")
	writeTimeout := p.fs.Int("write-timeout", config.WriteTimeout, "Write timeout in seconds")
	configFile := p.fs.String("config", "", "Path to configuration file")

	// Parse flags
	if err := p.fs.Parse(args); err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}

	// Update config with parsed values
	config.Port = *port
	config.Host = *host
	config.ReadTimeout = *readTimeout
	config.WriteTimeout = *writeTimeout
	config.ConfigFile = *configFile

	return config, nil
}
