# Data Model: Minimal HTTP API with Health Check Endpoint

## HealthStatus Entity

**Description**: Represents the health status of the service.

**Fields**:
- `status` (string, required): Current health status of the service. Valid values: "healthy", "unhealthy"
- `timestamp` (string, RFC3339 format, required): ISO8601 timestamp when the health check was performed
- `version` (string, optional): Version of the service
- `details` (object, optional): Additional details about the health status

**Validation Rules**:
- Status must be one of the predefined values
- Timestamp must be in RFC3339 format
- Version should follow semantic versioning if provided

## Configuration Entity

**Description**: Represents the configuration parameters for the service.

**Fields**:
- `port` (int, required): Port number the HTTP server should listen on. Valid range: 1-65535
- `host` (string, optional): Host address to bind to. Default: "localhost"
- `readTimeout` (int, optional): Read timeout in seconds. Default: 30
- `writeTimeout` (int, optional): Write timeout in seconds. Default: 30

**Validation Rules**:
- Port must be in valid range (1-65535)
- Host must be a valid hostname or IP address
- Timeouts must be positive integers
