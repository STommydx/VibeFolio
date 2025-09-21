package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/STommydx/VibeFolio/src/controllers"
	"github.com/STommydx/VibeFolio/src/observability"
	"github.com/STommydx/VibeFolio/src/services"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"go.uber.org/zap"
)

// App represents the main application
type App struct {
	server           *http.Server
	healthService    *services.HealthService
	healthController *controllers.HealthController
	config           *services.ConfigService
	logger           *observability.Logger
}

// NewApp creates a new App instance
func NewApp() *App {
	// Create logger
	logger, _ := observability.NewLogger("vibefolio")

	// Create services
	configService := services.NewConfigService()
	healthService := services.NewHealthService("1.0.0")

	// Create controllers
	healthController := controllers.NewHealthController(healthService)

	return &App{
		config:           configService,
		healthService:    healthService,
		healthController: healthController,
		logger:           logger,
	}
}

// Initialize initializes the application
func (a *App) Initialize() error {
	// Load configuration
	if err := a.config.LoadConfiguration(); err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Validate configuration
	if err := a.config.ValidateConfiguration(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	return nil
}

// Start starts the HTTP server
func (a *App) Start() error {
	// Get configuration
	config := a.config.GetConfiguration()

	// Create Huma API
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("VibeFolio Health Check", "1.0.0"))

	// Register routes
	a.healthController.RegisterRoutes(api)

	// Create server
	a.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:      mux,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
	}

	// Start server in a goroutine
	go func() {
		a.logger.Info("Server listening", zap.String("addr", a.server.Addr))
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Error("Server error", zap.Error(err))
		}
	}()

	return nil
}

// Stop gracefully stops the HTTP server
func (a *App) Stop() error {
	if a.server == nil {
		return nil
	}

	a.logger.Info("Stopping server")

	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error("Server shutdown error", zap.Error(err))
		return fmt.Errorf("server shutdown error: %w", err)
	}

	a.logger.Info("Server stopped")
	return nil
}

// Run runs the application
func (a *App) Run() error {
	// Initialize the application
	if err := a.Initialize(); err != nil {
		a.logger.Error("Failed to initialize application", zap.Error(err))
		return fmt.Errorf("failed to initialize application: %w", err)
	}

	// Start the server
	if err := a.Start(); err != nil {
		a.logger.Error("Failed to start server", zap.Error(err))
		return fmt.Errorf("failed to start server: %w", err)
	}

	a.logger.Info("Application started, waiting for shutdown signal")

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.logger.Info("Shutdown signal received")

	// Stop the server
	if err := a.Stop(); err != nil {
		a.logger.Error("Failed to stop server", zap.Error(err))
		return fmt.Errorf("failed to stop server: %w", err)
	}

	a.logger.Info("Application shutdown completed")
	return nil
}
