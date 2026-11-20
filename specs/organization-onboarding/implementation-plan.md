# Organization Onboarding — Implementation and Test Plan

**Status:** Approved by the owner for local implementation within the incremental delivery boundary. Not yet implemented. Remaining design/security gates and separate operational authorization still apply.
**Approval evidence:** “apruebo la espec, y el plan para autorizar la implementacion”.
**Baseline:** Initial documentation commit `0f396ee` on `develop`. [Behavior specification](spec.md) supplies acceptance IDs; G-01 is resolved by owner approval; remaining gates stay explicit.

## Recommendation

This is an approved temporary execution plan, not a policy source. Follow the [documentation lifecycle](../../docs/agent-workflow.md#plan-closeout-and-archive): preserve outcomes/evidence and move durable information to its owner before optional archival after completion. This active, unimplemented plan was approved separately by the owner, not by the documentation reorganization.

Implement the full approved onboarding behavior incrementally, beginning with protected organization creation without invitation issuance. This first slice avoids coupling initial persistence to email delivery while preserving the independent creation flow. The owner approved this incremental delivery boundary.

Complete invitation lifecycle and verified-email/MFA acceptance before calling the complete onboarding feature delivered. Local doubles may unblock isolated tests but do not prove Supabase/Resend behavior or make provider-facing delivery complete.

## Approved constraints

- [TD-001](../../docs/technical-decisions/001-api-foundation.md): Go, `net/http`, PostgreSQL, explicit SQL through `pgx`, and Flyway-only versioned migrations under `db/migrations/`. No ORM-owned schema or assumed `/v2` prefix.
- [TD-002](../../docs/technical-decisions/002-application-structure-and-api-contract.md): feature-owned Controller/Service/Repository interfaces, concrete `internal/postgres` adapters, explicit `internal/app` wiring, and `api/openapi.yaml` using OpenAPI 3.1.x with Swagger UI.
- [TD-004](../../docs/technical-decisions/004-initial-administrative-authentication.md) and [TD-005](../../docs/technical-decisions/005-initial-platform-administrator-bootstrap.md): provider-backed identity/MFA assurance, application authorization, and separate restricted bootstrap.
- [TD-006](../../docs/technical-decisions/006-structured-application-logging.md) and [TD-007](../../docs/technical-decisions/007-transactional-email-delivery.md): `log/slog` JSON diagnostics and Resend Free with separate authentication/application email integrations.
- No stack reselection, empty scaffolding, default privileged account, automatic migration-on-request, remote configuration, purchases, dependency installation, commits beyond those explicitly authorized, or publishing occurs during this planning step.

## Proposed incremental deliveries

| Slice       | Coherent outcome with code, contract, tests, and documentation together                                                                                                                          | Criteria and release gates                                                                                                                                                                  |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| P1          | Minimal local API/persistence foundation and protected organization creation with deferred invitation/configuration; reject missing data and unauthorized access without business-state changes. | OB-001–003, OB-007–009, OB-013, OB-034. G-01; trusted identity/role test setup and verification design required. No externally usable admin bypass.                                         |
| P2          | Restricted initial platform bootstrap and administrative enforcement needed to operate the first slice, including applicable password/session guarantees.                                        | OB-029–031, OB-035; G-04 and provider-enforcement evidence before real use. P1 and P2 may develop against isolated fixtures, but protected onboarding is not operationally ready before P2. |
| P3          | Approved pre-acceptance configuration and invitation issuance/revocation/replacement, including durable replacement failure/order semantics.                                                     | OB-004–006, OB-010, OB-017–023; G-02 and transition/transaction design. Configuration does not implement attendance.                                                                        |
| P4          | Existing/new-account acceptance with trusted verified email and TOTP, atomic role assignment/consumption, tenant isolation, race and expiry cases.                                               | OB-011–016, OB-019–021; provider verification and identity-mapping design.                                                                                                                  |
| P5          | Authentication email flows and application invitation delivery, privacy, rate limits, quotas, eligible manual retry, approved sender/language, and integration proof.                            | OB-024–028, OB-032–034; G-05 and authorized provider integration. Email adapters used in earlier slices remain explicitly partial until verified here.                                      |
| Conditional | Actual manual lost-factor recovery execution, only if the owner includes it and approves evidence/procedure/audit policy.                                                                        | OB-036; G-03–05. Otherwise defer execution, retain documented policy, expose no unsafe reset.                                                                                               |

Each slice is a reviewable behavior unit, not a file-type batch or a promise that one slice fits one commit. Tests and documentation accompany its behavior. The owner authorized local implementation under this plan. Define reviewable work units and resolve applicable design gates before source changes.

## Proposed API and persistence design

### Contract first, without selecting unapproved public routes

Describe organization creation/configuration, invitation issuance/revocation/replacement/resend, and acceptance as explicit operations in OpenAPI. Define per-operation authentication/MFA and authorization, request/response examples, errors, and state effects before implementing it. Exact paths, status codes, naming, and idempotency behavior are proposals for the contract review, not silently approved by this plan.

Separate invitation references supplied by clients from trusted organization authorization. Public/authenticated invitee operations must not become unguarded platform operations. Do not expose tokens, raw provider errors, or account-existence differences through API responses or logs.

### Proposed records and invariants

| Record or capability                   | Required invariant; concrete schema is proposed                                                                                                                                                                |
| -------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Organization                           | Name and intended initial administrator email; configuration can remain unset. Retain organization even if later invitation delivery fails.                                                                    |
| Account identity mapping               | Stable provider subject identifies an administrative identity; verified email is evidence for invitation eligibility, not the permanent authority key by itself. Constraints and normalization require design. |
| Organization administrative assignment | Organization-scoped association; prevent duplicate grants while allowing multiple administrators and an account's legitimate assignments in different organizations.                                           |
| Invitation                             | Organization/email binding, expiry origin, protected token representation, and explicit consumption/replacement/revocation evidence. No plaintext invitation secrets in logs.                                  |
| Replacement coordination               | Durable indication that replacement started for an organization/email; serialization shared by acceptance and replacement. Must survive successor-creation failure.                                            |
| Session and recovery enforcement       | Trusted assurance, clocks, activity/IP evidence and consistent limit counters. Storage and failure semantics must be specified, not inferred from provider defaults.                                           |

Create SQL constraints and migrations only after the schema proposal is reviewed. Do not couple a real user or privileged identity to migration history. Use named feature-owned repository operations representing invariants rather than having controllers sequence unrelated SQL calls.

### Critical invitation transactions

**Recommended candidate:** serialize acceptance and replacement on an organization/email coordination record. Use a durable generation or replacement marker; the exact mechanism is proposed.

1. In a first committed transition, record replacement start and invalidate predecessors for that combination. Preserve previously completed acceptances.
2. In a subsequent operation, create the successor. If this fails, predecessor invalidation remains durable and the failure is visible; only explicit retry starts another attempt.
3. Acceptance serializes on the same coordination key, checks original expiry and all states, and atomically commits administrative assignment with consumption.
4. Perform delivery after invitation creation is durable. Report delivery failure separately; do not undo organization/invitation creation or automatically reset expiry.

A durable operation ledger is an alternative. **Reject a single rollback boundary that restores predecessor usability when successor creation fails.** Define the exact replacement-start linearization point, crash recovery, and replay behavior before implementing this path.

Do not apply the invitation failure rule to authentication-link replacement: successful Auth issuance replaces its predecessors, but failed Auth issuance preserves only original predecessor eligibility. Verify the hosted provider's actual behavior before claiming that requirement is met.

## Security and integration design gates

| Area                       | Engineering proposal or missing approval                                                                                                                                                                                                                                              |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Identity and tenant guards | Verify trusted provider subject, token validity and MFA assurance, then application role/organization scope. No client-supplied assertion, debug flag, or public account creation assigns authority.                                                                                  |
| Session policies           | Prove current-session logout, 30-minute inactivity and 8-hour absolute expiry, fresh password/TOTP after expiry, and trusted classification of explicit versus automatic activity. Token refresh is not sufficient evidence.                                                          |
| Password policy            | Verify provider/configuration behavior, maximum length and Unicode handling. Select a common/compromised-password blocking mechanism compatible with the approved budget; do not assume a paid provider feature is available.                                                         |
| Restricted bootstrap       | Present exact command permissions/environment, immutable account identification and verified-email/MFA evidence, no-secret errors/audit, and safe repeat behavior for owner approval before implementation. Tests use isolated identities, never real privileged users in migrations. |
| Request protections        | Specify email normalization, trusted proxy/IP handling, counter windows and atomic updates, failures counted, HTTP/timing behavior, and additional policies for exposed non-recovery flows. Existing recovery limits remain unchanged.                                                |
| Auth-link replacement      | Define successful issuance versus failed/unknown outcomes without treating message submission as delivery. Prove account/purpose isolation and prior-token invalidation with provider-backed integration evidence.                                                                    |
| Email quotas               | Plan coordinated accounting across Auth SMTP and API delivery; provider quota is shared and account usage is uninspected. No automatic quota-rejected retry, expired-link delayed send, or paid upgrade.                                                                              |
| Operational visibility     | Propose a minimal non-email way for the Platform Administrator to inspect quota failures. No new dashboard, Grafana service, or secret-bearing log dump is assumed.                                                                                                                   |
| Audit                      | If a recovery operation is delivered, obtain applicable access, persistence, retention and audit-write-failure decisions. Diagnostic `slog` output is not an audit store by default.                                                                                                  |
| Documentation exposure     | Swagger UI serves the same OpenAPI contract; define exposure and interactive-operation protection before publishing, not public access by default.                                                                                                                                    |

## Test-first and verification plan

There is currently no runnable application/test harness established by this documentation. No RED/GREEN, migration, provider, build, or functional-test result is claimed here.

1. Select exact supported tool versions, Go module path, PostgreSQL/Flyway test execution, and OpenAPI validation/contract tools in the implementation proposal. Do not install or pin them merely to draft this plan.
2. For each deterministic behavior with a runnable harness, observe a failing test before implementation, then GREEN and refactor. Assert output/errors, exit/status where applicable, and persisted side effects through the public boundary.
3. Keep Go tests beside their packages. Run focused tests for the slice and the applicable suite; capture actual commands and results. Candidate command forms such as `go test ./...` become executable requirements only once the module/harness exists.
4. Run PostgreSQL integration tests against isolated scratch databases outside human data, initialized by the same Flyway SQL migrations. Validate persisted states after missing input, unauthorized access, expired/consumed invitations, creation/delivery failures, and races.
5. Coordinate concurrent acceptance/replacement tests deterministically; use injected clocks for the 15-minute, 30-minute, 8-hour, rolling-hour and UTC-midnight boundaries. Test just before, at, and after applicable expiry without real waiting.
6. Test existing/new accounts, wrong email, unverified email, password-only sessions, forged claims, cross-tenant identifiers, multiple administrators, role-grant uniqueness, and no attendance record creation.
7. Test count/normalization/failure semantics only after they are explicitly designed; do not mistake unchanged business data for unchanged security-counter state.
8. Check OpenAPI validity and response/authentication conformance for each operation. Verify Swagger uses the same contract; it is not evidence of runtime validation by itself.
9. Use local provider doubles for success/error/unknown transport cases and secret-redaction tests. Before real integration completion, verify actual supported provider behavior in a separately authorized environment; never infer it from doubles or upstream source alone.
10. Keep scratch state outside the worktree. No live SMTP/DNS/account operation, spending, or production fixture is authorized by this test plan.

## Required completion evidence

- Criterion-to-test traceability for each included OB ID, including every named alternate/failure case.
- Reviewed OpenAPI, Flyway SQL, role/tenant constraints, and replacement linearization/commit boundaries.
- Observed focused/full test, contract, migration, and security/integration results; failed, unavailable, or mocked-only evidence explicitly marked.
- Approved bootstrap and operational access before operating onboarding; all included security gates resolved without an MFA or tenant bypass.
- No unrelated changes, secrets, automatic retries/upgrades, implicit attendance scope, or claim that a local test proves remote configuration.

## Separate operational authorization

Before any real Supabase, Resend, DNS, hosting, or privileged bootstrap action, obtain explicit authorization naming the destination, operation, and credential/session to use. Never discover or reuse ambient authenticated sessions to fill that gap. Domain verification, SMTP/API credentials, safe redirect targets, email templates, capacity checks, and operational logging are distinct from permission to write local code.

## Review decision

The owner approved the specification and delivery boundary and authorized local implementation. G-01 is resolved. The approved scope retains pre-acceptance configuration only (G-02) and defers lost-factor recovery execution until its separate approvals (G-03). Review concrete contract/persistence/security proposals and resolve applicable G-04–G-05 policies before their implementation slices. G-06 engineering work should be proposed coherently rather than turning each table, helper, or file into a separate product question. No application code, dependency, schema, service, or additional commit has been created by this document.
