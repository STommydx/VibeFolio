package nats

import (
	"fmt"
	"os"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

// Server represents an embedded NATS server
type Server struct {
	*server.Server
	config *server.Options
}

// Config holds the NATS server configuration
type Config struct {
	URL     string
	Embedded bool
	DataDir string
	Port    int
}

// NewServer creates a new NATS server instance
func NewServer(cfg Config) (*Server, error) {
	// If not embedded, return nil to indicate we should connect to external server
	if !cfg.Embedded {
		return nil, nil
	}

	// Create data directory if it doesn't exist
	if cfg.DataDir == "" {
		cfg.DataDir = "./data/nats"
	}
	
	if err := os.MkdirAll(cfg.DataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create NATS data directory: %w", err)
	}

	// Configure the embedded server
	opts := &server.Options{
		ServerName: "VibeFolio-NATS",
		Host:       "localhost",
		Port:       cfg.Port,
		StoreDir:   cfg.DataDir,
		JetStream:  true,
		// Enable JetStream with file storage
		JetStreamMaxMemory: 1024 * 1024 * 1024, // 1GB
		JetStreamMaxStore:  10 * 1024 * 1024 * 1024, // 10GB
	}

	// Create the server
	ns, err := server.NewServer(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create NATS server: %w", err)
	}

	return &Server{
		Server: ns,
		config: opts,
	}, nil
}

// Start starts the embedded NATS server
func (s *Server) Start() error {
	if s.Server == nil {
		return nil // Nothing to start if not embedded
	}

	// Start the server in a goroutine
	go s.Server.Start()

	// Wait for the server to be ready
	if !s.Server.ReadyForConnections(10 * 1000000000) { // 10 seconds
		return fmt.Errorf("NATS server failed to start")
	}

	return nil
}

// Stop stops the embedded NATS server
func (s *Server) Stop() {
	if s.Server != nil {
		s.Server.Shutdown()
	}
}

// Connect creates a connection to the NATS server
func (s *Server) Connect() (*nats.Conn, error) {
	var url string
	if s.Server != nil {
		// Connect to our embedded server
		url = s.Server.ClientURL()
	} else {
		// Connect to external server
		// This would come from config in a real implementation
		url = nats.DefaultURL
	}

	// Connect to the server
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	return nc, nil
}

// ConnectJS creates a JetStream connection to the NATS server
func (s *Server) ConnectJS() (nats.JetStreamContext, error) {
	nc, err := s.Connect()
	if err != nil {
		return nil, err
	}

	// Create JetStream context
	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, fmt.Errorf("failed to create JetStream context: %w", err)
	}

	return js, nil
}