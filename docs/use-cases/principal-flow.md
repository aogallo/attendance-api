# Principal Attendance Flow

## Purpose

Describe the main business flow from organization setup to attendance review. This document captures business behavior, not implementation details.

## Actors

- **Platform Administrator:** creates and configures organizations and invites each organization's administrator.
- **Organization Administrator:** manages establishments, people, attendance configuration, and attendance review for their own organization.
- **Person:** an individual registered with an organization who uses an enabled attendance method to record attendance.

The Platform Administrator has no attendance or schedule. An Organization Administrator may only administer or may also participate as a Person within the organization. In the latter case, the same individual has a required personal schedule and follows the same attendance and reporting rules as other people, without needing a second login account merely to combine these responsibilities.

## Attendance-operation preconditions

These describe readiness for the attendance operations below, not requirements for initially creating the organization. Organization creation and administrator invitations may precede completion of attendance-profile setup.

- The organization exists and has an attendance profile configured.
- At least one attendance method is enabled for the organization.
- Exactly one person-registration mode is selected for the organization: administrator-managed registration or self-registration.
- Location validation may be enabled or disabled.
- At least one establishment exists before a person begins enrollment.

## Main Flow

Administrative operations in this narrative require the identity, MFA, and permission boundaries in the [current administrative access policy](../security/administrative-access.md). Recovery requirements and pending controls are maintained there, not in this flow. Ordinary Persons do not need individual login or TOTP for authorized capture-only sessions; this narrative does not expand attendance implementation scope.

1. **Create and configure the organization.**
   The Platform Administrator supplies the organization's name and the email address of the intended initial Organization Administrator; both are required business information. The account need not exist yet. Saving the email does not create an account, grant permissions, create a Person or schedule, or automatically issue an invitation. The invitation's 15-minute validity period starts when its own record is created, not when the organization is created.

   Organization creation precedes administrator invitations and does not require issuing one immediately. The Platform Administrator may invite administrators later, associating them with the existing organization. Organization existence alone does not imply readiness for attendance operations.

   Enabled attendance methods, person-registration mode, and the location-validation choice may be configured during subsequent setup rather than at creation. Before recording attendance, at least one supported method must be enabled, exactly one registration mode selected, and location validation explicitly enabled or disabled. Incomplete configuration must not be silently filled with unapproved defaults. Location validation remains optional, and the existing schedule-mismatch and unavailable-location rules still apply.

   Before an Organization Administrator accepts an invitation, the Platform Administrator may edit the organization's business information and attendance configuration, including enabled methods, person-registration mode, and optional location validation. These edits do not change an existing invitation's recipient email or restart its validity period. Additional organization fields and permissions after invitation acceptance remain to be defined.

   The Platform Administrator creates the organization and configures its attendance profile. The organization chooses which attendance methods are enabled, selects one registration mode, and decides whether location validation is enabled. Enabled methods apply across the organization's establishments by default. The organization configures the attendance workflow, and person-specific settings may further determine which event types apply to an individual. Every person must have their own associated schedule; it does not belong to an establishment. The schedule assignment process and the remaining division between organization and person settings remain to be specified.

