# Minimal HTTP API with Health Check Endpoint

## Summary

This pull request implements a minimal HTTP API with a health check endpoint as specified in feature 001. The service provides a `/healthz` endpoint that returns the health status of the service and supports configurable port settings.

## Features Delivered

### ✅ Core Implementation
- HTTP server built with Huma framework
- GET `/healthz` endpoint returning JSON health status
- Configurable port via environment variables
- Graceful shutdown handling

### ✅ Architecture
- Layered architecture following repository → service → controller pattern
- Dependency injection with Uber FX
- Structured logging with Zap
- Observability with OpenTelemetry tracing
- CLI interface with Cobra

### ✅ Testing
- Comprehensive test coverage with Ginkgo/Gomega
- Contract tests (3/3 passing)
- Integration tests (4/4 passing)
- Unit tests (27/27 passing)
- Performance and load tests

## Current Limitations

### Partial Configuration Implementation
While the specification requires support for three configuration methods (command-line flags, environment variables, HCL files), only environment variable parsing is currently fully integrated. The HCL and flag parsers have been implemented as components but are not yet integrated into the application.

**Files implemented but not integrated:**
- `src/config/hcl_parser.go`
- `src/config/flag_parser.go`

This represents a partial implementation of the configuration requirements and should be addressed in future work.

## Testing

All tests pass:
```bash
ginkgo -r
```

## Usage

```bash
# Build
go build -o bin/vibectl cmd/vibectl/main.go

# Run with default port (9090)
./bin/vibectl serve

# Run with custom port
PORT=8080 ./bin/vibectl serve

# Test health endpoint
curl http://localhost:9090/healthz
```
