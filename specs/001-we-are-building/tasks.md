# Tasks: Portfolio Management System

**Input**: Design documents from `/specs/001-we-are-building/`
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
- **Web app**: `backend/src/`, `frontend/src/`
- All paths assume web application structure per plan.md

## Phase 3.1: Setup
- [x] T001 Create project structure per implementation plan in backend/
- [x] T002 Initialize Go project with huma, ent, templ, NATS dependencies in backend/go.mod
- [x] T003 [P] Configure linting and formatting tools with golangci-lint in backend/.golangci.yml
- [x] T004 Set up uber/fx dependency injection framework in backend/src/di/
- [x] T005 Configure HCL-based configuration management in backend/src/config/
- [x] T006 Set up cobra CLI framework in backend/src/cli/
- [x] T007 Initialize embedded NATS JetStream server in backend/src/nats/
- [x] T008 Set up OpenTelemetry and zap logging in backend/src/telemetry/

## Phase 3.2: Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.3
**CRITICAL: These tests MUST be written and MUST FAIL before ANY implementation**

### Contract Tests [P]
- [x] T009 [P] Contract test for /auth/login endpoint in backend/tests/contract/test_auth_login.go
- [x] T010 [P] Contract test for /auth/callback endpoint in backend/tests/contract/test_auth_callback.go
- [x] T011 [P] Contract test for GET /api/profiles endpoint in backend/tests/contract/test_get_profiles.go
- [x] T012 [P] Contract test for POST /api/profiles endpoint in backend/tests/contract/test_create_profile.go
- [x] T013 [P] Contract test for GET /api/profiles/{profileId} endpoint in backend/tests/contract/test_get_profile.go
- [x] T014 [P] Contract test for PUT /api/profiles/{profileId} endpoint in backend/tests/contract/test_update_profile.go
- [x] T015 [P] Contract test for DELETE /api/profiles/{profileId} endpoint in backend/tests/contract/test_delete_profile.go
- [x] T016 [P] Contract test for GET /api/profiles/{profileId}/sections endpoint in backend/tests/contract/test_get_sections.go
- [x] T017 [P] Contract test for POST /api/profiles/{profileId}/sections endpoint in backend/tests/contract/test_create_section.go
- [x] T018 [P] Contract test for PUT /api/profiles/{profileId}/sections/{sectionId} endpoint in backend/tests/contract/test_update_section.go
- [x] T019 [P] Contract test for DELETE /api/profiles/{profileId}/sections/{sectionId} endpoint in backend/tests/contract/test_delete_section.go
- [x] T020 [P] Contract test for POST /api/profiles/{profileId}/resume endpoint in backend/tests/contract/test_generate_resume.go

### Integration Tests [P]
- [x] T021 [P] Integration test for OAuth 2.0 flow with Google provider in backend/tests/integration/test_oauth_google.go
- [x] T022 [P] Integration test for OAuth 2.0 flow with GitHub provider in backend/tests/integration/test_oauth_github.go
- [x] T023 [P] Integration test for profile creation and management in backend/tests/integration/test_profile_management.go
- [x] T024 [P] Integration test for profile section management in backend/tests/integration/test_section_management.go
- [x] T025 [P] Integration test for media upload and storage in backend/tests/integration/test_media_upload.go
- [x] T026 [P] Integration test for external link management in backend/tests/integration/test_link_management.go
- [x] T027 [P] Integration test for profile sharing functionality in backend/tests/integration/test_profile_sharing.go
- [x] T028 [P] Integration test for resume PDF generation in backend/tests/integration/test_resume_generation.go

## Phase 3.3: Core Implementation (ONLY after tests are failing)

### Data Layer [P]
- [x] T029 [P] Generate UserProfile entity with ent in backend/src/models/user_profile.go
- [x] T030 [P] Generate ProfileSection entity with ent in backend/src/models/profile_section.go
- [x] T031 [P] Generate MediaAsset entity with ent in backend/src/models/media_asset.go
- [x] T032 [P] Generate ExternalLink entity with ent in backend/src/models/external_link.go
- [x] T033 [P] Generate ShareConfiguration entity with ent in backend/src/models/share_configuration.go
- [x] T034 [P] Generate ResumeTemplate entity with ent in backend/src/models/resume_template.go
- [x] T035 [P] Generate GeneratedResume entity with ent in backend/src/models/generated_resume.go
- [x] T036 Create database migration scripts for all entities in backend/migrations/

### Service Layer [P]
- [x] T037 [P] Implement ProfileService for profile CRUD operations in backend/src/services/profile_service.go
- [x] T038 [P] Implement SectionService for profile section management in backend/src/services/section_service.go
- [x] T039 [P] Implement MediaService for file upload and management in backend/src/services/media_service.go
- [x] T040 [P] Implement LinkService for external link management in backend/src/services/link_service.go
- [x] T041 [P] Implement ShareService for profile sharing functionality in backend/src/services/share_service.go
- [x] T042 [P] Implement ResumeService for PDF generation in backend/src/services/resume_service.go
- [x] T043 [P] Implement AuthService for OAuth 2.0 integration in backend/src/services/auth_service.go

