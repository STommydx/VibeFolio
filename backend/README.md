# VibeFolio Backend

This is the backend service for the VibeFolio portfolio management system.

## Features

- User profile management with multiple sections
- OAuth 2.0 authentication with Google and GitHub
- Media file upload and management
- External link management
- Profile sharing with configurable privacy settings
- Resume PDF generation using LaTeX templates
- RESTful API with OpenAPI documentation
- Structured logging and monitoring with OpenTelemetry

## Getting Started

### Prerequisites

- Go 1.22 or later
- LaTeX distribution (for PDF generation)

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd VibeFolio/backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o vibefolio .
   ```

4. Run the application:
   ```bash
   ./vibefolio serve
   ```

### Configuration

The application can be configured using a HCL configuration file located at:
- Linux: `$XDG_CONFIG_HOME/vibefolio/config.hcl` or `~/.config/vibefolio/config.hcl`
- macOS: `~/Library/Application Support/vibefolio/config.hcl`
- Windows: `%APPDATA%\vibefolio\config.hcl`

Example configuration:
```hcl
server {
  port = 8080
  host = "localhost"
}

database {
  driver = "sqlite"
  dsn = "./data/vibefolio.db"
}

nats {
  url = "nats://localhost:4222"
  embedded = true
}

oauth {
  providers = [
    {
      name = "google"
      client_id = "your-google-client-id"
      client_secret = "your-google-client-secret"
      redirect_url = "http://localhost:8080/auth/callback"
    },
    {
      name = "github"
      client_id = "your-github-client-id"
      client_secret = "your-github-client-secret"
      redirect_url = "http://localhost:8080/auth/callback"
    }
  ]
}

logging {
  level = "info"
  format = "json"
}
```

## API Documentation

The API documentation is available at `/docs` when the server is running.

## Testing

Run tests with:
```bash
go test ./...
```

## License

This project is licensed under the MIT License.