# TD-001: Go HTTP API and SQL-owned persistence

- **Status:** Accepted
- **Date:** 2026-11-03
- **Owner:** Product/Engineering owner

## Context

Organization Onboarding is the first feature of a greenfield attendance API. The owner approved Go without an HTTP framework, PostgreSQL, explicit SQL through pgx, and Flyway. Database changes must be versioned SQL and must not be generated or applied automatically by an ORM.

## Options considered

### Option A — Go standard HTTP library with explicit SQL and Flyway

- Benefits: explicit request handling and SQL; existing owner familiarity with Flyway; independent schema lifecycle.
- Costs and risks: the application must establish consistent HTTP validation, errors, middleware, and operational configuration. Flyway is a separately managed tool.
- Onboarding implications: persistence and invitation state transitions can be designed explicitly without introducing an HTTP framework or ORM.

### Option B — HTTP framework, ORM, or a Go-native migration runner

- Benefits: additional conventions and helpers; Goose or golang-migrate can execute SQL migrations using Go tooling.
- Costs and risks: additional conventions and dependencies; changing migration tools is not necessary merely because the application uses Go.
- Onboarding implications: no currently identified requirement justifies replacing the approved standard-library and Flyway baseline. ORM-owned schema creation would violate the owner's constraint.

## Decision

- Implement the API in Go using `net/http`, without an HTTP framework.
- Use PostgreSQL and explicit SQL through `pgx`; an ORM is not part of the baseline.
- Manage schema changes exclusively through versioned SQL migrations executed by Flyway, independently of request handling. Do not create or alter tables automatically from application models or an ORM.
- Keep SQL migrations under `db/migrations/`. Do not introduce a second migration manager for the same schema.
- Preserve the API-generation versus HTTP-path distinction: this decision does not require a `/v2` URL prefix.
- Follow TD-002 for application organization and API documentation.

## Consequences

- Positive consequences: a small runtime stack, explicit persistence behavior, and reviewable schema changes independent of Go models.
- Trade-offs: application-level conventions must be maintained; PostgreSQL and Flyway must be available in the eventual development and delivery environment.
- Follow-up work: select supported tool versions, complete the administrative authentication/account-verification design, choose email delivery and specific hosting/deployment infrastructure, and define the detailed persistence/concurrency design. The operating model is recorded in [TD-003](003-delivery-model-and-customer-identity.md), and Supabase Auth is explicitly approved in [TD-004](004-initial-administrative-authentication.md). Other providers and infrastructure are not implicitly approved.
- This record approves the foundation, not application scaffolding, dependency installation, schema implementation, or deployment.
