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
- The initial administrative sign-in method is email and password. Do not silently substitute passwordless or social sign-in as the baseline.
- Both Platform Administrator and Organization Administrator access require MFA from the first delivery: email/password plus a TOTP code from an authenticator application. An email/password-only session must not execute operations authorized by either administrative role. Enforce this in the API using the verified authentication assurance of the session, not only through the login UI. This does not select SMS or facial recognition as the second factor.
- Organization Administrators must satisfy MFA before completing invitation acceptance and before exercising organization-level administrative authority, including activating attendance mode. Ordinary Person captures through an already authorized capture-only session do not require individual accounts or TOTP challenges. Detailed enrollment controls and reauthentication recency remain to be designed.
- Password recovery alone must not bypass either administrative role's MFA requirement. [TD-005](005-initial-platform-administrator-bootstrap.md) approves the initial platform-role assignment after verified account and completed MFA. Lost-authenticator recovery follows the approved responsibilities and manual process below; detailed verification and operational controls still need design. No fallback that grants administrative authority without satisfying the requirement is approved.
- Before completing an Organization Administrator invitation acceptance, the account must be authenticated, its invited email verified by the authentication provider, and TOTP MFA successfully completed. A matching but unverified email or a password-only session is insufficient to grant the invited role. The invitation must still be eligible when acceptance completes, within its original 15-minute validity period from record creation; email verification and MFA setup/completion do not extend or restart that period.
- Support forgotten-password recovery through a link sent to the account email, using Supabase Auth rather than implementing a separate credential-recovery system. The application does not store a duplicate administrative password.
- Account email confirmation and password recovery are separate from the application's administrative invitation. Their links do not consume the invitation, grant organization permissions, extend its original 15-minute validity period, or reactivate an unusable invitation. The invitation must still be eligible when acceptance completes.
- Ordinary Persons do not need a Supabase account or individual login merely to enroll in attendance or record it in the shared-device mode. Facial attendance identification is not Supabase administrative authentication.
- Keep organizations, administrative assignments, permissions, invitation records and lifecycle rules, Person data, schedules, and attendance in the application domain. Do not require one Supabase project or native provider organization for each customer organization.
- Existing-account reuse and account creation during invitation acceptance must follow the confirmed invited-email rules. Authentication or account creation alone must not grant the administrative role or create attendance-related records.
- The API must validate the authenticated identity before applying its own authorization rules; a client-supplied email or role is not proof of identity or authority.
- Do not substitute the provider's invitation lifecycle for the application's approved 15-minute, single-use, replacement, revocation, and resend behavior.
- TD-001 continues to govern the application-owned PostgreSQL schema and Flyway migrations. Do not directly manage the provider's internal authentication tables as application tables or assume that business data must be moved to the provider's database.
- Keep provider-specific authentication integration outside business rules, consistent with TD-002 and the deferred customer-operated scenario in TD-003. This does not require implementing multiple providers now.
- Begin development using the free plan if its current limits and terms fit the environment. Production plan selection, usage limits, costs, infrastructure, and email delivery require separate assessment and explicit authorization before spending or provisioning. Do not treat a pricing snapshot as a permanent guarantee.

### Email delivery responsibilities

- [TD-007](007-transactional-email-delivery.md) selects one managed transactional email provider, with custom SMTP for Supabase Auth verification/password-recovery messages and a separate API integration for application-owned organization invitations.
- TD-007 selects Resend, the sending subdomain `notificaciones.allangallo.dev`, and the sender `Attendance <no-reply@notificaciones.allangallo.dev>`. Contracting, domain verification, sender configuration, quotas, pricing, and detailed integration remain pending. This does not authorize service configuration, real email delivery, or purchases, and does not replace the application's invitation lifecycle with provider authentication invitations.

### Authentication email-link expiry

