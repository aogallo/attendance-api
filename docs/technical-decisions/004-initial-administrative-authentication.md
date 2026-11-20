# TD-004: Supabase Auth for initial administrative authentication

- **Status:** Accepted
- **Date:** 2026-11-03
- **Owner:** Product/Engineering owner

## Context

The initial platform is hosted by the owner. Ordinary people record attendance through facial recognition in a shared-device attendance mode and do not need login accounts. Administrators require accounts, and organization invitations must be accepted by the authenticated identity corresponding to the invited email.

The owner approved Supabase Auth as the initial administrative authentication provider. This is a design decision, not authorization to create a hosted project, transfer data, install dependencies, or purchase a plan.

## Options considered

### Option A — Supabase Auth with application-owned authorization

- Benefits: a managed account-authentication service, a free development starting point, and no requirement to operate an identity server initially.
- Costs and risks: provider dependency, plan and usage limits, email-delivery configuration, and ongoing integration/verification responsibilities. Enterprise SSO has different pricing and is not part of the initial scope.
- Onboarding implications: use the provider for account authentication while preserving the application's invitation and organization-permission rules.

### Option B — Operate Keycloak or build authentication in the Go application

- Benefits: control over infrastructure and account behavior; Keycloak supports enterprise identity integration.
- Costs and risks: operating a separate identity service or maintaining account-security functionality, recovery, and session handling in application code.
- Onboarding implications: viable alternatives, but not selected for the initial delivery.

## Decision

- Use Supabase Auth for administrative login accounts in the initial owner-hosted platform, covering Platform Administrators and Organization Administrators.
- Email/password with TOTP MFA and provider-owned credential recovery support application-owned administrative authorization. Maintain the detailed, confirmed requirements in the [current administrative access policy](../security/administrative-access.md), not as a second mutable rulebook in this ADR.
- Keep organizations, administrative assignments, permissions, invitation records and lifecycle rules, Person data, schedules, and attendance in the application domain. Do not require one Supabase project or native provider organization for each customer organization.
- Organization invitation lifecycle remains application-owned under [confirmed product decisions](../product-decisions.md#12-administrator-invitations-expire-are-single-use-and-are-replaced-manually); provider authentication must not replace it.
- TD-001 continues to govern the application-owned PostgreSQL schema and Flyway migrations. Do not directly manage the provider's internal authentication tables as application tables or assume that business data must be moved to the provider's database.
- Keep provider-specific authentication integration outside business rules, consistent with TD-002 and the deferred customer-operated scenario in TD-003. This does not require implementing multiple providers now.
- Begin development using the free plan if its current limits and terms fit the environment. Production plan selection, usage limits, costs, infrastructure, and email delivery require separate assessment and explicit authorization before spending or provisioning. Do not treat a pricing snapshot as a permanent guarantee.

### Email delivery responsibilities

[TD-007](007-transactional-email-delivery.md) records the integration choice. The [current email delivery policy](../operations/email-delivery.md) owns sender values, capacity caveats, and failure handling; it does not authorize real configuration or sending.

### Authentication email-link expiry

[Current policy](../security/administrative-access.md#authentication-email-link-expiry). This heading is retained for existing links; maintain the requirements at that destination.

### Password-recovery request privacy

[Current policy](../security/administrative-access.md#password-recovery-request-privacy).

### Password-recovery request limits

[Current policy](../security/administrative-access.md#password-recovery-request-limits).

### Administrative password policy

[Current policy](../security/administrative-access.md#administrative-password-policy).

### Sign-in before MFA completion

[Current policy](../security/administrative-access.md#sign-in-before-mfa-completion).

### Ordinary sign-out scope

[Current policy](../security/administrative-access.md#ordinary-sign-out-scope).

### Administrative session limits

[Current policy](../security/administrative-access.md#administrative-session-limits).

### Lost authenticator recovery

[Current policy](../security/administrative-access.md#lost-authenticator-recovery).

## Consequences

- Positive consequences: administrative authentication can use a managed service without making every attendance Person a billable login identity or operating Keycloak initially.
- Trade-offs: availability, pricing, migration portability, and session/account behavior depend partly on the provider. Changing providers later is not guaranteed to be configuration-only.
- Follow-up work: define enforcement and remaining details of the approved password policy, detailed email-confirmation/recovery handling, MFA enrollment controls, restricted recovery procedure controls and detailed recovery evidence/session controls, session/token validation and renewal, the bootstrap command controls under TD-005, and email delivery. Email/password sign-in, verified invited email and completed TOTP MFA before invitation acceptance, email-link password recovery, mandatory TOTP MFA for both administrative roles, manual identity-verified lost-factor recovery with new-factor enrollment, exclusive Platform Administrator authorization of Organization Administrator recovery, and owner/operator responsibility for restricted Platform Administrator recovery are approved; their configuration and error handling still need design. Shared-device capture-session lifetime and return-to-administration challenge/recency details remain separately defined attendance work and must preserve the administrative MFA requirement.
- Provider configuration and secret handling must be defined before integration; do not expose backend credentials to browsers or shared devices.
- No application implementation, remote account/project creation, purchase, deployment, or schema change is performed or authorized by this record alone.

## References

- [Supabase production checklist: shared authentication email-link expiry](https://supabase.com/docs/guides/security/production-checklist).
- [NIST SP 800-63B-4: Password verifiers](https://pages.nist.gov/800-63-4/sp800-63b.html#passwordver).
- [Supabase password security](https://supabase.com/docs/guides/auth/password-security).
- [Supabase Auth documentation](https://supabase.com/docs/guides/auth).
- [Supabase pricing](https://supabase.com/pricing): recheck before choosing a plan.
- [Auth email/SMTP requirements](https://supabase.com/docs/guides/auth/auth-smtp).
- [TD-001](001-api-foundation.md), [TD-002](002-application-structure-and-api-contract.md), and [TD-003](003-delivery-model-and-customer-identity.md).
