package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stommydx/VibeFolio/backend/src/config"
	"go.uber.org/zap"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "vibefolio",
		Short: "VibeFolio is a portfolio management system",
		Long:  "A comprehensive portfolio management system with resume generation capabilities",
	}

	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the VibeFolio server",
		Run: func(cmd *cobra.Command, args []string) {
			// Load configuration
			cfg, err := config.Load()
			if err != nil {
				fmt.Printf("Failed to load configuration: %v\n", err)
				os.Exit(1)
			}

			// Initialize logger
			logger, _ := zap.NewProduction()
			defer logger.Sync()

			logger.Info("Starting VibeFolio server...",
				zap.String("host", cfg.Server.Host),
				zap.Int("port", cfg.Server.Port))

			// Server implementation will go here
			fmt.Printf("Server would start on %s:%d\n", cfg.Server.Host, cfg.Server.Port)
		},
	}

	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
		Long:  "View or edit the VibeFolio configuration",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Load()
			if err != nil {
				fmt.Printf("Failed to load configuration: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Configuration loaded successfully:\n")
			fmt.Printf("  Server: %s:%d\n", cfg.Server.Host, cfg.Server.Port)
			fmt.Printf("  Database: %s://%s\n", cfg.Database.Driver, cfg.Database.DSN)
			fmt.Printf("  NATS: %s (embedded: %t)\n", cfg.NATS.URL, cfg.NATS.Embedded)
			fmt.Printf("  OAuth Providers: %d\n", len(cfg.OAuth.Providers))
			fmt.Printf("  Logging: %s format=%s\n", cfg.Logging.Level, cfg.Logging.Format)
		},
	}

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(configCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}