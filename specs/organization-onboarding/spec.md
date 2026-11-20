# Organization Onboarding — Consolidated Behavior Specification

**Status:** Accepted by the owner. The owner also approved the companion plan and authorized local implementation within its scope, subject to the remaining gates below.
**Approval evidence:** “apruebo la espec, y el plan para autorizar la implementacion”.
**Companion:** [Approved implementation and test plan](implementation-plan.md). Remaining design and operational approvals are listed separately below.

## Outcome and boundary

A Platform Administrator can create an organization independently of invitation issuance, configure it before administrator acceptance, and invite one or more Organization Administrators. An invitee obtains organization-scoped authority only after authenticated, verified-email, MFA-completed acceptance of an eligible invitation.

This specification describes the complete onboarding behavior in the existing intent. The owner approved the smaller implementation slices in the companion plan. No requirement here authorizes attendance implementation, production setup, or execution of privileged operations.

## Sources and precedence

- [Confirmed product decisions](../../docs/product-decisions.md) and [original intent](intent.md).
- [Current administrative access policy](../../docs/security/administrative-access.md) and [current email delivery policy](../../docs/operations/email-delivery.md) own cross-cutting requirements. The intent is historical clarification context, not a second current rulebook.
- [Platform Administrator](../../docs/actors/platform-administrator.md) and [Organization Administrator](../../docs/actors/organization-administrator.md).
- [Administrative authentication](../../docs/technical-decisions/004-initial-administrative-authentication.md), [initial bootstrap](../../docs/technical-decisions/005-initial-platform-administrator-bootstrap.md), and [email delivery](../../docs/technical-decisions/007-transactional-email-delivery.md).

Confirmed source behavior is restated as acceptance criteria below. More specific approved policies take precedence over older general lists of pending work. Open questions are not resolved by this document; a newly discovered source conflict requires owner resolution, not a silent assumption.

## Actors and preconditions

| Actor | Boundary |
| --- | --- |
| Platform Administrator | Organization creation, invitations, explicit replacement/revocation, and approved pre-acceptance edits; requires authenticated platform authority and completed TOTP MFA. |
| Invitee | May reuse or create the account for the invited email, verify email, complete MFA, and explicitly accept an eligible invitation. Authentication alone grants no role. |
| Organization Administrator | Receives only the intended organization authority after acceptance. Multiple administrators are supported without a product-defined maximum. |
| Owner/operator | Restricted initial platform-role provisioning, separately from onboarding; no public privilege grant or recovery bypass. |

Before operating protected onboarding, the initial platform account must be provisioned under TD-005. This is not a requirement to configure remote services before writing or locally testing the implementation.

## Main flow

1. The Platform Administrator authenticates with email/password and TOTP MFA.
2. They create an organization with its name and intended initial administrator email. The account need not exist. Invitation issuance and attendance configuration may be deferred.
3. They may edit organization information and attendance configuration before invitation acceptance. Existing invitation recipients and validity are not changed by those edits.
4. They explicitly issue an invitation for the existing organization and intended email. Its original 15-minute validity starts at invitation-record creation.
5. The invitee uses the matching account, verifies the invited email, and successfully completes TOTP MFA. These steps do not extend invitation validity.
6. The invitee explicitly accepts. Eligibility is checked at completion; role assignment and permanent invitation consumption succeed together or neither occurs.
7. The account becomes an Organization Administrator only for that organization. No Person, schedule, establishment, method enrollment, or attendance record is created by acceptance.

## Organization and configuration criteria

| ID | Required observable behavior |
| --- | --- |
| OB-001 | Creating an organization requires its name and intended initial administrator email. Reject missing required information without creating the organization. No additional required business field is approved. |
| OB-002 | Creation succeeds even when the intended administrator account does not exist and no invitation is issued. Saving the intended email does not create an identity, invitation, administrative assignment, Person, or schedule. |
| OB-003 | Attendance configuration is optional at creation. Preserve unset configuration rather than silently selecting defaults or claiming attendance readiness. |
| OB-004 | Before attendance is enabled in a later feature, configuration requires at least one supported method, exactly one registration mode, and an explicit location-validation choice. First attendance delivery supports facial recognition only; it is not automatically enabled. |
| OB-005 | Registration modes are administrator-managed registration or self-registration. Optional location validation does not become a universal attendance requirement. No attendance operation is implemented by this feature. |
| OB-006 | Platform-authorized pre-acceptance edits preserve existing invitation recipient and original validity. Configuration changes are prospective and must not rewrite historical attendance. Post-acceptance editing permissions remain an approval gate if included. |

## Identity, permissions, and acceptance criteria

