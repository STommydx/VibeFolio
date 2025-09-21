# Example Configuration

This directory contains example configuration files for VibeFolio.

## Basic Configuration

```hcl
# Basic configuration
port = 9090
host = "localhost"
read_timeout = 30
write_timeout = 30
```

## Production Configuration

```hcl
# Production configuration
port = 8080
host = "0.0.0.0"
read_timeout = 60
write_timeout = 60
```

## Development Configuration

```hcl
# Development configuration
port = 9090
host = "localhost"
read_timeout = 30
write_timeout = 30
debug = true
```
