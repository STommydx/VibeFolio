# API Documentation

## Health Check Endpoint

### GET /healthz

Returns the health status of the service.

#### Response

**200 OK** - Service is healthy

```json
{
  "status": "healthy",
  "timestamp": "2025-09-21T10:00:00Z",
  "version": "1.0.0",
  "details": {
    "service": "VibeFolio Health Check"
  }
}
```

**503 Service Unavailable** - Service is unhealthy

```json
{
  "status": "unhealthy",
  "timestamp": "2025-09-21T10:00:00Z",
  "version": "1.0.0",
  "details": {
    "service": "VibeFolio Health Check",
    "error": "Database connection failed"
  }
}
```

#### Response Fields

| Field | Type | Description |
|-------|------|-------------|
| status | string | Current health status ("healthy" or "unhealthy") |
| timestamp | string | ISO8601 timestamp when the health check was performed |
| version | string | Version of the service |
| details | object | Additional details about the health status |

## Error Responses

All error responses follow the RFC 7807 problem details format:

```json
{
  "type": "https://example.com/probs/out-of-credit",
  "title": "You do not have enough credit.",
  "detail": "Your current balance is 30, but that costs 50.",
  "instance": "/account/12345/msgs/abc",
  "balance": 30,
  "accounts": ["/account/12345", "/account/67890"]
}
```