- Email-verification and password-recovery links have a validity period of 30 minutes. This applies to administrative accounts of either role.
- Use 1800 seconds as the intended Supabase Email OTP Expiration setting, which is shared by these authentication email-link flows. This records the approved target configuration, not an applied or verified provider setting.
- Application-owned Organization Administrator invitations retain their separate 15-minute validity from invitation-record creation. The shared provider setting does not control or extend those invitations.
- If an email-verification or password-recovery link has expired, clearly inform the person that it has expired and allow an explicit manual request for a new link. Do not automatically send another link upon expiry or when the expired link is opened.
- Requesting or sending a new authentication link does not extend, restart, replace, or reactivate an application-owned organization invitation. An expired invitation still requires explicit replacement by a Platform Administrator.
- After successful issuance of a new email-verification or password-recovery link, earlier links for the same account and same purpose become unusable. The new link has 30 minutes of validity from its successful issuance; links for other accounts or purposes and application-owned organization invitations are unaffected.
- A mere request, a rate-limited request, or a failed issuance is not a successful replacement. In those cases, an earlier link retains only its original eligibility and expiry; it must not be extended or reactivated if already unusable.
- Successful issuance is not proof of mailbox delivery. The exact provider-backed success condition and enforcement of previous-link invalidation must be designed and verified during integration; this is an approved requirement, not an asserted guarantee about the deployed provider.
- Verification or password recovery does not grant administrative roles, waive TOTP MFA, or extend invitation validity. Remaining message copy, rate-limit and abuse-protection enforcement, and provider integration still require design. This records approved behavior, not an implemented resend flow or authorization to send real emails.

### Password-recovery request privacy

- Password-recovery requests must not reveal whether an account exists for the submitted email. Use the same user-facing response for existing and nonexistent accounts: “Si existe una cuenta asociada a ese correo, recibirás instrucciones para recuperar tu contraseña”.
- This conditional response does not claim that a message has actually been sent or delivered. It applies to requests for password-recovery instructions, including manual requests for a new recovery link, not to the already-approved notice that an opened link has expired.
- This is an approved product rule, not an implemented account-enumeration defense. HTTP status/body details, timing behavior, provider-error handling, request rate limits, and abuse protections remain to be designed and tested so they do not undermine the rule.
- This does not change the invitation lifecycle or allow password recovery to bypass TOTP MFA.

### Password-recovery request limits

- **Approved:** 2026-11-04.
- For each submitted email address, process password-recovery requests at least 60 seconds apart, with a maximum of 5 processed requests in a rolling one-hour window. This includes manual requests for a new recovery link.
- Apply the same limits and generic user-facing response whether or not an account exists for the email address. The response remains “Si existe una cuenta asociada a ese correo, recibirás instrucciones para recuperar tu contraseña”.
- Reaching either limit prevents additional recovery processing and email sending for the limited requests; it must not block the account's sign-in or send additional emails.
- Additionally, process a maximum of 30 password-recovery requests per IP address in a rolling one-hour window, summed across all submitted email addresses. Both the per-email and per-IP limits must be satisfied; the IP limit does not replace the 60-second interval or 5-request per-email limit.
- At the IP limit, stop processing additional recovery requests from that IP without sending additional recovery emails, blocking accounts, or blocking sign-in. Preserve the same generic user-facing response whether or not the submitted email has an account.
- The IP threshold is an initial approved value to validate against real usage, not a measured capacity or provider default. One IP may represent an entire school or company network; this shared-network impact must inform validation.
- Additionally, process a maximum of 20 password-recovery requests globally per UTC calendar day, summed across all submitted email addresses and IP addresses. The daily window resets at midnight UTC; it is not a rolling 24-hour window. The global, per-IP, and per-email limits must all be satisfied.
- The global threshold is an initial approved budget for Resend Free, not a guarantee of available email capacity or a reservation of the remaining quota for other flows. Other sending, provider monthly quotas, and Supabase limits may reduce actual available capacity. TD-007 records the approved provider-quota exhaustion policy: no delivery promise, automatic retry, or automatic paid upgrade; make the failure operationally visible without relying on another email, retain the generic recovery response, and allow manual retry when capacity returns without extending expiry. Capacity allocation and technical enforcement remain pending.
- This defines a recovery-request policy, not sign-in attempt limits, invitation resend limits, or email-verification request limits. Reaching any recovery limit stops additional recovery processing and email sending for the limited requests, preserves the generic response, and does not block accounts or sign-in.
- Address normalization, trusted client-IP determination, counter storage and enforcement, failure/counting semantics, HTTP response details, and provider integration remain implementation-design work. No counter, provider setting, or live sending behavior has been implemented or changed by this approval.

### Administrative password policy

- Apply the same baseline to Platform Administrator and Organization Administrator accounts: a minimum of 15 characters, allowing long passphrases.
- Do not require a mixture of uppercase letters, lowercase letters, digits, and symbols.
- Do not require periodic password changes. Require a change when there is evidence that the password has been compromised.
- Permit password managers, autofill, and pasting passwords.
- Reject common or known-compromised passwords. The blocklist mechanism, configuration, and operational cost remain pending; this requirement does not authorize a paid plan or external service.
- Supabase's documented native leaked-password protection requires Pro or higher. Recheck availability and pricing before selecting or provisioning it; no such purchase or configuration is approved here.
- Detailed enforcement, maximum supported length, Unicode/length handling, error behavior, and compromise-response execution still require design and verification. The approved policy is not evidence that provider settings, UI controls, or API enforcement have been implemented.

