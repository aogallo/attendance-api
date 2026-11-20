# TD-007: Managed transactional email with separate authentication and invitation flows

- **Status:** Accepted
- **Date:** 2026-11-03
- **Owner:** Product/Engineering owner

## Context

The initial platform needs administrative email verification, password recovery, and Organization Administrator invitations. Supabase Auth owns authentication flows; the application owns organization invitations and their lifecycle. The owner approved using one managed transactional email provider for both, with separate integrations, and subsequently selected Resend with the sending subdomain `notificaciones.allangallo.dev`. Contracting and operational configuration remain pending.

This records a delivery approach, not authorization to implement, configure services, send email, or incur costs.

## Options considered

### Option A — One managed transactional email provider with separate integrations

- Benefits: avoids operating a mail server while allowing one provider to deliver both categories of email.
- Costs and risks: provider dependency, quotas, deliverability, domain configuration, and operating costs.
- Selected with Resend Free as the initial plan; credentials and operational configuration are not approved here.

### Option B — Operate a mail server or rely on the default Auth email service in production

- Operating a mail server adds delivery and maintenance responsibilities not needed for the approved initial approach.
- Supabase's default Auth email service is intended for testing and has recipient/rate restrictions; it is not the selected production delivery solution.
- Not selected.

## Decision

- Use Resend as the initial managed transactional email provider for the platform's authentication and organization-invitation messages.
- Start with the Resend Free transactional plan, as approved on 2026-11-04. Paid upgrades or additional spending require separate explicit authorization; this plan selection does not authorize account access or configuration.
- Maintain sender/domain/language, verification prerequisites, and operating limits in the [current email delivery policy](../operations/email-delivery.md). This ADR preserves the provider/integration decision and its rationale, not a second mutable configuration inventory.
- Configure Supabase Auth to use that provider through custom SMTP for email verification and password recovery. Supabase remains responsible for those authentication flows.
- Authentication-link lifetime/replacement and recovery privacy/limits belong to the [administrative access policy](../security/administrative-access.md); application invitation lifecycle remains separately owned by [product decisions](../product-decisions.md#12-administrator-invitations-expire-are-single-use-and-are-replaced-manually).
- The API sends application-owned organization invitations through a separate integration with the same provider. The API integration's transport and implementation remain pending; custom SMTP for Auth does not automatically deliver application invitations.
- Do not replace application invitations with Supabase's native authentication invitation lifecycle. Preserve the original 15-minute validity from invitation-record creation, single-use acceptance, explicit replacement/revocation, and approved resend behavior.

### Free-plan capacity snapshot — 2026-11-04

[Dated capacity snapshot and caveats](../operations/email-delivery.md#free-plan-capacity-snapshot--2026-11-04). This heading remains for existing links; update observations in the policy page.

### Provider quota exhaustion

[Current quota-exhaustion policy](../operations/email-delivery.md#provider-quota-exhaustion).

## Consequences

- Authentication and organization invitations can share email delivery infrastructure without sharing ownership of their business rules or token lifecycles.
- Sending a message does not verify an email address, grant a role, complete MFA, or accept an invitation.
- This decision does not authorize purchases, provisioning, DNS changes, obtaining or using credentials, application implementation, or real email delivery.

## References

- [TD-004: Initial administrative authentication](004-initial-administrative-authentication.md).
- [Supabase custom SMTP](https://supabase.com/docs/guides/auth/auth-smtp).
- [Resend with Supabase SMTP](https://resend.com/docs/send-with-supabase-smtp).
- [Resend sending domains and subdomains](https://resend.com/docs/dashboard/domains/introduction).
- [Resend pricing](https://resend.com/pricing).
- [Resend account quotas and limits](https://resend.com/docs/knowledge-base/account-quotas-and-limits).
- [Confirmed product decisions](../product-decisions.md).
