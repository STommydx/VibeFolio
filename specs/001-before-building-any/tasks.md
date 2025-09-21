# Tasks: Minimal HTTP API with Health Check Endpoint

**Input**: Design documents from `/specs/001-before-building-any/`
**Prerequisites**: plan.md (required), research.md, data-model.md, contracts/

## Execution Flow (main)
```
1. Load plan.md from feature directory
   → If not found: ERROR "No implementation plan found"
   → Extract: tech stack, libraries, structure
2. Load optional design documents:
   → data-model.md: Extract entities → model tasks
   → contracts/: Each file → contract test task
   → research.md: Extract decisions → setup tasks
3. Generate tasks by category:
   → Setup: project init, dependencies, linting
   → Tests: contract tests, integration tests
   → Core: models, services, CLI commands
   → Integration: DB, middleware, logging
   → Polish: unit tests, performance, docs
4. Apply task rules:
   → Different files = mark [P] for parallel
   → Same file = sequential (no [P])
   → Tests before implementation (TDD)
5. Number tasks sequentially (T001, T002...)
6. Generate dependency graph
7. Create parallel execution examples
8. Validate task completeness:
   → All contracts have tests?
   → All entities have models?
   → All endpoints implemented?
9. Return: SUCCESS (tasks ready for execution)
```

## Format: `[ID] [P?] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- Include exact file paths in descriptions

## Path Conventions
- **Single project**: `src/`, `tests/` at repository root
- **Web app**: `backend/src/`, `frontend/src/`
- **Mobile**: `api/src/`, `ios/src/` or `android/src/`
- Paths shown below assume single project - adjust based on plan.md structure

## Phase 3.1: Setup
- [x] T001 Create project structure per implementation plan
- [x] T002 Initialize Go project with Huma, Fx, OpenTelemetry, Zap, HCL, Cobra, Ginkgo dependencies
- [x] T003 [P] Configure linting and formatting tools (golangci-lint)
- [x] T004 [P] Set up Git pre-commit hooks for code formatting and linting

## Phase 3.2: Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.3
**CRITICAL: These tests MUST be written and MUST FAIL before ANY implementation**

- [ ] T005 [P] Contract test GET /healthz endpoint in tests/contract/healthz_test.go
- [ ] T006 [P] Integration test health check endpoint in tests/integration/healthz_test.go
- [ ] T007 [P] Integration test configuration loading in tests/integration/config_test.go
- [ ] T008 [P] Unit test health service logic in tests/unit/health_service_test.go
- [ ] T009 [P] Unit test configuration parsing in tests/unit/config_test.go

## Phase 3.3: Core Implementation (ONLY after tests are failing)

### Models
- [ ] T010 [P] HealthStatus model in src/models/health_status.go
- [ ] T011 [P] Configuration model in src/models/config.go

### Services
- [ ] T012 Health service in src/services/health_service.go
- [ ] T013 Configuration service in src/services/config_service.go

### CLI Commands
- [ ] T014 [P] Root CLI command in src/cli/root.go
- [ ] T015 [P] Serve command in src/cli/serve.go

### API Endpoints
- [ ] T016 GET /healthz endpoint in src/controllers/health_controller.go

### Core Components
- [ ] T017 Input validation for configuration
- [ ] T018 Error handling and response formatting
- [ ] T019 Application startup and shutdown logic

## Phase 3.4: Integration

### Dependency Injection
- [ ] T020 Set up Fx dependency injection container in src/di/container.go
- [ ] T021 Register services with Fx in src/di/modules.go

### Observability
- [ ] T022 Integrate OpenTelemetry tracing in src/observability/tracer.go
- [ ] T023 Set up Zap structured logging in src/observability/logger.go
- [ ] T024 Request/response logging middleware

### Configuration
- [ ] T025 HCL configuration parsing in src/config/hcl_parser.go
- [ ] T026 Environment variable configuration in src/config/env_parser.go
- [ ] T027 Command line flag configuration in src/config/flag_parser.go

### HTTP Server
- [ ] T028 HTTP server setup with Huma in src/server/server.go
- [ ] T029 Server lifecycle management (start/stop)

## Phase 3.5: Polish

### Testing
- [ ] T030 [P] Unit tests for error handling in tests/unit/error_test.go
- [ ] T031 [P] Unit tests for logging in tests/unit/logging_test.go
- [ ] T032 Performance tests for health endpoint (<100ms)
- [ ] T033 Load testing for concurrent health checks

### Documentation
- [ ] T034 [P] Update README.md with project documentation
- [ ] T035 [P] Update docs/api.md with API documentation
- [ ] T036 [P] Create example configuration files
- [ ] T037 [P] Update DEVELOPMENT.md with development guide

### Quality
- [ ] T038 Remove code duplication
- [ ] T039 Run manual-testing.md scenarios
- [ ] T040 Verify all acceptance criteria from spec.md
- [ ] T041 Security audit for HTTP server configuration

## Dependencies

- Setup (T001-T004) before everything
- Tests (T005-T009) before implementation (T010-T019)
- Models (T010-T011) before services (T012-T013)
- Services (T012-T013) before controllers (T016)
- Core implementation (T010-T019) before integration (T020-T029)
- Integration (T020-T029) before polish (T030-T041)
- T010 blocks T012
- T011 blocks T013
- T020 blocks T028
- T025-T027 block T013

## Parallel Example

```
# Launch T005-T009 together:
Task: "Contract test GET /healthz endpoint in tests/contract/healthz_test.go"
Task: "Integration test health check endpoint in tests/integration/healthz_test.go"
Task: "Integration test configuration loading in tests/integration/config_test.go"
Task: "Unit test health service logic in tests/unit/health_service_test.go"
Task: "Unit test configuration parsing in tests/unit/config_test.go"

# Launch T010-T011 together:
Task: "HealthStatus model in src/models/health_status.go"
Task: "Configuration model in src/models/config.go"

# Launch T014-T015 together:
Task: "Root CLI command in src/cli/root.go"
Task: "Serve command in src/cli/serve.go"

# Launch T030-T037 together:
Task: "Unit tests for error handling in tests/unit/error_test.go"
Task: "Unit tests for logging in tests/unit/logging_test.go"
Task: "Update README.md with project documentation"
Task: "Update docs/api.md with API documentation"
Task: "Create example configuration files"
Task: "Update DEVELOPMENT.md with development guide"
```

## Notes

- [P] tasks = different files, no dependencies
- Verify tests fail before implementing
- Commit after each task
- Avoid: vague tasks, same file conflicts

## Task Generation Rules
*Applied during main() execution*

1. **From Contracts**:
   - Each contract file → contract test task [P]
   - Each endpoint → implementation task

2. **From Data Model**:
   - Each entity → model creation task [P]
   - Relationships → service layer tasks

3. **From User Stories**:
   - Each story → integration test [P]
   - Quickstart scenarios → validation tasks

4. **Ordering**:
   - Setup → Tests → Models → Services → Endpoints → Polish
   - Dependencies block parallel execution

## Validation Checklist
*GATE: Checked by main() before returning*

- [x] All contracts have corresponding tests
- [x] All entities have model tasks
- [x] All tests come before implementation
- [x] Parallel tasks truly independent
- [x] Each task specifies exact file path
- [x] No task modifies same file as another [P] task
