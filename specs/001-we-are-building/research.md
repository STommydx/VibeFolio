# Research Findings: Portfolio Management System

## Summary
This document captures the research findings for the portfolio management system implementation, focusing on the technology stack and architectural decisions.

## Research Topics

### 1. Go Web Framework Options
**Decision**: Use danielgtaylor/huma
**Rationale**: Huma provides OpenAPI generation, validation, and serialization out of the box. It's designed for building REST APIs with automatic documentation.
**Alternatives considered**: 
- Gin (more manual setup required for OpenAPI)
- Echo (similar to Gin but with slightly better performance)
- Fiber (fast but less mature ecosystem)

### 2. UI Templating and Rendering
**Decision**: Use a-h/templ for server-side rendering with HTMX for reactive components
**Rationale**: templ provides type-safe templates that compile to Go code, ensuring compile-time checking. HTMX allows for reactive UI without heavy client-side JavaScript.
**Alternatives considered**:
- Standard html/template (less type safety)
- GoFiber with templating (requires separate templating engine)
- React/Vue with Go backend (more complex deployment)

### 3. Database ORM
**Decision**: Use ent ORM
**Rationale**: Ent provides a powerful and flexible ORM with code generation, schema migration support, and good performance. It allows for easy database switching.
**Alternatives considered**:
- GORM (popular but can be slower)
- SQLBoiler (good performance but less flexible)
- Raw SQL with database/sql package (more control but more boilerplate)

### 4. Messaging and Storage
**Decision**: Use NATS JetStream for work queue and object storage
**Rationale**: NATS JetStream provides reliable messaging with persistence and can also serve as an object store. For initial development, we'll use an embedded NATS server.
**Alternatives considered**:
- RabbitMQ (heavier, requires separate installation)
- Apache Kafka (more complex for this use case)
- Redis with Pub/Sub (less persistence guarantees)

### 5. Dependency Injection
**Decision**: Use uber/fx
**Rationale**: Fx provides a robust dependency injection framework that works well with Go's explicit error handling. It helps manage the application lifecycle.
**Alternatives considered**:
- Wire (compile-time DI but more complex setup)
- DIG (simpler but less feature-rich)
- Manual dependency injection (more boilerplate)

### 6. Telemetry and Logging
**Decision**: Use OpenTelemetry for telemetry and zap for logging
**Rationale**: OpenTelemetry is the industry standard for observability, providing vendor-neutral APIs. Zap is a fast, structured logging library.
**Alternatives considered**:
- Prometheus metrics with custom logging (more manual setup)
- Other logging libraries (less performance or features)

### 7. Configuration Management
**Decision**: Use HCL format with XDG compliant directory storage
**Rationale**: HCL provides human-readable configuration with good tooling support. XDG compliance ensures proper placement on different operating systems.
**Alternatives considered**:
- YAML/JSON (less expressive than HCL)
- Environment variables only (less structured)
- TOML (similar to HCL but less tooling)

### 8. CLI Framework
**Decision**: Use cobra
**Rationale**: Cobra is the de facto standard for CLI applications in Go, used by major projects like Kubernetes and Hugo.
**Alternatives considered**:
- kingpin (simpler but less features)
- flag (standard library but limited)

### 9. Authentication Approach
**Decision**: Use standard OAuth 2.0 with external identity providers
**Rationale**: Leveraging established identity providers (Google, GitHub, LinkedIn) reduces development complexity and security risks. Users benefit from familiar authentication flows and established trust relationships.
**Implementation Approach**: 
- Integrate with OAuth 2.0 providers using standard protocols
- Store only OAuth subject identifiers in UserProfile entity
- Use OAuth access tokens for API authorization
- Support multiple identity providers simultaneously

## Implementation Approach
Based on the research, we'll implement a layered architecture:
1. Repository Layer: Using ent ORM for data persistence
2. Service Layer: Business logic implementation
3. Controller Layer: Huma handlers for API endpoints and templ/gin handlers for UI pages

The application will be packaged as a single binary with embedded NATS server for simplified deployment.