# Implementation Plan: Minimal HTTP API with Health Check Endpoint

**Branch**: `001-before-building-any` | **Date**: 2025-09-21 | **Spec**: /specs/001-before-building-any/spec.md
**Input**: Feature specification from `/specs/001-before-building-any/spec.md`

## Execution Flow (/plan command scope)
```
1. Load feature spec from Input path
   → If not found: ERROR "No feature spec at {path}"
2. Fill Technical Context (scan for NEEDS CLARIFICATION)
   → Detect Project Type from context (web=frontend+backend, mobile=app+api)
   → Set Structure Decision based on project type
3. Fill the Constitution Check section based on the content of the constitution document.
4. Evaluate Constitution Check section below
   → If violations exist: Document in Complexity Tracking
   → If no justification possible: ERROR "Simplify approach first"
   → Update Progress Tracking: Initial Constitution Check
5. Execute Phase 0 → research.md
   → If NEEDS CLARIFICATION remain: ERROR "Resolve unknowns"
6. Execute Phase 1 → contracts, data-model.md, quickstart.md, agent-specific template file (e.g., `CLAUDE.md` for Claude Code, `.github/copilot-instructions.md` for GitHub Copilot, `GEMINI.md` for Gemini CLI, `QWEN.md` for Qwen Code or `AGENTS.md` for opencode).
7. Re-evaluate Constitution Check section
   → If new violations: Refactor design, return to Phase 1
   → Update Progress Tracking: Post-Design Constitution Check
8. Plan Phase 2 → Describe task generation approach (DO NOT create tasks.md)
9. STOP - Ready for /tasks command
```

**IMPORTANT**: The /plan command STOPS at step 7. Phases 2-4 are executed by other commands:
- Phase 2: /tasks command creates tasks.md
- Phase 3-4: Implementation execution (manual or via tools)

## Summary
Build a minimal HTTP API with only a healthz endpoint for bootstrapping the project. The endpoint must return a status code based on the service health with a JSON response. The user must be able to configure the port the HTTP server is listening to. The API project must be architected to be future proof for multi layer structure mentioned in constitution.

## Technical Context
**Language/Version**: Go 1.21 or later
**Primary Dependencies**: danielgtaylor/huma (API framework), uber/fx (dependency injection), opentelemetry (observability), zap (logging), hclsimple (configuration), cobra (CLI), ginkgo (testing)
**Storage**: N/A
**Testing**: ginkgo/gomega
**Target Platform**: Linux server
**Project Type**: single (backend API only)
**Performance Goals**: NEEDS CLARIFICATION
**Constraints**: NEEDS CLARIFICATION
**Scale/Scope**: Small, single service

## Constitution Check
*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

Based on VibeFolio Constitution v1.1.0:

1. **Focus on Correctness and Best Practices**: Implementation must prioritize correctness over speed; no assumptions should be made; documentation must be consulted; best practices must be researched.
2. **API First Development**: All functionality must be designed with API endpoints before UI implementation.
3. **Test-Driven Development**: All code must have corresponding tests; unit tests before implementation where feasible.
4. **Security by Default**: Access control and encryption must be considered for all features.
5. **Observability Through Telemetry**: All components must emit structured logs and metrics.
6. **Layered Architecture**: Implementation must follow repository → service → controller pattern.

Violations must be documented in Complexity Tracking with justification.

## Project Structure

### Documentation (this feature)
```
specs/001-before-building-any/
├── plan.md              # This file (/plan command output)
├── research.md          # Phase 0 output (/plan command)
├── data-model.md        # Phase 1 output (/plan command)
├── quickstart.md        # Phase 1 output (/plan command)
├── contracts/           # Phase 1 output (/plan command)
└── tasks.md             # Phase 2 output (/tasks command - NOT created by /plan)
```

### Source Code (repository root)
```
# Option 1: Single project (DEFAULT)
src/
├── models/
├── services/
├── cli/
└── lib/

tests/
├── contract/
├── integration/
└── unit/
```

**Structure Decision**: DEFAULT to Option 1 (single project) as this is a simple backend API with no frontend component.

