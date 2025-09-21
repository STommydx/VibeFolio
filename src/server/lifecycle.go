package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/STommydx/VibeFolio/src/observability"
	"go.uber.org/zap"
)

// ServerManager manages the lifecycle of the HTTP server
type ServerManager struct {
	server *Server
	logger *observability.Logger
}

// NewServerManager creates a new server manager
func NewServerManager(server *Server, logger *observability.Logger) *ServerManager {
	return &ServerManager{
		server: server,
		logger: logger,
	}
}

// Run runs the server and handles graceful shutdown
func (sm *ServerManager) Run() error {
	// Start the server
	if err := sm.server.Start(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	sm.logger.Info("Server started successfully")

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	sm.logger.Info("Shutdown signal received")

	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Stop the server
	if err := sm.server.Stop(ctx); err != nil {
		sm.logger.Error("Failed to stop server", zap.Error(err))
		return err
	}

	sm.logger.Info("Server shutdown completed")
	return nil
}
