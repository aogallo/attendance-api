# Organization Administrator

## Purpose

Organization Administrator is the role responsible for managing attendance operations for one organization.

## Business Context

An organization may have multiple Organization Administrators from the first release, with no product-defined maximum number. Each administrator receives organization-scoped authority by accepting a valid invitation; supporting multiple administrators does not broaden that authority beyond the organization.

The Organization Administrator creates and manages the organization's establishments, manages people according to the selected registration mode, maintains attendance-related configuration, and reviews attendance information and relevant application errors.

For the first attendance release, the Organization Administrator signs in with MFA and activates attendance mode on a shared device for their authorized organization. This creates a capture-only session, not unrestricted administrative access, and does not require permanently registering the device to an organization or establishment. Returning to administrative functions requires administrator reauthentication that satisfies the MFA policy; the challenge/recency details and session lifecycle remain to be defined. People recording attendance in that mode do not need individual login accounts or TOTP challenges.

## Business Purpose

The Organization Administrator allows a customer to operate its own attendance process without depending on the platform operator for routine organization-level changes.

## Scope

From the first delivery, Organization Administrator access requires email/password authentication plus a TOTP second factor from an authenticator application. The API must reject password-only access to organization-level administrative operations. Invitation acceptance does not waive MFA, and password recovery does not bypass it. If the authenticator is lost, recovery is manual with identity verification and enrollment/verification of a new TOTP factor before administrative access is restored. An authorized recovery invalidates the affected account's previous authentication factors and existing sessions. Only a Platform Administrator may authorize that recovery; an Organization Administrator cannot authorize their own or another administrator's recovery. Detailed evidence, technical enforcement of factor/session invalidation, and enrollment controls remain to be designed. Recovery does not grant access to another organization or platform-wide operations.

The role is limited to the organization it administers. It is responsible for creating the initial and any additional establishments, and an establishment must exist before a person can enroll. It must not administer another organization or platform-wide settings.

Administrative authority does not require attendance participation. An Organization Administrator may only administer, without an associated Person record, attendance enrollment, or schedule.

An Organization Administrator may also participate in attendance, which is expected to be a common scenario. In that case, the same individual has administrative authority and an associated Person record within the organization. That Person must have a schedule and follow the same attendance enrollment, capture, and schedule-compliance reporting rules as other people. This does not require a second login account merely because the individual also records attendance. Administrative privileges do not exempt the person from those rules.

The exact permission matrix—including whether and how it may edit or deactivate establishments, change the organization's profile, view failed attempts, and manage administrator accounts—still needs to be defined.

## Core Principle

An Organization Administrator must operate only within the scope of the assigned organization. Administrative authority over one organization must not grant access to another organization's records or settings.