| ID | Required observable behavior |
| --- | --- |
| OB-007 | Reject unauthenticated, password-only, or insufficiently authorized platform operations without their requested organization/invitation mutations. Configuring a factor without successfully verifying MFA is insufficient. |
| OB-008 | Validate authenticated identity and session assurance; do not trust a supplied email, organization identifier, or role as authority. Enforce organization isolation on every applicable operation. |
| OB-009 | An Organization Administrator cannot acquire platform authority or another organization's authority through onboarding. A platform role does not imply unrestricted access to customer attendance records. |
| OB-010 | Support multiple Organization Administrators with independently accepted eligible invitations, without inventing a product-defined cap. |
| OB-011 | Acceptance reuses an existing account for the invited email or permits account creation when absent. Do not require a duplicate account. |
| OB-012 | Acceptance requires the matching authenticated account, provider-verified invited email, and successfully completed TOTP MFA. Reject anonymous, wrong-email, unverified-email, or password-only acceptance without role assignment or invitation consumption. Forwarding a link does not authorize a different email identity. |
| OB-013 | Authentication, email verification, and MFA completion alone do not accept an invitation or assign administrative roles. An account may sign in for TOTP setup without administrative access. |
| OB-014 | If invitation validity ends during account creation, verification, or MFA enrollment, reject acceptance at completion without granting authority. A Platform Administrator must explicitly replace an expired invitation. |
| OB-015 | Successful acceptance assigns the invited organization role and permanently consumes the invitation as one consistent outcome. Repeated or concurrent acceptance must not grant duplicate authority. |
| OB-016 | Acceptance does not create attendance participation records. A Platform Administrator has no Person or attendance schedule; an Organization Administrator may later participate through a separate organization-scoped Person with a required personal schedule. |

## Invitation lifecycle and failure criteria

| ID | Required observable behavior |
| --- | --- |
| OB-017 | An invitation is organization/email-bound and valid for 15 minutes from its record creation, not delivery or first opening. Reject expired, consumed, replaced, or revoked invitations. |
| OB-018 | Only a Platform Administrator may explicitly revoke or replace these invitations. Organization Administrators cannot revoke them. Do not automatically replace expired invitations. |
| OB-019 | New issuance for the same organization/email replaces all earlier pending invitations for that combination. Preserve other organizations, recipient emails, and completed acceptances. |
| OB-020 | If acceptance completed before replacement began, preserve the accepted result. If replacement began first, reject acceptance of a predecessor even if successor creation is still pending. |
| OB-021 | If replacement begins but successor creation fails, the predecessor stays unusable. Clearly report failure, allow explicit retry, and do not claim a successor invitation exists. A transaction rollback must not reactivate the predecessor. |
| OB-022 | Delivery failure after record creation preserves the organization and invitation, clearly reports the failure to the Platform Administrator, and does not imply acceptance or successful delivery. This applies to initial and replacement invitations. |
| OB-023 | Manual resend uses the same still-eligible invitation and does not restart validity, create a replacement, or reactivate an unusable invitation. An expired invitation requires explicit replacement. |

## Supporting administrative access and email criteria

These criteria verify the linked current policies; they do not independently own password/session/email values. Update affected criteria and tests when the policy owner changes an approved requirement, rather than editing a value here alone. Which components are delivered in each implementation slice is proposed in the companion plan; mocks are not proof of actual provider enforcement.

| ID | Required observable behavior |
| --- | --- |
| OB-024 | Verification and password-recovery links last 30 minutes. An expired link produces a clear expiry notice and permits a manual new-link request, without automatic sending or invitation extension. |
| OB-025 | Successful new authentication-link issuance invalidates earlier links for the same account and purpose and gives the new link 30 minutes. A mere request, rejection, or failed issuance preserves only predecessor original eligibility/expiry; no extension/reactivation. Other purposes and invitations are unaffected. Issuance is not delivery proof. |
| OB-026 | Recovery requests for existing and nonexistent accounts use the same response: “Si existe una cuenta asociada a ese correo, recibirás instrucciones para recuperar tu contraseña”. Include manual new-link requests; do not expose account existence through errors or claim delivery. Technical HTTP/timing protections still require design. |
| OB-027 | Recovery processing satisfies every approved limit: at least 60 seconds apart and at most 5 per email per rolling hour; at most 30 per IP per rolling hour across emails; at most 20 globally per UTC calendar day, reset at midnight UTC. Apply the same treatment regardless of account existence. Limits do not block accounts/sign-in or trigger extra email. |
| OB-028 | Recovery limits do not silently become verification, invitation-resend, or sign-in limits. Trusted IP determination, normalization, failure counting, and enforcement require a concrete design before those capabilities are implemented. |
| OB-029 | Both roles use the approved password baseline: minimum 15 characters, long passphrases, no mandatory character-type mixtures or periodic changes, change upon compromise evidence, managers/autofill/paste permitted, common/compromised passwords rejected. Enforcement and a free-compatible blocking mechanism remain to be verified. |
| OB-030 | Ordinary sign-out invalidates only the current administrative session. Session access expires after 30 minutes of inactivity or an 8-hour maximum, whichever comes first; fresh email/password and TOTP are required afterward. Refresh alone cannot bypass either limit. |
| OB-031 | Explicit navigation, manual queries, and saving changes reset inactivity, not the 8-hour maximum. Mouse movement, an open tab, background queries/updates, and automatic renewals do not reset inactivity or revive an expired session. Trusted tracking remains an engineering prerequisite. |
| OB-032 | Initial verification, recovery, and organization-invitation emails use neutral professional Spanish and the approved sender Attendance <no-reply@notificaciones.allangallo.dev>. Authentication mail and application invitations retain separate ownership even when delivered through the same provider. |
| OB-033 | On provider daily/monthly quota exhaustion, do not promise delivery, automatically retry rejected sends, extend validity, or upgrade/pay automatically. Record failure and expose it to the Platform Administrator without relying on another email. Preserve records and recovery privacy; permit only explicit eligible retry when capacity returns, never delayed expired links. |
| OB-034 | Diagnostics use the approved structured format and exclude passwords, tokens, MFA secrets, and biometric data. Diagnostic logging does not satisfy security-audit persistence by itself. |
| OB-035 | Initial platform-role assignment follows the restricted TD-005 procedure only after trusted account verification and completed MFA. No public self-elevation, real privileged-user migration seed, startup grant, or recovery reuse. Repeated invocation must not silently regrant privileges. Detailed command authorization and replay policy require approval before implementation. |
| OB-036 | If lost-factor recovery execution is included, preserve approved manual identity verification, authorizing actor, invalidation of prior factors/all sessions, new verified TOTP before restored access, and audit fields without secrets. It cannot ship using invented identity evidence or bootstrap as a bypass. If deferred by the owner, do not expose a reset path. |