2. **Invite the Organization Administrator.**
   An organization may have multiple Organization Administrators from the first release, with no product-defined maximum number. The invitation and acceptance flow applies to each administrator independently and does not grant authority over other organizations.

   The invitation is bound to the intended organization and a specific recipient email address. Only the identity corresponding to that email may accept it. Forwarding the link to an identity with a different email does not authorize administrative access.

   Before completing acceptance, the invitee must be authenticated with the account corresponding to the invited email, that email must be verified by the authentication provider, and TOTP MFA must be successfully completed. If the account does not exist, it can be created within the same flow and must satisfy those requirements before completing acceptance. Opening the invitation link alone, a matching but unverified email, a password-only session, or authentication without an eligible invitation does not grant the administrative role. Email verification and MFA setup/completion do not extend or restart the invitation's original 15-minute validity period from record creation. If it expires during those steps, acceptance is rejected without granting the role; a Platform Administrator must explicitly issue a replacement.

   Authentication-link, password, session, and recovery behavior is confirmed in the [administrative access policy](../security/administrative-access.md); its enforcement details remain pending. Those authentication flows do not replace invitation acceptance or alter its original validity. Delivery values and quota failures follow the [email policy](../operations/email-delivery.md).

   If an account already exists for the invited email, the acceptance flow reuses it without requiring a duplicate account. If no account exists for that email, the invitee can create one during acceptance. Both cases grant the administrative role only through successful invitation acceptance and do not automatically create a Person record, schedule, or attendance-method enrollment. The provider and initial access flow are approved in TD-004; detailed account integration remains to be designed.

   The Platform Administrator invites the designated Organization Administrator. The invitation is valid for 15 minutes from the creation of its record, not from delivery or opening the link. The Platform Administrator must explicitly issue a replacement invitation; the system does not issue one automatically.

   Successful acceptance permanently consumes the invitation. Expired, consumed, replaced, or explicitly revoked invitations cannot be accepted. Issuing a new invitation for the same recipient email and the same organization replaces all earlier pending invitations for that combination and makes them unusable, without affecting other recipient emails, other organizations, or completed acceptances. Only the Platform Administrator may explicitly revoke an Organization Administrator invitation by changing its state. Organization Administrators cannot revoke these invitations.

   If replacement has already begun, acceptance of that invitation must be rejected. Acceptance completed before replacement began must not be silently undone. The technical ordering mechanism remains to be specified.

   If replacement begins but fails before the new invitation record is created, the prior invitation remains unusable and is not automatically reactivated. The failure is clearly shown to the Platform Administrator, who may explicitly retry the replacement. The failed creation must not be presented as a newly created invitation.

   If an initial or replacement invitation record is created but delivery fails, the organization and invitation remain created. Show the failure clearly to the Platform Administrator, who may manually resend the same invitation while it remains eligible for acceptance. Resending does not restart the original 15-minute period or reactivate an unusable invitation. If it has expired, the Platform Administrator must explicitly issue a replacement. The technical delivery mechanism remains undecided.

   Accepting the invitation grants administrative authority only. It does not automatically create a Person record, assign a schedule, or enroll the administrator in attendance. An administrator who also needs attendance participates through the person registration and enrollment process below.

3. **Create establishments.**
   The Organization Administrator creates the organization's establishments, including the initial establishment. At least one establishment must exist before a person can enroll. Location validation uses the relevant establishment when configured and applicable.

4. **Register a person and enroll an attendance method.**
   The person is registered using the one registration mode selected for the organization:
   - In administrator-managed mode, the Organization Administrator registers the person.
   - In self-registration mode, the person registers themselves.

   The person must complete enrollment for one attendance method enabled for the organization. A self-registered person can record attendance once registration and required method enrollment are complete; a separate Organization Administrator approval is not required. The exact method assignment/selection and enrollment steps remain to be specified.

5. **Activate attendance mode and capture attendance.**
   For the first attendance release, an Organization Administrator signs in with MFA and activates attendance mode on a shared tablet or other organization shared device. The resulting capture session is scoped to an organization that administrator is authorized to administer and cannot perform administrative operations. Returning to administrative functions requires administrator reauthentication that satisfies the administrative MFA policy; challenge/recency details and the capture-session lifecycle remain to be defined. Captures by ordinary people do not require individual login or TOTP challenges.

   Permanent device registration or binding to an organization or establishment is not required. A future requirement for a particular organization belongs to an optional organization-configured extension, not the universal shared core. How establishment context is resolved when needed remains open; attendance mode alone does not establish that physical location was validated.

   The ordinary Person is identified by facial recognition without an individual login account or sign-in. Personal-device capture is deferred. Facial attendance identification does not itself grant administrative authority. A person becoming an administrator follows the separate invitation/account flow, without replacing their Person or attendance history.

   Each person must have an associated schedule. This requirement does not define a new enrollment approval or capture-validation step; the assignment process and timing remain open.

   The person uses one attendance method enabled for the organization. The configured attendance workflow and the settings applicable to the person determine the meaning of a capture when the workflow uses typed events.

   Examples of configurable event types may include entry, departure, break start, and break end. The identification method and event type are separate: for example, facial recognition may identify the person while an explicit button or selection indicates that the capture is a departure. The exact event-type options and how they are selected remain to be defined.

   - Each successful capture creates a separate confirmed Attendance event.
   - The system must not infer entry or departure solely from the number or order of previous captures.
   - If the applicable configuration does not distinguish event types, the capture is recorded as another attendance event without automatically labeling it as entry or departure.
   - If the capture or an applicable required validation fails, a separate unsuccessful Attendance Capture Attempt is recorded; it is not counted as confirmed attendance.
   - If location information is unavailable, attendance can still be recorded. The result must not imply that location was successfully validated.
   - Being outside the expected schedule does not by itself reject a successful capture.
   - A method disabled in the organization's current configuration cannot be used to produce a successful attendance capture.
   - The person receives immediate application-level feedback when a capture fails. Relevant application errors are available to the Organization Administrator; technical platform errors are available through operational logging and monitoring.

6. **Review attendance and schedule comparisons.**
   The Organization Administrator reviews confirmed attendance for their organization. All lateness and schedule-compliance reports must compare confirmed events with the person's applicable schedule and configured event meanings, not with establishment opening hours. The first report is intended to identify lateness, early departure, and absence. The report calculations remain to be specified, including how holidays, time zones, overlapping schedules, and person-specific exceptions are handled.

