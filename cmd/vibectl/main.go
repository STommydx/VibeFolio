package main

import (
	"fmt"
	"os"

	"github.com/STommydx/VibeFolio/src/app"
	"github.com/STommydx/VibeFolio/src/cli"
)

func main() {
	// If no arguments are provided, run the CLI
	if len(os.Args) == 1 {
		cli.Execute()
		return
	}

	// If the first argument is "serve", run the application directly
	if os.Args[1] == "serve" {
		application := app.NewApp()
		if err := application.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Otherwise, run the CLI
	cli.Execute()
}