## Scenario expansion and proof

- Expand every criterion into tests of all named observable effects: output, success/failure status, role assignments, invitation state, expiry, and absence of unrelated mutations. Preserve existing behavior when a shared path changes.
- Exercise both account cases and exact invitation boundary/order cases. Use an injected clock for deterministic timing and isolated PostgreSQL state for real persistence/concurrency tests.
- Compare hashes or equivalent snapshots of business state after rejected commands. Technical security counters may change only as explicitly specified by the eventual counting design; do not assert that all state is immutable on every rejection.
- Race acceptance against acceptance and replacement using deterministic coordination. Inject successor-creation and delivery failures; inspect persisted results, not just messages.
- Keep native provider integration evidence separate from local adapter doubles. No tests or integration checks have been run merely by creating this specification.

## Approval gates and unresolved details

| ID | Question or proposal | Blocking boundary |
| --- | --- | --- |
| G-01 | Resolved: the owner approved this consolidated specification, the companion plan and its incremental boundary, and authorized local implementation. | Approval recorded; remaining gates still apply. |
| G-02 | Proposed: first delivery includes only approved platform-controlled configuration before administrator acceptance; post-acceptance edits require separately defined permissions. | Before implementing any additional post-acceptance operation. |
| G-03 | Decide whether actual lost-factor recovery execution ships inside this delivery. Proposed: defer execution until identity evidence, operational authorization, and audit controls are approved; fail closed meanwhile. | Before exposing any recovery-reset operation, not automatically before organization-creation coding. |
| G-04 | Approve restricted bootstrap authorization, identity/MFA evidence, and repeat-execution behavior; specify sensitive-operation reauthentication recency where required. | Before those security-sensitive mechanisms are implemented or operated. |
| G-05 | Adopt concrete abuse policies for other publicly exposed authentication/email flows and approve applicable audit access/retention/failure rules. Recovery thresholds are not a substitute. | Before exposing the affected flows. |
| G-06 | Complete API/error schemas, SQL constraints and ordering, normalization, counters, trusted activity/IP handling, provider success semantics, test tools and versions. These are engineering proposals in the plan, not additional confirmed product rules. | Before implementing the corresponding behavior; raise only consequential policy choices to the owner. |

Maximum password length/Unicode handling, exact email copy and redirect destinations, and an operational failure-viewing surface are not invented here. They must be specified and verified at their relevant integration boundary.

## Explicit exclusions

- Establishments, people, schedules, biometric enrollment, attendance capture/reporting, event-type semantics, payroll, memberships, and access authorization.
- Shared-device attendance mode and its capture-session lifecycle, PIN/fingerprint delivery, and permanent device registration.
- Client v1 migration, customer-operated installations, and enterprise identity federation.
- Grafana/Loki deployment, production hosting/provisioning, DNS/provider configuration, credentials, paid services, and real emails.

## Review and delivery state

This is an accepted consolidated specification with stable acceptance IDs and owner authorization for local implementation under the approved plan. It is not an approved release or an assertion that any implementation/security control exists. Resolve remaining applicable gates before coding the affected slice. Any owner correction updates the affected criteria and plan without reopening unrelated confirmed decisions.