### Sign-in before MFA completion

- An account may sign in with email and password to configure and verify its TOTP factor before completing MFA. This applies to both administrative roles and does not grant administrative access.
- Until MFA is successfully completed, the API must reject operations requiring either administrative role. Creating or configuring a factor without successfully verifying it is insufficient.
- Completing MFA satisfies the authentication requirement only; it does not assign roles, accept an invitation, or expand organization permissions. Initial Platform Administrator assignment still follows TD-005.
- Verified invited email and successfully completed TOTP MFA are prerequisites for completing invitation acceptance, not substitutes for it. If the invitation expires while the invitee verifies email or sets up/completes MFA, acceptance must be rejected without granting the invited role; an explicit replacement by a Platform Administrator is required. Neither step extends invitation validity or authorizes a self-service bypass of lost-factor recovery.

### Ordinary sign-out scope

- For either administrative role, ordinary sign-out invalidates only the current session. Other sessions of the same account, including sessions on other devices, remain unaffected.
- This is distinct from authorized lost-authenticator recovery, which invalidates all existing sessions of the affected account under the recovery rules below.
- This defines the required behavior, not a verified provider default or implemented revocation mechanism. API enforcement, token invalidation timing, and renewal mechanics still require design; approved administrative session limits are specified below.
- This decision does not define the separate lifecycle of capture-only attendance sessions.

### Administrative session limits

- For both administrative roles, the session expires after 30 minutes of inactivity or at a maximum duration of 8 hours, whichever occurs first.
- Explicit administrator actions reset the 30-minute inactivity timer: navigating between sections, manually querying information, or saving changes. Moving the mouse, leaving a tab open, receiving automatic updates, automatic token renewals, and background queries do not reset it.
- These actions do not reset or extend the 8-hour maximum session duration. After either expiry limit is reached, activity must not revive the expired administrative session without the required reauthentication.
- After either limit is reached, administrative access requires fresh email/password authentication and successful TOTP MFA. A token refresh alone must not bypass either limit or substitute for that reauthentication.
- These are approved requirements, not evidence that provider configuration or API enforcement already exists. Trusted tracking of the approved user activities, timing enforcement, token validation/renewal mechanics, and provider configuration still require design and verification before implementation.
- These limits apply to administrative sessions, not to the separately undecided lifecycle of capture-only attendance sessions.

### Lost authenticator recovery

- Every recovery must have an audit record identifying who authorized it, the affected account, the date/time, and the outcome. Do not record passwords, tokens, or TOTP secrets. This applies to recovery of either administrative role, including the restricted owner/operator procedure.
- For either administrative role, an authorized recovery invalidates the affected account's previous authentication factors and existing sessions. Administrative access requires enrollment and successful verification of a new TOTP factor; neither an old factor nor a pre-recovery session may restore that access.
- Recovery after loss of the TOTP authenticator is manual and requires verification of the account holder's identity.
- The administrator must enroll and successfully verify a new TOTP factor before administrative functions are restored. Recovery must not grant an MFA exemption or substitute a password/email-recovery link for the required second factor.
- This applies to both administrative roles. It restores access to the existing account and its authorized role scope; it does not grant new roles or access to other organizations.
- Only a Platform Administrator may authorize lost-authenticator recovery for an Organization Administrator, after verifying that account holder's identity. This is a platform administrative action and requires the authorizing administrator's own valid MFA-authenticated access. An Organization Administrator cannot approve their own recovery or another administrator's recovery.
- Platform Administrator lost-authenticator recovery is reserved to the platform owner/operator through a separate restricted operational procedure. It requires verification of the account holder's identity and enrollment and successful verification of a new TOTP factor before administrative access is restored. This responsibility is separate from authorization of Organization Administrator recovery and does not permit public self-service recovery, an MFA exemption, or reuse of the initial bootstrap command.
- Acceptable identity evidence, how the restricted procedure is authorized and executed, technical enforcement of factor/session invalidation, audit storage/access/retention and failure handling, and abuse protections remain to be approved. The audit fields and secret-exclusion rule above are approved; their implementation is not. Factor and session invalidation is an approved rule, not a verified provider capability or completed implementation. The owner/operator responsibility does not by itself define sufficient identity evidence or operational credentials.
- This approval defines the general recovery rule only; do not invoke privileged provider operations or reset any factor without the separate operational authorization and approved procedure.

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