### API Layer
- [x] T044 Implement OAuth login handler in backend/src/api/auth_handler.go
- [x] T045 Implement OAuth callback handler in backend/src/api/auth_callback_handler.go
- [x] T046 Implement GET /api/profiles endpoint in backend/src/api/profile_handler.go
- [x] T047 Implement POST /api/profiles endpoint in backend/src/api/profile_handler.go
- [x] T048 Implement GET /api/profiles/{profileId} endpoint in backend/src/api/profile_handler.go
- [x] T049 Implement PUT /api/profiles/{profileId} endpoint in backend/src/api/profile_handler.go
- [x] T050 Implement DELETE /api/profiles/{profileId} endpoint in backend/src/api/profile_handler.go
- [x] T051 Implement GET /api/profiles/{profileId}/sections endpoint in backend/src/api/section_handler.go
- [x] T052 Implement POST /api/profiles/{profileId}/sections endpoint in backend/src/api/section_handler.go
- [x] T053 Implement PUT /api/profiles/{profileId}/sections/{sectionId} endpoint in backend/src/api/section_handler.go
- [x] T054 Implement DELETE /api/profiles/{profileId}/sections/{sectionId} endpoint in backend/src/api/section_handler.go
- [x] T055 Implement POST /api/profiles/{profileId}/resume endpoint in backend/src/api/resume_handler.go

### CLI Layer [P]
- [ ] T056 [P] Implement profile creation CLI command in backend/src/cli/profile_create.go
- [ ] T057 [P] Implement profile section management CLI commands in backend/src/cli/section_manage.go
- [ ] T058 [P] Implement resume generation CLI command in backend/src/cli/resume_generate.go

### UI Layer
- [ ] T059 Create templ components for profile display in backend/src/ui/components/
- [ ] T060 Create templ pages for profile management in backend/src/ui/pages/
- [ ] T061 Implement HTMX interactions for dynamic UI updates in backend/src/ui/htmx/
- [ ] T062 Create server-side rendered handlers for UI pages in backend/src/ui/handlers/

## Phase 3.4: Integration
- [ ] T063 Connect ent ORM to SQLite database in backend/src/database/
- [ ] T064 Integrate NATS JetStream for work queue in backend/src/nats/queue.go
- [ ] T065 Integrate NATS Object Storage for file storage in backend/src/nats/storage.go
- [ ] T066 Implement OAuth 2.0 middleware for authentication in backend/src/middleware/auth.go
- [ ] T067 Implement authorization middleware for access control in backend/src/middleware/access.go
- [ ] T068 Set up LaTeX PDF generation service in backend/src/pdf/
- [ ] T069 Implement structured logging with zap in backend/src/telemetry/logging.go
- [ ] T070 Implement OpenTelemetry metrics and tracing in backend/src/telemetry/metrics.go
- [ ] T071 Add input validation for all API endpoints in backend/src/validation/
- [ ] T072 Implement error handling and standardized error responses in backend/src/errors/

## Phase 3.5: Polish
- [ ] T073 [P] Unit tests for ProfileService in backend/tests/unit/profile_service_test.go
- [ ] T074 [P] Unit tests for SectionService in backend/tests/unit/section_service_test.go
- [ ] T075 [P] Unit tests for MediaService in backend/tests/unit/media_service_test.go
- [ ] T076 [P] Unit tests for AuthService in backend/tests/unit/auth_service_test.go
- [ ] T077 Performance tests for profile operations (<200ms) in backend/tests/performance/profile_test.go
- [ ] T078 Performance tests for resume generation (<5s) in backend/tests/performance/resume_test.go
- [ ] T079 Security audit of OAuth 2.0 implementation in backend/docs/security.md
- [ ] T080 Security audit of access controls in backend/docs/security.md
- [ ] T081 Observability validation (logs, metrics, traces) in backend/tests/observability/
- [ ] T082 [P] Update API documentation from OpenAPI spec in backend/docs/api.md
- [ ] T083 [P] Update user documentation and quickstart guide in backend/docs/user-guide.md
- [ ] T084 Remove code duplication and refactor as needed
- [ ] T085 Run manual testing based on quickstart guide in backend/tests/manual/

## Dependencies
- Tests (T009-T028) before implementation (T029-T072)
- Data layer (T029-T036) before service layer (T037-T043)
- Service layer (T037-T043) before API layer (T044-T055)
- Core implementation (T029-T062) before integration (T063-T072)
- Implementation before polish (T073-T085)
- T044-T055 depend on T037-T043
- T059-T062 depend on T029-T036

