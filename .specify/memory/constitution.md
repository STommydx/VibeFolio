<!-- SYNC IMPACT REPORT
Version change: 1.0.0 → 1.1.0
Modified principles: 
- I. Focus on Correctness and Best Practices (new principle added as first priority)
- II. API-First Design (renumbered from I)
- III. Test-Driven Development (renumbered from II)
- IV. Security by Default (renumbered from III)
- V. Observability through Telemetry (renumbered from IV)
- VI. Layered Architecture (renumbered from V)
Added sections: None
Removed sections: None
Templates requiring updates: None
Follow-up TODOs: None
-->

# VibeFolio Constitution

## Core Principles

### I. Focus on Correctness and Best Practices
All development MUST prioritize correctness over speed. NO ASSUMPTIONS should be made during implementation. Documentation MUST be consulted, and best practices MUST be researched with references to existing projects when applicable. Each task MUST have a concrete and verifiable goal to ensure correctness before moving on to the next task.

### II. API-First Design
All functionality MUST be accessible through well-defined APIs before any UI implementation. UI components MUST use the same APIs that external consumers would use. This ensures consistent behavior and enables multiple clients to leverage the same functionality.

### III. Test-Driven Development
All code MUST be developed using Test-Driven Development practices. Tests MUST be written before implementation, and all code MUST have associated tests that validate its behavior. Integration tests MUST cover all API endpoints and critical user flows.

### IV. Security by Default
Security MUST be considered at every stage of development. Access control mechanisms MUST be implemented for all resources. Authentication and authorization MUST be enforced at the API level. All data transmission MUST be encrypted, and sensitive data MUST be properly protected.

### V. Observability through Telemetry
Structured logging and telemetry MUST be integrated into all components. All significant operations MUST generate appropriate log entries with consistent formatting. Metrics MUST be collected to monitor system health and performance. Error conditions MUST be properly tracked and reported.

### VI. Layered Architecture
The system MUST follow a layered architecture with clear separation of concerns:
- Repository Layer: Data access and persistence logic
- Service Layer: Business logic and domain operations
- Controller Layer: Transport layer and UI components
Each layer MUST be independently testable and loosely coupled with other layers.

## Architecture Requirements

The system MUST follow a layered architecture pattern with clear separation of concerns. Each layer MUST be independently deployable and testable:

1. **Repository Layer**: Handles data access and persistence operations
   - Encapsulates data storage mechanisms
   - Provides a clean API for data operations
   - Isolates data concerns from business logic

2. **Service Layer**: Contains business logic and domain operations
   - Implements core application functionality
   - Coordinates between different repositories
   - Enforces business rules and validation

3. **Controller Layer**: Manages transport layer and user interface
   - Handles HTTP requests and responses
   - Converts between external and internal data formats
   - Manages user interactions and presentation

Each layer MUST be independently testable through unit tests, and integration between layers MUST be verified through integration tests.

## Security Standards

All development MUST adhere to security best practices:

1. **Authentication and Authorization**
   - All API endpoints MUST require authentication unless explicitly public
   - Role-based access control MUST be implemented
   - Permissions MUST be validated for all operations

2. **Data Protection**
   - Sensitive data MUST be encrypted at rest and in transit
   - Personal information MUST be handled according to privacy regulations
   - Data retention and deletion policies MUST be enforced

3. **Secure Coding Practices**
   - Input validation MUST be performed on all external data
   - SQL injection and other common vulnerabilities MUST be prevented
   - Dependencies MUST be regularly updated and security-audited

## Governance

The Constitution supersedes all other development practices and guidelines. Any amendments to this Constitution MUST follow a formal process including documentation, approval, and implementation plan.

All pull requests and code reviews MUST verify compliance with these constitutional principles. Significant violations MUST be addressed before merging.

Versioning of this Constitution follows semantic versioning:
- MAJOR version for backward-incompatible governance changes
- MINOR version for new principles or materially expanded guidance
- PATCH version for clarifications and non-semantic refinements

**Version**: 1.1.0 | **Ratified**: 2025-09-21 | **Last Amended**: 2025-09-21