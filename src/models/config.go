package models

// Configuration represents the configuration parameters for the service
type Configuration struct {
	// Port is the port number the HTTP server should listen on (required)
	// Valid range: 1-65535
	Port int `hcl:"port" validate:"required,min=1,max=65535"`

	// Host is the host address to bind to (optional)
	// Default: "localhost"
	Host string `hcl:"host,optional" validate:"omitempty,hostname"`

	// ReadTimeout is the read timeout in seconds (optional)
	// Default: 30
	ReadTimeout int `hcl:"read_timeout,optional" validate:"omitempty,min=1"`

	// WriteTimeout is the write timeout in seconds (optional)
	// Default: 30
	WriteTimeout int `hcl:"write_timeout,optional" validate:"omitempty,min=1"`

	// ConfigFile is the path to the configuration file (optional)
	ConfigFile string `hcl:"-" validate:"omitempty"`
}
