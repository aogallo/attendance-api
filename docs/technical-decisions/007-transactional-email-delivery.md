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
- Use neutral, professional Spanish as the initial language for organization-invitation, email-verification, and password-recovery emails. This approves their language only; exact templates, content, and delivery configuration remain pending. It does not select a localization mechanism or change the language of API contracts or technical documentation.
- Use `notificaciones.allangallo.dev` as the sending subdomain. This uses the owner's current domain and therefore exposes the owner's personal domain in sender addresses rather than a separate product domain.
- The subdomain must be added and verified for sending before use. The owner's earlier Resend setup with a different domain does not establish verification of this subdomain. No verification or DNS configuration has been performed as part of this design decision.
- Configure Supabase Auth to use that provider through custom SMTP for email verification and password recovery. Supabase remains responsible for those authentication flows.
- Email-verification and password-recovery links use the approved 30-minute validity in TD-004 (intended shared Supabase Email OTP Expiration: 1800 seconds). Application-owned organization invitations retain their independent 15-minute validity from record creation. These are design values; no provider setting has been changed.
- Under TD-004, successful issuance of a new authentication link invalidates earlier links for the same account and same purpose, with a new 30-minute validity from issuance. A rejected request or failed issuance does not replace an earlier link or extend its original eligibility/expiry. Other purposes and organization invitations remain unaffected. Successful issuance is not proof of delivery; provider-backed success and invalidation enforcement remain to be designed and verified.
- The API sends application-owned organization invitations through a separate integration with the same provider. The API integration's transport and implementation remain pending; custom SMTP for Auth does not automatically deliver application invitations.
- Do not replace application invitations with Supabase's native authentication invitation lifecycle. Preserve the original 15-minute validity from invitation-record creation, single-use acceptance, explicit replacement/revocation, and approved resend behavior.
- Delivery failure preserves an already-created organization and invitation, is clearly reported to the Platform Administrator, and allows manual resend of the same eligible invitation without extending expiry. An expired invitation requires explicit replacement.
- Use `no-reply@notificaciones.allangallo.dev` as the sender address and `Attendance` as the display name for the initial authentication and organization-invitation emails. These are approved design values, not evidence of a configured sender, verified domain, or working delivery.
- Credentials, account/plan verification, templates, detailed delivery integration, and operational safeguards remain pending. This sender identity decision does not define an inbound mailbox, Reply-To address, or support channel.

### Free-plan capacity snapshot — 2026-11-04

- Resend publishes a Free transactional allowance of 100 emails per day and 3,000 per month. The daily quota uses a UTC calendar day and resets at midnight UTC; it is not a rolling 24-hour window.
- The provider documents that sent and received emails consume the transactional quota, with multiple recipients counted separately. Authentication mail and application invitations share the applicable provider quota; other workloads on the same account may also consume it. No current account usage or available balance has been inspected.
- These published figures are a dated planning snapshot, not a guarantee of available capacity or a substitute for checking provider/account limits before integration. Supabase's own Auth sending limits must also be evaluated separately.
- TD-004 sets an initial global budget of 20 processed password-recovery requests per UTC calendar day, summed across all email addresses and IP addresses, in addition to its per-email and per-IP limits. This is not a guarantee of capacity or a reservation of the remaining provider quota. Capacity allocation across email flows and technical enforcement of the approved quota-exhaustion policy below remain pending.

### Provider quota exhaustion

- When Resend's daily or monthly quota is exhausted, do not promise delivery or automatically retry sends rejected for that reason.
- Record the failure and make it visible to the Platform Administrator through an operational surface that does not depend on sending another email. That surface and its implementation remain to be designed; this decision does not create a new dashboard or notification service.
- Preserve already-created organizations and invitations without extending invitation or authentication-link expiry. Do not report a failed send as successful delivery.
- For password-recovery requests, preserve the approved generic response without revealing whether an account exists or claiming that delivery succeeded. Operational failure visibility is separate from that public response.
- Once capacity is available again, allow explicit manual retry subject to the existing eligibility, expiry, and rate-limit rules. An expired organization invitation requires explicit replacement by a Platform Administrator; an expired authentication link follows the manual new-link request rules in TD-004. Do not send expired links as a delayed retry.
- Do not purchase services or upgrade the plan automatically.
- This policy is approved behavior, not implemented quota monitoring or provider integration. Detection, counters, logging, operational visibility, and retry enforcement remain to be designed and tested.

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