## Established Behavior

- The product's primary purpose is recording attendance, not managing membership eligibility or deciding whether a person may enter a facility.
- Initial capture uses facial recognition in administrator-enabled attendance mode on organization shared devices. Ordinary people need no individual login account. The capture session is organization-scoped and has no administrative authority; returning to administrative functions requires administrator reauthentication.
- Permanent device registration or organization/establishment binding is not required by the baseline. Organization-specific requirements may be added later through optional configured extensions, preserving core invariants. Personal-device capture is deferred.
- An organization may enable multiple identification/capture methods. A person uses one enabled method for a given capture; enabled methods apply across establishments by default.
- The meaning of an attendance event depends on the configured workflow and the settings applicable to the person.
- Each successful capture creates a separate event. The system does not infer entry, departure, or break type solely from capture order or count.
- An organization selects one registration mode at a time: administrator-managed registration or self-registration.
- A self-registered person can take attendance after registration and method enrollment are complete, without separate administrator approval.
- The Organization Administrator creates all establishments, and an establishment must exist before person enrollment.
- Location validation is optional. If location is unavailable, attendance can still be recorded, but it must not be described as location-validated.
- Confirmed attendance and unsuccessful capture attempts are separate records.
- Schedule comparison is for reporting and does not reject attendance by itself.
- Every person, including a teacher or member, must have an associated schedule. It belongs to the person, not the establishment, and is the reference for all lateness and schedule-compliance reports.
- Organization configuration changes apply to future captures and do not rewrite historical attendance records. A disabled attendance method cannot be used for future successful captures.
- Person information is isolated by organization.
- An Organization Administrator may only administer without a Person record or schedule. If also participating in attendance, the administrator has an associated organization-scoped Person with a required schedule and the same attendance and reporting rules as other people. The Platform Administrator has no attendance or schedule.
- An expired Organization Administrator invitation must be replaced manually by the Platform Administrator.
- Invitation validity starts at record creation and lasts 15 minutes. Acceptance is single-use; replacement or explicit revocation also makes the prior invitation unusable. Acceptance is rejected if replacement has already begun, while an acceptance completed earlier is not silently undone.
- Application-level errors provide immediate feedback and are visible to the relevant Organization Administrator; technical platform errors are sent to operational logging and monitoring.

## Scope Notes

- This is the principal flow, not a complete description of every exception or alternate path.
- Membership management, payroll calculation, and access-authorization decisions are outside the shared attendance core unless introduced through a separately approved extension.
- This flow does not specify database structure, API endpoints, credential storage, biometric processing details, or implementation architecture.
- Email confirmation by code is a candidate enrollment approach, not an approved requirement for every attendance method.

## Open Product Questions

See `docs/product-decisions.md` for the central decision log. The following questions most directly affect this flow:

1. When multiple attendance methods are enabled, who assigns or selects the method for a person, and can it change later?
2. What is the enrollment, replacement, and recovery process for each method?
3. Which event types will be supported in the first release, and how will the person/device indicate the intended event type?
4. Which attendance settings belong to the organization and which can vary by person? How are conflicts resolved?
5. What radius and behavior apply when the reported location is outside the permitted area or too imprecise?
6. What exact calculation rules define lateness, early departure, and absence, including holidays, time zones, overlapping schedules, and person-specific exceptions?
7. Which application errors or failed attempts can Organization Administrators review, with what detail and retention period?
8. What exact permissions belong to each administrative role?
9. What happens to a person's existing enrollment when an organization disables and later re-enables that method?
10. Can establishments be edited, deactivated, or deleted after being used?
11. Who assigns and maintains each person's required schedule, when is it assigned, and how is the applicable schedule selected if multiple schedules are supported?
12. What are the capture-session expiry, renewal, revocation, and administrator reauthentication rules for attendance mode?
13. How is establishment context resolved when applicable, without requiring permanent device binding or falsely asserting physical-location validation?

## Acceptance Summary

The principal flow is supported when an organization can be configured with one registration mode and one or more attendance methods; a Platform Administrator can invite an Organization Administrator; that administrator can create establishments and manage people; every person has an associated schedule and can enroll in one enabled method and record attendance; each successful capture creates a distinct event whose type is only assigned according to the applicable configuration; unsuccessful captures are represented separately; unavailable location does not block attendance or falsely indicate successful location validation; all lateness and schedule-compliance reports use the person's applicable schedule rather than establishment opening hours, without schedule mismatches blocking attendance; configuration changes apply to future captures without changing historical attendance; and an Organization Administrator can review only their organization's information.
