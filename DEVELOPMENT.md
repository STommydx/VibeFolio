# Development Guide

This guide provides instructions for developing and contributing to VibeFolio.

## Prerequisites

- Go 1.21 or later
- Git
- [Ginkgo](https://onsi.github.io/ginkgo/) for testing

## Project Structure

```
VibeFolio/
├── cmd/
│   └── healthcheck/          # Main application entry point
├── src/
│   ├── app/                  # Application core
│   ├── cli/                  # Command-line interface
│   ├── config/               # Configuration parsing
│   ├── controllers/           # HTTP controllers
│   ├── di/                   # Dependency injection
│   ├── errors/               # Error handling
│   ├── models/               # Data models
│   ├── observability/        # Logging and tracing
│   ├── responses/            # Response formatting
│   ├── server/               # HTTP server
│   └── services/             # Business logic
├── tests/
│   ├── contract/             # Contract tests
│   ├── integration/          # Integration tests
│   └── unit/                 # Unit tests
├── docs/                     # Documentation
└── specs/                    # Feature specifications
```

## Building

```bash
# Build the application
go build -o vibectl cmd/healthcheck/main.go

# Build with race detector
go build -race -o vibectl cmd/healthcheck/main.go
```

## Testing

```bash
# Run all tests
ginkgo -r

# Run unit tests
ginkgo ./tests/unit/...

# Run integration tests
ginkgo ./tests/integration/...

# Run contract tests
ginkgo ./tests/contract/...

# Run tests with coverage
ginkgo -r -cover
```

## Code Quality

```bash
# Run linter
golangci-lint run

# Format code
go fmt ./...

# Vet code
go vet ./...
```

## Adding New Features

1. Create a new specification in `specs/`
2. Run the planning commands to generate tasks
3. Implement tests first (TDD)
4. Implement the feature
5. Update documentation
6. Run all tests
7. Submit a pull request

## Dependency Injection

VibeFolio uses Uber FX for dependency injection. When adding new components:

1. Register them in `src/di/modules.go`
2. Use parameter objects for constructors with multiple dependencies
3. Use result objects for constructors that return multiple values

## Observability

VibeFolio includes structured logging with Zap and tracing with OpenTelemetry:

1. Use the provided logger in `src/observability/logger.go`
2. Add spans for important operations
3. Include relevant context in log messages

## Configuration

VibeFolio supports configuration through:

1. HCL files
2. Environment variables
3. Command-line flags

When adding new configuration options:

1. Add them to the `Configuration` model in `src/models/config.go`
2. Update the HCL parser in `src/config/hcl_parser.go`
3. Update the environment variable parser in `src/config/env_parser.go`
4. Update the flag parser in `src/config/flag_parser.go`
5. Add validation in `src/config/validation.go`

## API Development

VibeFolio uses Huma for API development:

1. Create controllers in `src/controllers/`
2. Register routes in the controller's `RegisterRoutes` method
3. Use Huma's built-in validation and documentation features
4. Follow REST conventions

## Error Handling

VibeFolio uses structured error handling:

1. Use the `AppError` type in `src/errors/app_error.go`
2. Return appropriate HTTP status codes
3. Include meaningful error messages and details
4. Log errors appropriately

## Release Process

1. Update version numbers
2. Update CHANGELOG.md
3. Create a git tag
4. Build binaries for supported platforms
5. Create a GitHub release
