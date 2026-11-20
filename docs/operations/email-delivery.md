# Transactional Email — Current Policy

**Status:** Confirmed delivery requirements; no account, DNS, sender, provider, or real delivery configuration is established by this document.
**Owner:** This page owns current email delivery values and failure rules. [TD-007](../technical-decisions/007-transactional-email-delivery.md) preserves provider/integration rationale. [Administrative access policy](../security/administrative-access.md) owns authentication-link lifetime, replacement, privacy, and recovery limits; [product decisions](../product-decisions.md#12-administrator-invitations-expire-are-single-use-and-are-replaced-manually) own organization-invitation lifecycle rules.

## Sender and delivery responsibilities

- Use Resend Free initially. Paid upgrades or additional spending require separate explicit authorization; selecting the plan does not authorize account access or configuration.
- Use neutral, professional Spanish as the initial language for organization-invitation, email-verification, and password-recovery emails. This approves their language only; exact templates, content, and delivery configuration remain pending. It does not select a localization mechanism or change the language of API contracts or technical documentation.
- Use `notificaciones.allangallo.dev` as the sending subdomain. This uses the owner's current domain and therefore exposes the owner's personal domain in sender addresses rather than a separate product domain.
- The subdomain must be added and verified for sending before use. The owner's earlier Resend setup with a different domain does not establish verification of this subdomain. No verification or DNS configuration has been performed as part of this design decision.
- Configure Supabase Auth to use that provider through custom SMTP for email verification and password recovery. Supabase remains responsible for those authentication flows.
- The API sends application-owned organization invitations through a separate integration with the same provider. The API integration's transport and implementation remain pending; custom SMTP for Auth does not automatically deliver application invitations.
- Delivery failure preserves an already-created organization and invitation, is clearly reported to the Platform Administrator, and allows manual resend of the same eligible invitation without extending expiry. An expired invitation requires explicit replacement.
- Use `no-reply@notificaciones.allangallo.dev` as the sender address and `Attendance` as the display name for the initial authentication and organization-invitation emails. These are approved design values, not evidence of a configured sender, verified domain, or working delivery.
- Credentials, account/plan verification, templates, detailed delivery integration, and operational safeguards remain pending. This sender identity decision does not define an inbound mailbox, Reply-To address, or support channel.

## Free-plan capacity snapshot — 2026-11-04

- Resend publishes a Free transactional allowance of 100 emails per day and 3,000 per month. The daily quota uses a UTC calendar day and resets at midnight UTC; it is not a rolling 24-hour window.
- The provider documents that sent and received emails consume the transactional quota, with multiple recipients counted separately. Authentication mail and application invitations share the applicable provider quota; other workloads on the same account may also consume it. No current account usage or available balance has been inspected.
- These published figures are a dated planning snapshot, not a guarantee of available capacity or a substitute for checking provider/account limits before integration. Supabase's own Auth sending limits must also be evaluated separately.
- The [administrative access policy](../security/administrative-access.md#password-recovery-request-limits) owns the global recovery budget and its per-email/per-IP constraints. It is not a guarantee of capacity or a reservation of the remaining provider quota. Capacity allocation across email flows and technical enforcement of the approved quota-exhaustion policy below remain pending.

## Provider quota exhaustion

- When Resend's daily or monthly quota is exhausted, do not promise delivery or automatically retry sends rejected for that reason.
- Record the failure and make it visible to the Platform Administrator through an operational surface that does not depend on sending another email. That surface and its implementation remain to be designed; this decision does not create a new dashboard or notification service.
- Preserve already-created organizations and invitations without extending invitation or authentication-link expiry. Do not report a failed send as successful delivery.
- For password-recovery requests, preserve the approved generic response without revealing whether an account exists or claiming that delivery succeeded. Operational failure visibility is separate from that public response.
- Once capacity is available again, allow explicit manual retry subject to the existing eligibility, expiry, and rate-limit rules. An expired organization invitation requires explicit replacement by a Platform Administrator; an expired authentication link follows the manual new-link request rules in the [administrative access policy](../security/administrative-access.md#authentication-email-link-expiry). Do not send expired links as a delayed retry.
- Do not purchase services or upgrade the plan automatically.
- This policy is approved behavior, not implemented quota monitoring or provider integration. Detection, counters, logging, operational visibility, and retry enforcement remain to be designed and tested.

## Maintenance and references

Update operational values and dated capacity checks here rather than copying them into feature plans or ADRs. A provider-selection change with significant consequences requires an approved decision record; a quota observation is not automatically a new architecture decision.

- [Resend pricing](https://resend.com/pricing).
- [Resend account quotas and limits](https://resend.com/docs/knowledge-base/account-quotas-and-limits).
- [Supabase custom SMTP](https://supabase.com/docs/guides/auth/auth-smtp).
