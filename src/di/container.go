package di

import (
	"context"
	"fmt"

	"github.com/STommydx/VibeFolio/src/observability"
	"go.uber.org/fx"
)

// Container represents the dependency injection container
type Container struct {
	app    *fx.App
	logger *observability.Logger
}

// NewContainer creates a new dependency injection container
func NewContainer() *Container {
	logger, _ := observability.NewLogger("fx-container")
	return &Container{
		logger: logger,
	}
}

// Build builds the dependency injection container with all modules
func (c *Container) Build() error {
	c.app = fx.New(
		fx.Provide(
		// Providers will be added here
		),
		fx.Invoke(
		// Invokers will be added here
		),
		fx.Logger(fxLogger{logger: c.logger}),
	)

	return nil
}

// Start starts the dependency injection container
func (c *Container) Start(ctx context.Context) error {
	if c.app == nil {
		return fmt.Errorf("container not built")
	}

	c.logger.Info("Starting Fx container")
	return c.app.Start(ctx)
}

// Stop stops the dependency injection container
func (c *Container) Stop(ctx context.Context) error {
	if c.app == nil {
		return nil
	}

	c.logger.Info("Stopping Fx container")
	return c.app.Stop(ctx)
}

// fxLogger implements fx.Printer interface for logging
type fxLogger struct {
	logger *observability.Logger
}

func (l fxLogger) Printf(str string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf("[FX] "+str, args...))
}
