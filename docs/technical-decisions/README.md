# Technical Decisions

This folder records technical decisions that have been discussed and explicitly
approved for Attendance API v2.

API v2 names the implementation generation. It does not establish an HTTP path, a `/v2` route prefix, or an approved API versioning strategy.

## Accepted decisions

| Record | Scope |
| --- | --- |
| [TD-001: Go HTTP API and SQL-owned persistence](001-api-foundation.md) | Go, `net/http`, PostgreSQL, `pgx`, and Flyway SQL migrations. |
| [TD-002: Feature-oriented layers and an OpenAPI contract](002-application-structure-and-api-contract.md) | Controller/Service/Repository responsibilities, concrete PostgreSQL adapters, feature packages, OpenAPI, and Swagger UI. |
| [TD-003: Owner-hosted initial delivery and deferred customer installations](003-delivery-model-and-customer-identity.md) | Initial owner-operated platform; customer-operated delivery with enterprise identity is documented for later evaluation, not initial implementation. |
| [TD-004: Supabase Auth for initial administrative authentication](004-initial-administrative-authentication.md) | Initial administrative accounts use Supabase Auth; organizations, permissions, invitations, and attendance remain application-owned. |
| [TD-005: Controlled initial Platform Administrator bootstrap](005-initial-platform-administrator-bootstrap.md) | Verified account and completed TOTP MFA, followed by an owner-invoked restricted command assigning the initial platform role; no public or migration-seeded privileged account. |
| [TD-006: Structured application logging with Go log/slog](006-structured-application-logging.md) | Standard-library logger with JSON output, approved fields, sensitive-data exclusions, and separate security audit requirements. |
| [TD-007: Managed transactional email with separate authentication and invitation flows](007-transactional-email-delivery.md) | Resend with `notificaciones.allangallo.dev`, custom SMTP for Supabase Auth, and a separate API integration for application-owned invitations; configuration and contracting pending. |

These TD records already serve as ADRs: they preserve approved choices, alternatives, rationale, and consequences. Keep their IDs and paths. Current requirements are owned by the [administrative access policy](../security/administrative-access.md) and [email delivery policy](../operations/email-delivery.md), not duplicated in this index. Those pages separate confirmed behavior from remaining enforcement, configuration, and operational work. TD-005 continues to own the restricted initial bootstrap decision. Accepted foundations do not authorize implementation, purchases, remote provisioning, or deployment by themselves.

## Approved logging baseline and pending observability design

See [TD-006](006-structured-application-logging.md) for the approved logger and diagnostics boundary; do not copy its complete field/security rules into this navigation index.

Grafana remains a candidate for evaluation. No collection pipeline, storage backend, deployment, or paid service has been selected or authorized. Evaluate application log generation separately from log collection/storage and Grafana-based querying/visualization before implementation. Recovery audit requirements in TD-004 remain distinct; choosing a monitoring tool does not settle audit storage, access, retention, or failure handling.

## Recording decisions

A proposal from an AI agent is not an approved decision. When evaluating a proposal, record:

- **Status:** Proposed, Accepted, Deferred, or Superseded.
- **Context:** What problem requires a decision?
- **Options considered:** Include material alternatives.
- **Decision:** The approved choice, or why the decision is deferred.
- **Consequences:** Benefits, costs, constraints, and migration implications.

Do not create an ADR for every small implementation detail. Record durable choices that affect future work, security, integration, operations, or maintainability.

At the start of this project, no technical choices are assumed to be approved solely because an agent recommended them.