## Parallel Example
```
# Launch T009-T020 together (contract tests):
Task: "Contract test for /auth/login endpoint in backend/tests/contract/test_auth_login.go"
Task: "Contract test for /auth/callback endpoint in backend/tests/contract/test_auth_callback.go"
Task: "Contract test for GET /api/profiles endpoint in backend/tests/contract/test_get_profiles.go"
Task: "Contract test for POST /api/profiles endpoint in backend/tests/contract/test_create_profile.go"
Task: "Contract test for GET /api/profiles/{profileId} endpoint in backend/tests/contract/test_get_profile.go"
Task: "Contract test for PUT /api/profiles/{profileId} endpoint in backend/tests/contract/test_update_profile.go"
Task: "Contract test for DELETE /api/profiles/{profileId} endpoint in backend/tests/contract/test_delete_profile.go"
Task: "Contract test for GET /api/profiles/{profileId}/sections endpoint in backend/tests/contract/test_get_sections.go"
Task: "Contract test for POST /api/profiles/{profileId}/sections endpoint in backend/tests/contract/test_create_section.go"
Task: "Contract test for PUT /api/profiles/{profileId}/sections/{sectionId} endpoint in backend/tests/contract/test_update_section.go"
Task: "Contract test for DELETE /api/profiles/{profileId}/sections/{sectionId} endpoint in backend/tests/contract/test_delete_section.go"
Task: "Contract test for POST /api/profiles/{profileId}/resume endpoint in backend/tests/contract/test_generate_resume.go"

# Launch T021-T028 together (integration tests):
Task: "Integration test for OAuth 2.0 flow with Google provider in backend/tests/integration/test_oauth_google.go"
Task: "Integration test for OAuth 2.0 flow with GitHub provider in backend/tests/integration/test_oauth_github.go"
Task: "Integration test for profile creation and management in backend/tests/integration/test_profile_management.go"
Task: "Integration test for profile section management in backend/tests/integration/test_section_management.go"
Task: "Integration test for media upload and storage in backend/tests/integration/test_media_upload.go"
Task: "Integration test for external link management in backend/tests/integration/test_link_management.go"
Task: "Integration test for profile sharing functionality in backend/tests/integration/test_profile_sharing.go"
Task: "Integration test for resume PDF generation in backend/tests/integration/test_resume_generation.go"

# Launch T029-T036 together (data layer):
Task: "Generate UserProfile entity with ent in backend/src/models/user_profile.go"
Task: "Generate ProfileSection entity with ent in backend/src/models/profile_section.go"
Task: "Generate MediaAsset entity with ent in backend/src/models/media_asset.go"
Task: "Generate ExternalLink entity with ent in backend/src/models/external_link.go"
Task: "Generate ShareConfiguration entity with ent in backend/src/models/share_configuration.go"
Task: "Generate ResumeTemplate entity with ent in backend/src/models/resume_template.go"
Task: "Generate GeneratedResume entity with ent in backend/src/models/generated_resume.go"

# Launch T037-T043 together (service layer):
Task: "Implement ProfileService for profile CRUD operations in backend/src/services/profile_service.go"
Task: "Implement SectionService for profile section management in backend/src/services/section_service.go"
Task: "Implement MediaService for file upload and management in backend/src/services/media_service.go"
Task: "Implement LinkService for external link management in backend/src/services/link_service.go"
Task: "Implement ShareService for profile sharing functionality in backend/src/services/share_service.go"
Task: "Implement ResumeService for PDF generation in backend/src/services/resume_service.go"
Task: "Implement AuthService for OAuth 2.0 integration in backend/src/services/auth_service.go"
```

## Notes
- [P] tasks = different files, no dependencies
- Verify tests fail before implementing
- Commit after each task
- Avoid: vague tasks, same file conflicts

## Task Generation Rules
*Applied during main() execution*

1. **From Contracts**:
   - Each contract endpoint → contract test task [P]
   - Each endpoint → controller implementation task
   
2. **From Data Model**:
   - Each entity → model creation task [P]
   - Relationships → service layer tasks
   
3. **From User Stories**:
   - Each story → integration test [P]
   - Quickstart scenarios → validation tasks

4. **Layered Architecture Compliance**:
   - Repository tasks for data persistence
   - Service tasks for business logic
   - Controller tasks for API endpoints
   
5. **Security Requirements**:
   - Auth middleware tasks for access control
   - Encryption tasks for data transmission
   
6. **Observability Requirements**:
   - Logging tasks for structured telemetry
   - Metrics tasks for performance tracking

7. **Ordering**:
   - Setup → Tests → Models → Services → Controllers → Integration → Polish
   - Dependencies block parallel execution
   - Tests before implementation (TDD)

## Validation Checklist
*GATE: Checked by main() before returning*

- [x] All contracts have corresponding tests
- [x] All entities have model tasks
- [x] All tests come before implementation
- [x] Parallel tasks truly independent
- [x] Each task specifies exact file path
- [x] No task modifies same file as another [P] task
- [x] API endpoints implemented before UI components
- [x] Security controls implemented for all sensitive operations
- [x] Structured logging implemented for all key operations
- [x] Implementation follows repository → service → controller pattern