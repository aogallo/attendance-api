# Platform Administrator

## Purpose

Platform Administrator is the role responsible for administering the attendance platform as a whole.

## Business Context

This role belongs to the platform operator rather than to any particular customer organization. It creates and configures organizations and invites each organization's administrator. Organization Administrators are responsible for creating and managing their organization's establishments.

## Business Purpose

The Platform Administrator supports platform governance and enables customer organizations to begin using the platform.

## Scope

The initial Platform Administrator is provisioned under [TD-005](../technical-decisions/005-initial-platform-administrator-bootstrap.md): the owner creates or reuses a verified Supabase Auth account, completes TOTP MFA, and explicitly runs a restricted bootstrap command to assign the platform role in the application database. Public registration and SQL migrations do not create the real privileged account or grant this role. The command is not an automatic startup seed or an MFA-recovery bypass; its detailed operational safeguards remain to be designed.

Platform administrative access follows the [current administrative access policy](../security/administrative-access.md). Its authentication, MFA, sessions, and recovery requirements do not broaden platform or organization permissions; maintain those cross-cutting rules at the policy owner.

After organization creation and before an Organization Administrator accepts an invitation, the Platform Administrator may correct the organization's business information and adjust its attendance configuration, including enabled methods, person-registration mode, and optional location validation. These edits do not change an existing invitation's recipient email or restart its validity period. Permissions after invitation acceptance remain to be specified.

The role operates at platform level. It is distinct from an Organization Administrator, who manages one specific organization. The Platform Administrator issues an invitation replacement when an invitation expires; replacements are not sent automatically.

Only the Platform Administrator may explicitly revoke an Organization Administrator invitation by changing its state. Organization Administrators cannot revoke these invitations.

Platform Administrator authorization of Organization Administrator recovery is governed by the [lost-authenticator recovery policy](../security/administrative-access.md#lost-authenticator-recovery). It does not authorize recovery of the Platform Administrator's own account; the separate owner/operator responsibility and pending operational controls remain in that policy.

The Platform Administrator does not participate in attendance and has no attendance schedule or Person record for this platform role. Its administrative registration does not enroll it in an organization's attendance process.

Platform-level administration does not automatically mean unrestricted access to every organization's person records or attendance history. Any customer-data access required for support should have a defined purpose and appropriate controls. The exact permission matrix remains to be defined.

## Core Principle

Platform administration and organization administration must remain separate. The term "root" may describe the role informally, but Platform Administrator is the clearer business name for its responsibilities.
