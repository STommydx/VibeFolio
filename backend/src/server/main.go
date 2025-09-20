package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"go.uber.org/zap"

	"github.com/stommydx/VibeFolio/backend/src/api"
	"github.com/stommydx/VibeFolio/backend/src/config"
	"github.com/stommydx/VibeFolio/backend/src/telemetry"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger and tracer
	logger, _, err := telemetry.Setup(telemetry.Config{
		LogLevel:  cfg.Logging.Level,
		LogFormat: cfg.Logging.Format,
	})
	if err != nil {
		fmt.Printf("Failed to set up telemetry: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// Log startup
	logger.Info("Starting VibeFolio server...",
		zap.String("host", cfg.Server.Host),
		zap.Int("port", cfg.Server.Port))

	// Create the huma API
	mux := http.NewServeMux()
	humaAPI := humago.New(mux, huma.DefaultConfig("VibeFolio API", "1.0.0"))

	// Create handlers
	authHandler := api.NewAuthHandler(logger.Logger)
	profileHandler := api.NewProfileHandler(logger.Logger)
	sectionHandler := api.NewSectionHandler(logger.Logger)
	resumeHandler := api.NewResumeHandler(logger.Logger)

	// Register routes
	authHandler.RegisterRoutes(humaAPI)
	profileHandler.RegisterRoutes(humaAPI)
	sectionHandler.RegisterRoutes(humaAPI)
	resumeHandler.RegisterRoutes(humaAPI)

	// Start the server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	logger.Info("Server listening", zap.String("address", addr))
	
	if err := http.ListenAndServe(addr, mux); err != nil {
		logger.Error("Server failed to start", zap.Error(err))
		os.Exit(1)
	}
}