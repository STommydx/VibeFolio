# VibeFolio

VibeFolio is a modern, cloud-native application platform built with Go. It provides a minimal HTTP API with health check endpoints and is designed to be future-proof for multi-layer architecture.

## Features

- **Minimal HTTP API**: Lightweight and fast HTTP server
- **Health Check Endpoint**: Built-in health monitoring at `/healthz`
- **Configurable**: Support for HCL configuration files, environment variables, and command-line flags
- **Observability**: Integrated logging and tracing with OpenTelemetry and Zap
- **Dependency Injection**: Powered by Uber FX for clean, testable code
- **Modern Go**: Built with the latest Go features and best practices

## Quick Start

### Prerequisites

- Go 1.21 or later
- Git

### Building

```bash
git clone <repository-url>
cd VibeFolio
go build -o vibectl cmd/healthcheck/main.go
```

### Running

```bash
# Run with default configuration
./vibectl serve

# Run with custom port
./vibectl serve --port 9090

# Run with configuration file
./vibectl serve --config config.hcl
```

### Testing the Health Check

```bash
curl http://localhost:9090/healthz
```

Expected response:
```json
{
  "status": "healthy",
  "timestamp": "2025-09-21T10:00:00Z",
  "version": "1.0.0"
}
```

## Configuration

VibeFolio can be configured in three ways:

1. **Command line flags**:
   ```bash
   ./vibectl serve --port 9090 --host 0.0.0.0
   ```

2. **Environment variables**:
   ```bash
   PORT=9090 HOST=0.0.0.0 ./vibectl serve
   ```

3. **HCL configuration file**:
   ```hcl
   port = 9090
   host = "0.0.0.0"
   read_timeout = 30
   write_timeout = 30
   ```

## API Endpoints

- `GET /healthz` - Health check endpoint

## Architecture

VibeFolio follows a layered architecture:
- **Models**: Data structures and entities
- **Services**: Business logic and core functionality
- **Controllers**: HTTP request handling
- **CLI**: Command-line interface

## Development

See [DEVELOPMENT.md](DEVELOPMENT.md) for development guidelines and contribution instructions.

## License

MIT
