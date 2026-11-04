# Organization Profile

## Purpose

Organization Profile describes how an organization is configured to use the attendance platform, including its enabled attendance methods, person-registration mode, attendance workflow, and optional location validation.

## Business Context

Attendance-profile configuration is not required to create the organization. The organization may first be created with its name and intended initial administrator email; enabled methods, person-registration mode, and the location-validation choice may be configured later. Before recording attendance, at least one supported method must be enabled, exactly one registration mode selected, and location validation explicitly enabled or disabled. An incomplete profile must not be presented as configured or silently assigned unapproved defaults.

Different organizations may need different ways to register people and take attendance. An organization may enable one or more supported methods, such as facial recognition, PIN, or fingerprint recognition. The organization selects one person-registration mode at a time: administrator-managed registration or self-registration.

Facial recognition is the first method to be delivered and the product's primary attendance method. It is the only attendance method in the first attendance release, but is not an automatically enabled organization setting. PIN and fingerprint recognition are deferred to later releases and cannot be configured as available methods or fallbacks in the first release. Organization creation can still precede attendance-profile setup; facial recognition must be enabled before recording attendance in that release.

Enabled methods are configured at the organization level and apply across its establishments by default. Location validation is optional. The attendance workflow may distinguish event types such as entry, departure, break start, and break end; the actual supported event types and selection process remain to be defined. Every person must have an associated schedule, which belongs to the person rather than an establishment and is the reference for all lateness and schedule-compliance reports. The schedule assignment process and the remaining boundaries between organization-level and person-level workflow configuration are still open.

## Business Purpose

Organization Profile allows the platform to adapt to each customer's attendance needs without creating a separate product or changing the shared core for every industry.

## Scope

Organization Profile describes attendance-related configuration. It does not define industry-specific operations such as school enrollment, employment administration, gym memberships, access control, or payroll.

The first attendance release uses administrator-enabled attendance mode on an organization shared device, with an authorized organization-scoped capture session and no individual login required for an ordinary Person. Permanent device registration or binding to an organization or establishment is not a baseline requirement. If an organization later requires it, it is an optional organization-configured capability implemented as an isolated extension, without making it mandatory for every organization or weakening core isolation rules. Its detailed behavior is not yet approved for implementation.

An attendance method describes how a person is identified or how a capture begins; an attendance event type describes the purpose of a successful capture when the configured workflow distinguishes it. Schedule comparison is for reporting and a schedule mismatch alone does not prevent an otherwise successful attendance capture.

Changes to an organization's attendance profile apply to future capture attempts. They must not rewrite, reclassify, or delete historical attendance records. A method disabled in the current profile cannot be used for a future successful capture. The handling of existing enrollments when a method is disabled and later re-enabled remains to be defined.

## Core Principle

Each organization can configure a supported attendance experience that matches its operating needs, while shared attendance concepts remain consistent across organizations and industries. Configuration changes must not alter historical attendance records.
