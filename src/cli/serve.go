package cli

import (
	"fmt"

	"github.com/STommydx/VibeFolio/src/app"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Long: `Start the HTTP server with the health check endpoint.

The server will listen on the configured port (default: 9090) and respond
to requests on the /healthz endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting HTTP server...")

		// Create and run the application
		application := app.NewApp()
		if err := application.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	// Add flags for the serve command
	serveCmd.Flags().Int("port", 9090, "Port to listen on")
	serveCmd.Flags().String("host", "localhost", "Host to bind to")
	serveCmd.Flags().String("config", "", "Path to configuration file")

	// Add the serve command to the root command
	AddCommand(serveCmd)
}
