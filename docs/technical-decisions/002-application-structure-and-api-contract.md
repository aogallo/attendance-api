# TD-002: Feature-oriented layers and an OpenAPI contract

- **Status:** Accepted
- **Date:** 2026-11-03
- **Owner:** Product/Engineering owner

## Context

The owner approved Controller, Service, and Repository responsibilities, a feature-oriented Go layout, and API documentation from the first endpoint using OpenAPI and Swagger UI. The structure must remain small and understandable without requiring SDD tooling.

## Options considered

### Option A — Feature-oriented packages and an explicit OpenAPI contract

- Benefits: related behavior stays together; persistence implementations remain replaceable; clients can review the HTTP contract independently of implementation.
- Costs and risks: dependencies between features must avoid cycles; contract and implementation need consistency checks.
- Onboarding implications: organization and invitation behavior have explicit service boundaries, with PostgreSQL details kept outside those packages.

### Option B — Global layer directories and documentation generated from controller comments

- Benefits: familiar layer-based navigation and documentation located near endpoint implementation.
- Costs and risks: changes span global directories as features grow; documentation becomes coupled to a particular annotation tool.
- Onboarding implications: viable, but not the organization selected by the owner.

## Decision

### Responsibilities and dependencies

- Controllers handle HTTP decoding, transport-level validation, service invocation, and response/error mapping. They contain no SQL or business rules.
- Services execute use cases, business rules, and coordination. They depend on small, use-case-specific Repository interfaces, not HTTP request/response types or concrete PostgreSQL repositories.
- Repository interfaces live in the feature package that consumes them, for example `internal/organization/repository.go` declares `organization.Repository`.
- Files such as `internal/postgres/organization_repository.go` contain concrete structs and methods implementing those interfaces using `pgx` and SQL. They are not interfaces, inherited base classes, or templates imported into feature repositories.
- Implementation satisfaction is implicit in Go: a concrete type implements an interface by providing its methods. Feature packages must not import `internal/postgres`; the PostgreSQL adapter may import feature types and interfaces without creating a cycle.
- `internal/app` constructs and connects concrete repositories, services, controllers, and routing. Use explicit construction rather than a dependency-injection framework.
- Do not create generic base repositories or an interface for every struct. Add abstractions when a consumer needs a behavioral contract.
- Transaction boundaries must preserve the approved business rules; do not scatter use-case orchestration into controllers or treat separate repository calls as automatically atomic.

### Layout

```text
cmd/api/main.go
internal/app/app.go
internal/httpapi/router.go
internal/httpapi/middleware.go
internal/organization/controller.go
internal/organization/service.go
internal/organization/repository.go
internal/organization/model.go
internal/invitation/controller.go
internal/invitation/service.go
internal/invitation/repository.go
internal/invitation/model.go
internal/postgres/pool.go
internal/postgres/organization_repository.go
internal/postgres/invitation_repository.go
api/openapi.yaml
db/migrations/
docs/
specs/
go.mod
go.sum
```

This is a placement convention, not a requirement to generate empty files or directories. Keep tests alongside the packages they exercise using `*_test.go`. Go logic resides under `internal/` and the executable entry point under `cmd/`; other directory and layer choices are project conventions, not an official universal Go layout.

### API contract and documentation

- Define the HTTP/JSON contract explicitly using OpenAPI 3.1.x in `api/openapi.yaml`, versioned with the implementation.
- Use Swagger UI to present the same contract; OpenAPI and Swagger UI are complementary, not competing sources of truth.
- Document request and response schemas, errors, authentication requirements, and examples from the first implemented endpoint.
- Validate the OpenAPI document and use contract checks to keep it consistent with implementation. Select the checking tools during implementation planning; do not assume generated documentation enforces runtime behavior.
- Do not require controller annotations, server code generation, or SDD tooling to contribute.
- Choose documentation exposure and interactive-operation protections before publishing. This decision does not make documentation public or authorize API operations without normal authentication and authorization.

## Consequences

- Positive consequences: discoverable feature boundaries, testable business behavior, isolated SQL adapters, and an independently reviewable API contract.
- Trade-offs: contract consistency needs checks; feature boundaries need care to avoid circular dependencies.
- Follow-up work: define the initial endpoint contract, concrete service/repository methods, transaction coordination, tool versions, and documentation asset delivery. Do not prebuild unused abstractions.

## References

- [Official Go module organization](https://go.dev/doc/modules/layout).
- [Community project-layout disclaimer](https://github.com/golang-standards/project-layout#overview): this is not an official Go-team standard.
- [Go interfaces and types](https://go.dev/doc/effective_go#interfaces_and_types).
- [Swagger UI](https://github.com/swagger-api/swagger-ui).
