# Platform Administrator

## Purpose

Platform Administrator is the role responsible for administering the attendance platform as a whole.

## Business Context

This role belongs to the platform operator rather than to any particular customer organization. It creates and configures organizations and invites each organization's administrator. Organization Administrators are responsible for creating and managing their organization's establishments.

## Business Purpose

The Platform Administrator supports platform governance and enables customer organizations to begin using the platform.

## Scope

The initial Platform Administrator is provisioned under [TD-005](../technical-decisions/005-initial-platform-administrator-bootstrap.md): the owner creates or reuses a verified Supabase Auth account, completes TOTP MFA, and explicitly runs a restricted bootstrap command to assign the platform role in the application database. Public registration and SQL migrations do not create the real privileged account or grant this role. The command is not an automatic startup seed or an MFA-recovery bypass; its detailed operational safeguards remain to be designed.

From the first delivery, Platform Administrator access requires email/password authentication plus a TOTP second factor from an authenticator application. A password-only session must not perform platform-level administrative operations; the API must enforce the requirement, not merely hide UI controls. Password recovery does not bypass MFA. If the authenticator is lost, recovery is reserved to the platform owner/operator through a separate restricted operational procedure, with identity verification and enrollment/verification of a new TOTP factor before administrative access is restored. This is not public self-service recovery or reuse of the initial bootstrap. Detailed procedure authorization, evidence, technical enforcement of factor/session invalidation, and initial enrollment controls remain to be defined. Organization Administrators must also use TOTP MFA for their own organization-scoped administrative operations; MFA does not broaden either role's permissions.

After organization creation and before an Organization Administrator accepts an invitation, the Platform Administrator may correct the organization's business information and adjust its attendance configuration, including enabled methods, person-registration mode, and optional location validation. These edits do not change an existing invitation's recipient email or restart its validity period. Permissions after invitation acceptance remain to be specified.

The role operates at platform level. It is distinct from an Organization Administrator, who manages one specific organization. The Platform Administrator issues an invitation replacement when an invitation expires; replacements are not sent automatically.

Only the Platform Administrator may explicitly revoke an Organization Administrator invitation by changing its state. Organization Administrators cannot revoke these invitations.

Only a Platform Administrator with valid MFA-authenticated access may authorize an Organization Administrator's lost-authenticator recovery, after verifying the account holder's identity. The affected administrator must verify a new TOTP factor before administrative functions are restored. This permission does not authorize recovery of the Platform Administrator's own account; that separate case follows the restricted owner/operator procedure above. For either administrative role, an authorized recovery invalidates the affected account's previous authentication factors and existing sessions. Detailed identity evidence, execution, technical enforcement of factor/session invalidation, and audit controls remain to be designed.

The Platform Administrator does not participate in attendance and has no attendance schedule or Person record for this platform role. Its administrative registration does not enroll it in an organization's attendance process.

Platform-level administration does not automatically mean unrestricted access to every organization's person records or attendance history. Any customer-data access required for support should have a defined purpose and appropriate controls. The exact permission matrix remains to be defined.

## Core Principle

Platform administration and organization administration must remain separate. The term "root" may describe the role informally, but Platform Administrator is the clearer business name for its responsibilities.
