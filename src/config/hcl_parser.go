package config

import (
	"fmt"

	"github.com/STommydx/VibeFolio/src/models"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

// HCLParser parses HCL configuration files
type HCLParser struct{}

// NewHCLParser creates a new HCL parser
func NewHCLParser() *HCLParser {
	return &HCLParser{}
}

// ParseConfig parses an HCL configuration file and returns a Configuration model
func (p *HCLParser) ParseConfig(filename string) (*models.Configuration, error) {
	var config models.Configuration

	// Set default values
	config.Port = 9090
	config.Host = "localhost"
	config.ReadTimeout = 30
	config.WriteTimeout = 30

	// Parse the HCL file
	err := hclsimple.DecodeFile(filename, nil, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HCL configuration file: %w", err)
	}

	return &config, nil
}
