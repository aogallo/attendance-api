# TD-005: Controlled initial Platform Administrator bootstrap

- **Status:** Accepted
- **Date:** 2026-11-03
- **Owner:** Product/Engineering owner

## Context

The first Platform Administrator must exist before organization onboarding can be operated. Public registration must not let a user acquire this platform-wide role. Supabase Auth manages accounts and credentials, while the application owns administrative assignments and Flyway manages its schema.

The owner approved an explicit restricted bootstrap command to assign the first platform role to the owner's verified account after MFA setup. This records the design only; no command or account exists merely because this record is accepted.

## Options considered

### Option A — Restricted command after verified account and MFA setup

- Benefits: separates schema management, identity creation, and privilege assignment; avoids embedded passwords and public self-elevation.
- Costs and risks: requires an authorized operator, correct target-environment selection, and a carefully verified account identifier.
- Onboarding implications: initial platform setup is completed before its administrative endpoints are used.

### Option B — Create a privileged user in a migration, startup seed, or public registration flow

- Benefits: fewer explicit setup steps.
- Costs and risks: can couple a real identity to schema history, introduce default credentials, repeatedly grant privileges on startup, or expose self-elevation.
- Onboarding implications: not selected for provisioning the initial platform account.

## Decision

1. The owner creates or reuses the intended account in Supabase Auth.
2. The account's email is verified and TOTP MFA is configured and successfully completed before the initial platform-role assignment.
3. The owner explicitly invokes a restricted bootstrap command from an authorized operational environment.
4. The command identifies the exact intended provider account and records its Platform Administrator assignment in the application database.

The bootstrap must satisfy these boundaries:

- No public endpoint or registration form grants the Platform Administrator role.
- Keep the bootstrap command's implementation versioned, but do not put real passwords, tokens, backend secrets, or fixed privileged accounts into source code or SQL migrations.
- Flyway creates application tables and schema constraints. It may hold approved reference-role definitions, but not the real initial privileged account or its password. Creating an identity and assigning its role remain separate operations.
- Run the command deliberately for the intended environment, not automatically on every API startup. A repeated invocation must not silently reassign, duplicate, or regrant privileges; exact replay/refusal behavior remains to be specified.
- Do not directly insert accounts into the provider's internal authentication tables or store a duplicate password in the application database.
- The command does not create a Person, schedule, attendance-method enrollment, or attendance record for the Platform Administrator.
- Administrative operations still require MFA enforcement under TD-004. Bootstrap is not an approved password-reset, lost-factor recovery, or emergency MFA-bypass mechanism.
- Account creation/reuse, email verification, and MFA must not be marked complete merely because an operator supplied an email or set an unverified local flag. The detailed evidence checks and command permissions remain implementation-design work.

## Consequences

- Positive consequences: the initial role is explicitly assigned without placing credentials or a real privileged user in migration history.
- Trade-offs: initial setup requires an authorized manual action; environment and identity mistakes must be prevented and tested.
- Follow-up work: define the command interface, operational authorization and credential scope, account/MFA verification evidence, repeat-execution behavior, safe error reporting, and audit evidence before implementation. Later Platform Administrator appointments require a separate decision. TD-004 reserves Platform Administrator lost-authenticator recovery to the platform owner/operator through a separate restricted operational procedure, with identity verification and a successfully verified new TOTP factor before access returns. Its detailed controls remain pending; this bootstrap must not be reused for recovery.
- This decision does not authorize creating users, obtaining credentials, invoking privileged provider APIs, running migrations, or executing bootstrap in any local or remote environment now.

## Related decisions

- [TD-001: Go HTTP API and SQL-owned persistence](001-api-foundation.md).
- [TD-004: Supabase Auth for initial administrative authentication](004-initial-administrative-authentication.md).
