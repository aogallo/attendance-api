# Product Decisions — Attendance Core

This document records product decisions and separates confirmed behavior, candidate proposals, and unresolved questions. It describes business behavior rather than technical implementation.

## Confirmed product decisions

### 1. The product's primary purpose is attendance

The shared core captures and records attendance across different industries. School, company, gym, and other capabilities may extend the product, but the core must not assume industry-specific business processes.

### 2. A confirmed Attendance record represents a successful capture

Each successful attendance capture creates a separate confirmed Attendance event. An unsuccessful capture is recorded separately as an Attendance Capture Attempt and is not counted as confirmed attendance.

The meaning of a successful event depends on the attendance workflow configured for the organization and the configuration applicable to the person. A workflow may distinguish event types such as work entry, work departure, break start, or break end. These are examples, not a final list of event types supported in the first release.

The system must not infer that a capture is an entry or exit merely because of its position in a sequence or because another event has already been recorded. When the applicable configuration does not assign a specific type, the capture is recorded as an additional attendance event without inventing an entry/exit meaning.

An attendance capture outside the expected schedule can still produce a confirmed Attendance record. Schedule comparison is used for reporting and does not, by itself, reject the capture.

### 3. An organization may enable multiple attendance methods

The platform may support several enabled methods for one organization, such as PIN, fingerprint recognition, and facial recognition. For a given capture, a person uses one method that is enabled for the organization.

Facial recognition is the first attendance method to be delivered and the product's primary attendance method. It is the only attendance method in the first attendance release. PIN and fingerprint recognition are deferred to later releases and must not be offered as available methods or fallbacks in the first release. This scope does not automatically enable facial recognition for an organization or approve a biometric provider, enrollment flow, matching approach, or fallback policy. The core remains capable of supporting additional methods in later releases.

Enabled methods are configured at the organization level and are available across the organization's establishments by default. A single enabled method is the only available method. The policy for assigning or selecting a method when multiple methods are enabled remains unresolved.

### 4. Person registration uses one configured registration mode

An organization chooses one registration mode at a time:

- Administrator-managed registration; or
- Self-registration.

The organization does not enable both modes simultaneously in the current product scope. A self-registered person may record attendance as soon as registration and the required attendance-method enrollment have been completed; a separate administrator approval is not required.

### 5. Person information is isolated by organization

Each organization manages its own people and related attendance information. The platform must not share or automatically link person information across organizations.

### 6. Location validation is optional, and unavailable location does not block attendance

Location validation is used only when the organization enables it. If location information is unavailable in the mobile or web application, the person must still be able to record attendance because attendance capture is the product's primary purpose.

When location is unavailable, the event must not imply that location was successfully validated. The acceptable radius and the behavior when a reported location is outside the permitted area still require definition.

### 7. Every person must have a schedule, and schedule comparison is for reporting

Every person must have an associated schedule within their organization, regardless of their contextual role, such as teacher, employee, or member. The schedule belongs to the person, not to an establishment, group, or activity. Establishment opening hours describe when a location operates and do not replace the person's schedule.

All lateness and schedule-compliance reports must compare confirmed attendance events with the person's applicable schedule and the configured meaning of those events. The first attendance report is intended to identify lateness, early departure, and absence. Person-specific schedule exceptions may be configured for an individual. Being outside the person's schedule does not by itself reject an otherwise successful capture.

The calculation rules remain open. In particular, reports must use the configured meaning of attendance events rather than infer arrival or departure from event sequence alone. Handling of holidays, time zones, overlapping schedules, and person-specific exceptions also needs definition.

### 8. Establishments must exist before person enrollment

An organization must have at least one establishment before a person can enroll. The Organization Administrator is responsible for creating and managing the organization's establishments, including the initial establishment.

### 9. There are two administrative roles

An organization may have multiple Organization Administrators from the first release. The product imposes no maximum number of administrators per organization. This does not change the requirement for each administrator to accept a valid invitation or the organization-scoped boundaries of their authority.

- **Platform Administrator:** creates and configures organizations and invites their Organization Administrators; operates at platform level.
- **Organization Administrator:** manages attendance operations for their own organization, including establishments, people, configuration, and attendance review as permitted by the product.

The detailed permission matrix is not yet defined. Platform-level administration must remain separate from organization-level administration, and organization data must remain isolated.

The [current administrative access policy](security/administrative-access.md) owns sign-in/MFA, password, session, authentication-link, recovery, and audit requirements. It preserves the approved rules and their pending enforcement/operational controls. [TD-004](technical-decisions/004-initial-administrative-authentication.md) records provider rationale; [TD-005](technical-decisions/005-initial-platform-administrator-bootstrap.md) records the restricted initial platform-role assignment. Do not maintain a second copy of those policies here.

