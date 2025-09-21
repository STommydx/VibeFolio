package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/STommydx/VibeFolio/src/controllers"
	"github.com/STommydx/VibeFolio/src/models"
	"github.com/STommydx/VibeFolio/src/observability"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"go.uber.org/zap"
)

// Server represents the HTTP server
type Server struct {
	httpServer *http.Server
	logger     *observability.Logger
	config     *models.Configuration
}

// NewServer creates a new HTTP server
func NewServer(config *models.Configuration, logger *observability.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}

// Setup sets up the HTTP server with all components
func (s *Server) Setup(healthController *controllers.HealthController) error {
	// Create Huma API
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("VibeFolio Health Check", "1.0.0"))

	// Register routes
	healthController.RegisterRoutes(api)

	// Add logging middleware
	handler := observability.LoggingMiddleware(s.logger)(mux)

	// Create server
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		Handler:      handler,
		ReadTimeout:  time.Duration(s.config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.WriteTimeout) * time.Second,
	}

	s.logger.Info("HTTP server configured",
		zap.String("host", s.config.Host),
		zap.Int("port", s.config.Port),
		zap.Int("read_timeout", s.config.ReadTimeout),
		zap.Int("write_timeout", s.config.WriteTimeout),
	)

	return nil
}

// Start starts the HTTP server
func (s *Server) Start() error {
	if s.httpServer == nil {
		return fmt.Errorf("server not configured")
	}

	s.logger.Info("Starting HTTP server", zap.String("addr", s.httpServer.Addr))

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("HTTP server error", zap.Error(err))
		}
	}()

	return nil
}

// Stop stops the HTTP server gracefully
func (s *Server) Stop(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}

	s.logger.Info("Stopping HTTP server")

	// Create a context with timeout for shutdown
	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Shutdown the server
	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		s.logger.Error("HTTP server shutdown error", zap.Error(err))
		return err
	}

	s.logger.Info("HTTP server stopped")
	return nil
}
