package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
			fmt.Println("Starting VibeFolio server...")
			// Server implementation will go here
		},
	}

	rootCmd.AddCommand(serveCmd)
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}