# TD-006: Structured application logging with Go log/slog

- **Status:** Accepted
- **Date:** 2026-11-03
- **Owner:** Product/Engineering owner

## Context

The owner approved structured JSON application logs with date/time, severity level, operation, and request identifier, excluding passwords, tokens, MFA secrets, and biometric data. The owner subsequently approved Go's standard-library `log/slog` as the initial logger.

This is a design decision only. Application implementation and observability infrastructure are not authorized by this record.

## Options considered

### Option A — Standard-library log/slog

- Supports structured attributes, severity levels, and JSON output through `JSONHandler` without adding a third-party logging dependency.
- Selected for the approved initial logging baseline. Collection, storage, visualization, and security auditing remain separate responsibilities.

### Option B — Third-party logging library

- May provide additional features or specialized performance characteristics, but adds a dependency and integration choices.
- Not selected: no requirement established so far calls for capabilities beyond the approved standard-library baseline.

## Decision

- Use `log/slog` as the initial application logger, with `JSONHandler` for structured JSON output.
- Include date/time, severity level, operation, and request identifier in the approved application logging format. Detailed field names and request-identifier handling remain implementation-design work.
- Exclude passwords, tokens, MFA secrets, and biometric data. Selecting a logger does not automatically enforce those exclusions; implementation and tests must enforce them.
- Keep security audit requirements separate from diagnostic logs. The recovery audit requirements in [TD-004](004-initial-administrative-authentication.md) are not satisfied merely by selecting or calling this logger.
- Grafana remains a visualization/querying candidate, not the application logger. Neither Grafana nor Loki deployment, a collection pipeline, a storage backend, or a paid service is approved here.

## Consequences

- The application can produce structured logs without a third-party logging library.
- Output destination, collection, storage, retention, access controls, field conventions, request correlation, and operational configuration still require design.
- This decision does not install dependencies, implement application code, provision services, authorize spending, or deploy anything.

## References

- [Go log/slog documentation](https://pkg.go.dev/log/slog).
- [Confirmed product decisions](../product-decisions.md).