Administrative authority and attendance participation are separate:

- The Platform Administrator does not record attendance and has no attendance schedule or Person record for its platform role.
- An Organization Administrator may only administer, without a Person record, attendance enrollment, or schedule.
- An Organization Administrator may also record attendance, which is expected to be a common scenario. The same individual then has an associated Person record within the organization, with a required schedule and the same attendance enrollment, capture, and lateness and schedule-compliance reporting rules as other people. A second login account is not required merely to combine administration and attendance.

Accepting an administrative invitation grants the administrative role; it does not by itself create a Person record, assign a schedule, or enroll the administrator in attendance. The mandatory schedule rule applies to every Person participating in attendance, not to administrative authority alone. If an account already exists for the invited email, the acceptance flow reuses it; otherwise, the invitee can create an account during acceptance. The technical mechanism remains to be specified.

Recovery responsibilities and required audit evidence are maintained in the [lost-authenticator recovery policy](security/administrative-access.md#lost-authenticator-recovery), not implied by administrative attendance participation.

### 10. Organization configuration changes apply prospectively

Changes to an organization's attendance configuration affect future capture attempts. They must not rewrite, reclassify, or delete historical attendance records.

If an attendance method is disabled for an organization, future captures using that method must not be accepted as successful attendance. The behavior for existing method enrollments when a method is disabled and later re-enabled remains unresolved.

### 11. Attendance methods and attendance event types are different concepts

An Attendance Method describes how a person is identified or how a capture is initiated, such as PIN, fingerprint recognition, or facial recognition. An Attendance Event Type describes the purpose of the successful capture when the configured workflow distinguishes it, such as entry, departure, break start, or break end.

The organization and the applicable person configuration determine the attendance workflow. The exact rules for selecting an event type at capture time, and how organization-level and person-level configuration interact, remain to be specified.

### 12. Administrator invitations expire, are single-use, and are replaced manually

Each Organization Administrator invitation is bound to a specific recipient email address and the intended organization. Only the identity corresponding to that recipient email may accept it. Forwarding the link does not authorize an identity with a different email address to gain administrative access. If an account already exists for the invited email, it is reused rather than requiring a duplicate account; otherwise, account creation is supported during acceptance.

Before completing acceptance, the invitee must be authenticated with the account corresponding to the invited email, that email must be verified by the authentication provider, and TOTP MFA must be successfully completed. If no such account exists, the invitee may create it within the same flow and must satisfy those requirements before completing acceptance. Opening the invitation link alone, an unverified matching email, a password-only session, or authentication without an eligible invitation does not grant the invited role. If the invitation expires during email verification or MFA setup/completion, acceptance is rejected without granting the role; a Platform Administrator must explicitly issue a replacement.

The [administrative access policy](security/administrative-access.md#authentication-and-authorization) owns authentication prerequisites. Verification and recovery must not consume an invitation, grant organization permissions, or extend/reactivate invitation validity. This does not approve the separate email verification-code proposal for attendance-method enrollment.

An Organization Administrator invitation is valid for 15 minutes from the creation of its record, not from delivery or the first time its link is opened. After that period, it cannot be accepted. A Platform Administrator must explicitly issue a replacement invitation; the platform does not automatically send another invitation.

Successful acceptance permanently consumes the invitation. It cannot be accepted again, even within its original 15-minute validity period.

Issuing a new invitation for the same recipient email and the same organization replaces all earlier pending Organization Administrator invitations for that combination and makes them unusable. Invitations for other recipient emails or other organizations are unaffected, including invitations for the same email in a different organization. This does not undo an already completed acceptance. Whether replaced invitations are retained with a changed state or deleted is not yet decided; their unusability is required regardless of storage approach.

Only the Platform Administrator may explicitly revoke an Organization Administrator invitation by changing its state. Organization Administrators cannot revoke these invitations. Expiration, successful acceptance, replacement, and explicit revocation are distinct reasons why an invitation can no longer be accepted.

If replacement has already begun, acceptance of the invitation being replaced must be rejected. If acceptance completed before replacement began, the accepted result must not be silently undone. The technical mechanism for enforcing this ordering is not yet selected.

If replacement begins but fails before the new invitation record is created, the prior invitation remains unusable and must not be automatically reactivated. The failure must be clearly shown to the Platform Administrator, who may explicitly retry the replacement. A failed creation does not mean a new invitation exists.

If an invitation record is successfully created but delivery fails, the organization and invitation remain created. The delivery failure must be clearly shown to the Platform Administrator, who may manually retry sending the same invitation while it remains eligible for acceptance. Resending does not create a replacement, restart the 15-minute period, or reactivate an unusable invitation. If the invitation has expired, the Platform Administrator must explicitly issue a replacement instead. This applies to both initial invitations and replacements. Provider values and failure handling are maintained in the [email delivery policy](operations/email-delivery.md); its technical integration remains pending and does not change this lifecycle.

Related policies are maintained once:

- [Authentication-link lifetime and replacement](security/administrative-access.md#authentication-email-link-expiry), distinct from this invitation lifecycle.
- [Recovery privacy](security/administrative-access.md#password-recovery-request-privacy) and [request limits](security/administrative-access.md#password-recovery-request-limits).
- [Email sender/language and delivery responsibilities](operations/email-delivery.md#sender-and-delivery-responsibilities), [capacity snapshot](operations/email-delivery.md#free-plan-capacity-snapshot--2026-11-04), and [quota exhaustion](operations/email-delivery.md#provider-quota-exhaustion).

These are confirmed requirements, not implemented security controls or authorization to configure providers, send emails, or spend money.

### 13. Application failures must be visible and technical failures must be observable

The person taking attendance should receive immediate feedback when an application-level capture or validation fails. Organization Administrators should be able to see relevant application errors for their organization. Technical platform errors should also be available through the platform's operational logging and monitoring tools, such as the monitoring solution selected for deployment.

The exact failure categories, detail shown to Organization Administrators, retention period, and access boundaries still need definition. User-facing messages should communicate the outcome without exposing raw internal diagnostics.

The approved minimum application logging format is structured JSON with date/time, severity level, operation, and request identifier. Exclude passwords, tokens, MFA secrets, and biometric data. Security audit requirements remain separate from diagnostic logging. [TD-006](technical-decisions/006-structured-application-logging.md) records the approved initial logger, Go's standard-library `log/slog` with `JSONHandler`. This does not approve application implementation, Grafana/Loki deployment, or other operational infrastructure.

### 14. Organization creation precedes administrator invitations

The Platform Administrator creates the organization before issuing an invitation that associates an Organization Administrator with it. Organization creation and invitation issuance are separate business actions: an organization may be created without immediately issuing an invitation and may exist before any administrator has accepted one. Invitations may be issued later for that existing organization. This does not grant administrative authority before successful invitation acceptance or automatically make the organization ready for attendance operations.

### 15. The Platform Administrator may edit the organization before invitation acceptance

After creating an organization, the Platform Administrator may edit its business information and attendance configuration before an Organization Administrator accepts an invitation. This includes correcting the organization's name or contact information, changing enabled attendance methods, selecting the person-registration mode, and configuring optional location validation. The required creation fields are confirmed in decision 16; any additional fields remain to be defined.

These organization edits do not change an existing invitation's recipient email or restart its validity period. They do not replace the invitation issuance, acceptance, replacement, or revocation rules. Attendance configuration changes remain prospective and must not rewrite historical attendance. This decision does not define the permission matrix after invitation acceptance.

### 16. Organization creation requires a name and the initial administrator's email

The required business information for creating an organization is its name and the email address of the intended initial Organization Administrator. The email is required even when that administrator's account does not yet exist and the invitation will be issued later.

Saving the intended administrator's email does not create an account, grant administrative authority, create a Person or schedule, or automatically issue or send an invitation. The invitation's 15-minute validity period starts only when its own record is created, not when the organization or intended recipient email is saved. Account creation or reuse remains part of the later invitation acceptance flow.

No additional business-information fields have been approved as required for organization creation. Attendance-profile setup may be completed later, as described in decision 17.

### 17. Attendance configuration may follow organization creation

An organization may be created with its name and intended initial administrator email before its attendance profile is configured. Enabled attendance methods, the person-registration mode, and whether location validation is enabled may be configured during subsequent setup rather than being required at creation.

Before the organization begins recording attendance, that configuration must be completed: at least one supported attendance method is enabled, exactly one person-registration mode is selected, and location validation is explicitly enabled or disabled. An incomplete profile must not be treated as configured by inventing default settings. Requiring this setup does not make location validation mandatory or change the rules allowing captures outside a person's schedule or when location is unavailable.

### 18. Initial attendance uses an administrator-enabled shared-device mode

The first attendance release records attendance only from a shared tablet or other shared device used by the organization. Capture from personal phones or computers is deferred. An ordinary Person does not need a login account or individual sign-in to record attendance: facial recognition identifies the enrolled Person within the authorized organization, and the person's required schedule remains the reporting reference. Administrative access still requires an account and the approved invitation flow.

An Organization Administrator signs in with MFA and activates attendance mode for an organization they are authorized to administer. Attendance mode uses an organization-scoped capture session, not an unrestricted administrative session. Its capture authority must not permit administrative operations. Returning to administrative functions requires administrator reauthentication that satisfies the administrative MFA policy; challenge/recency details remain open. Facial identification for attendance does not itself grant administrative access or replace TOTP MFA.

Attendance mode does not require permanent registration of the device or a permanent association of that device with an organization or establishment. This does not remove the need for an authorized organization context or a relevant establishment context when the configured attendance rules require one. Establishment resolution remains to be specified. Activating attendance mode alone does not prove the device's physical location or change the optional-location and unavailable-location rules.

If a particular organization later requires permanent device registration or binding, treat it as an optional organization-configured capability implemented through an isolated extension, not a universal requirement of the shared attendance core. Its detailed behavior and delivery scope require separate approval. Such extensions must preserve organization isolation and the shared attendance invariants.

## Candidate proposal — not yet approved

### Attendance-method enrollment by email verification code

A proposed approach is to verify the person's email with a confirmation code as part of enrollment. This idea is not yet a final product decision. Email verification can confirm control of an email address, but it does not by itself enroll or validate a PIN, fingerprint, or facial reference. The enrollment flow for each supported method must be defined before implementing that method.

## Intentionally outside the shared core

The following are not part of the shared attendance core's primary responsibility:

- Membership plans, renewals, and membership eligibility.
- Decisions about whether a person is allowed to enter a facility.
- Student enrollment, employee contracts, payroll, class management, fitness plans, and other industry-specific processes.

These capabilities may be implemented as separate business modules or integrations. The core may provide attendance information for those capabilities to consume, but must not absorb their rules by default.

## Open product questions

An implementation agent must not silently choose answers to these questions if the choice would change business behavior.

1. **Method assignment/selection:** When multiple methods are enabled, does the organization assign one method to each person, or can the person choose from the enabled methods? Can the method be changed later?
2. **Method enrollment:** What is the enrollment, confirmation, recovery, replacement, and re-enrollment flow for each method? Is email verification required, and at what point?
3. **Attendance event configuration:** Which event types are supported in the first release? How does a person or device indicate the event type (for example, by choosing an Entry, Exit, Break Start, or Break End option)? When no event type is explicitly assigned, what label or default should be shown?
4. **Organization/person workflow precedence:** The workflow depends on organization and person configuration. Which settings are defined at organization level, which can vary per person, and how are conflicts or overrides resolved?
5. **Location radius and mismatch:** What radius is accepted, where is it configured, and what happens when the reported location is outside the permitted area or too imprecise? When location is unavailable, attendance may proceed, but it must not be represented as location-validated.
6. **Report definitions:** How are lateness, early departure, and absence calculated from the applicable schedule and configured event types? How are holidays, time zones, overlapping schedules, and person-specific exceptions applied?
7. **Failure visibility:** Which categories of application errors and unsuccessful attempts should Organization Administrators see? What details and history are appropriate, and how long should those details be retained?
8. **Administrative permissions:** What exact actions can each role perform, including organization configuration, invitations, attendance-method configuration, person management, attendance review, and account management?
9. **Disabled method enrollment:** When a method is disabled and later re-enabled, are existing enrollments retained and usable again, or must people enroll again?
10. **Establishment lifecycle:** Can an Organization Administrator edit, deactivate, or delete an establishment after people or attendance records are associated with it? Historical attendance should remain interpretable, but the lifecycle rules need confirmation.
11. **Personal schedule assignment:** Who assigns and maintains each person's required schedule, and at what point in registration or enrollment? Can a person have multiple schedules, and how is the applicable schedule selected? These questions do not reopen the requirement that every person must have a schedule or its use as the reporting reference.
12. **Attendance-mode session lifecycle:** How long does a capture session remain valid, how is it renewed or revoked, and what challenge/recency rules apply to administrator reauthentication when returning to administrative functions? The administrative session must satisfy the approved email/password plus TOTP MFA policy; no facial-only administrative unlock is approved. These details must preserve the separation of capture and administrative authority.
13. **Establishment context in attendance mode:** How is the relevant establishment determined for a person and capture when establishment/location rules apply, without requiring permanent device binding? A reference establishment must not be represented as a successfully validated physical location merely because it was selected or associated.

## Guidance for product and implementation work

- Treat confirmed product decisions above and the linked specialized policies as their respective current sources of truth; a proposed specification or plan cannot override them.
- Treat the email verification-code approach as a proposal, not an approved requirement.
- Keep the core independent of any one industry.
- Keep identification method separate from attendance event type.
- Never infer entry, departure, or break events merely from the sequence or count of previous captures.
- Do not add membership management or access-authorization decisions to the shared core without an explicitly approved extension.
- If a feature depends on an unresolved question, raise it during clarification and record the decision before implementation.
- Update this document when its product decisions change. Update administrative and email rules at their linked policy owners, plus affected acceptance criteria/tests, rather than copying their full text here. Follow the [documentation lifecycle](agent-workflow.md#one-owner-per-kind-of-information).