## Phase 0: Outline & Research
1. **Extract unknowns from Technical Context** above:
   - Performance Goals: What are the expected performance requirements for the health check endpoint?
   - Constraints: What are the specific constraints for this service?

2. **Generate and dispatch research agents**:
   ```
   For each unknown in Technical Context:
     Task: "Research performance requirements for health check endpoints"
     Task: "Research deployment constraints for Go services"
   For each technology choice:
     Task: "Find best practices for danielgtaylor/huma in health check services"
     Task: "Find best practices for uber/fx dependency injection in Go"
     Task: "Find best practices for OpenTelemetry integration with zap logging"
     Task: "Find best practices for HCL configuration with hclsimple"
     Task: "Find best practices for cobra CLI applications"
     Task: "Find best practices for ginkgo testing framework"
   ```

3. **Consolidate findings** in `research.md` using format:
   - Decision: [what was chosen]
   - Rationale: [why chosen]
   - Alternatives considered: [what else evaluated]

**Output**: research.md with all NEEDS CLARIFICATION resolved

## Phase 1: Design & Contracts
*Prerequisites: research.md complete*

1. **Extract entities from feature spec** → `data-model.md`:
   - HealthStatus entity with fields for status, timestamp, and any relevant service information
   - Configuration entity for port and other configurable parameters

2. **Generate API contracts** from functional requirements:
   - GET /healthz endpoint that returns HealthStatus
   - Configuration contract for port setting
   - Output OpenAPI schema to `/contracts/`

3. **Generate contract tests** from contracts:
   - One test file per endpoint
   - Assert request/response schemas
   - Tests must fail (no implementation yet)

4. **Extract test scenarios** from user stories:
   - Each story → integration test scenario
   - Quickstart test = story validation steps

5. **Update agent file incrementally** (O(1) operation):
   - Run `.specify/scripts/bash/update-agent-context.sh qwen` for your AI assistant
   - If exists: Add only NEW tech from current plan
   - Preserve manual additions between markers
   - Update recent changes (keep last 3)
   - Keep under 150 lines for token efficiency
   - Output to repository root

**Output**: data-model.md, /contracts/*, failing tests, quickstart.md, agent-specific file

## Phase 2: Task Planning Approach
*This section describes what the /tasks command will do - DO NOT execute during /plan*

**Task Generation Strategy**:
- Load `.specify/templates/tasks-template.md` as base
- Generate tasks from Phase 1 design docs (contracts, data model, quickstart)
- Each contract → contract test task [P]
- Each entity → model creation task [P]
- Each user story → integration test task
- Implementation tasks to make tests pass

**Ordering Strategy**:
- TDD order: Tests before implementation
- Dependency order: Models before services before UI
- Mark [P] for parallel execution (independent files)

**Estimated Output**: 25-30 numbered, ordered tasks in tasks.md

**IMPORTANT**: This phase is executed by the /tasks command, NOT by /plan

## Phase 3+: Future Implementation
*These phases are beyond the scope of the /plan command*

**Phase 3**: Task execution (/tasks command creates tasks.md)
**Phase 4**: Implementation (execute tasks.md following constitutional principles)
**Phase 5**: Validation (run tests, execute quickstart.md, performance validation)

## Complexity Tracking
*Fill ONLY if Constitution Check has violations that must be justified*

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| Layered Architecture | Even for a simple health check service, following the layered architecture ensures future extensibility and maintains consistency with constitutional principles | A flat architecture would be simpler but would violate our constitutional requirement for layered architecture |

## Progress Tracking
*This checklist is updated during execution flow*

**Phase Status**:
- [x] Phase 0: Research complete (/plan command)
- [x] Phase 1: Design complete (/plan command)
- [x] Phase 2: Task planning complete (/plan command - describe approach only)
- [ ] Phase 3: Tasks generated (/tasks command)
- [ ] Phase 4: Implementation complete
- [ ] Phase 5: Validation passed

**Gate Status**:
- [x] Initial Constitution Check: PASS
- [x] Post-Design Constitution Check: PASS
- [x] All NEEDS CLARIFICATION resolved
- [x] Complexity deviations documented

---
*Based on Constitution v1.1.0 - See `/memory/constitution.md`*
