package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "A minimal HTTP API with health check endpoint",
	Long: `healthcheck is a minimal HTTP API service that provides a health check endpoint.

This service is part of the VibeFolio project and is designed to be a simple,
lightweight health check service that can be used as a foundation for more
complex services.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// AddCommand adds a command to the root command
func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
