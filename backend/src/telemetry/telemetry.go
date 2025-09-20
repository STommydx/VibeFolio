package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger represents a structured logger
type Logger struct {
	*zap.Logger
}

// Tracer represents an OpenTelemetry tracer
type Tracer struct {
	trace.Tracer
}

// Config holds telemetry configuration
type Config struct {
	LogLevel  string
	LogFormat string
}

// Setup initializes telemetry (logging and tracing)
func Setup(cfg Config) (*Logger, *Tracer, error) {
	// Set up logging
	logger, err := setupLogging(cfg)
	if err != nil {
		return nil, nil, err
	}

	// Set up tracing
	tracer, err := setupTracing()
	if err != nil {
		logger.Error("Failed to set up tracing", zap.Error(err))
		return logger, nil, err
	}

	return logger, tracer, nil
}

// setupLogging initializes the zap logger
func setupLogging(cfg Config) (*Logger, error) {
	var zapConfig zap.Config
	
	if cfg.LogFormat == "json" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Set log level
	level, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = zapcore.InfoLevel
	}
	zapConfig.Level.SetLevel(level)

	zapLogger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{zapLogger}, nil
}

// setupTracing initializes OpenTelemetry tracing
func setupTracing() (*Tracer, error) {
	// Create a stdout exporter for development
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	// Create a resource describing the service
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("vibefolio"),
			semconv.ServiceVersion("0.1.0"),
		),
	)
	if err != nil {
		return nil, err
	}

	// Create a trace provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// Set the global trace provider
	otel.SetTracerProvider(tp)

	// Set the propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Get a tracer
	tracer := otel.Tracer("vibefolio")

	return &Tracer{tracer}, nil
}

// StartSpan starts a new span
func (t *Tracer) StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return t.Tracer.Start(ctx, name)
}

// WithLogger adds a logger to the context
func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

// FromContext retrieves a logger from the context
func FromContext(ctx context.Context) *Logger {
	if logger, ok := ctx.Value("logger").(*Logger); ok {
		return logger
	}
	// Return a default logger if none is found in context
	logger, _ := setupLogging(Config{LogLevel: "info", LogFormat: "json"})
	return logger
}