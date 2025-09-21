# Quickstart Guide: Minimal HTTP API with Health Check Endpoint

## Prerequisites

- Go 1.21 or later
- Git

## Building the Service

1. Clone the repository:
   ```
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Build the service:
   ```
   go build -o healthcheck ./cmd/healthcheck
   ```

## Running the Service

1. Run the service with default configuration:
   ```
   ./healthcheck serve
   ```

2. Run the service with custom port:
   ```
   ./healthcheck serve --port 9090
   ```

3. Run the service with configuration file:
   ```
   ./healthcheck serve --config config.hcl
   ```

## Testing the Health Check Endpoint

1. Check if the service is healthy:
   ```
   curl -v http://localhost:8080/healthz
   ```

   Expected response:
   ```
   HTTP/1.1 200 OK
   Content-Type: application/json

   {
     "status": "healthy",
     "timestamp": "2025-09-21T10:00:00Z",
     "version": "1.0.0"
   }
   ```

## Configuration

The service can be configured in three ways:

1. Command line flags:
   ```
   ./healthcheck serve --port 9090 --host 0.0.0.0
   ```

2. Environment variables:
   ```
   PORT=9090 HOST=0.0.0.0 ./healthcheck serve
   ```

3. HCL configuration file:
   ```hcl
   port = 9090
   host = "0.0.0.0"
   read_timeout = 30
   write_timeout = 30
   ```

## Stopping the Service

To stop the service, send an interrupt signal (Ctrl+C) or SIGTERM.
