# Attendance Platform — Product and Domain Documentation

This documentation describes the business meaning, decisions, and principal use cases of the shared attendance platform. It intentionally avoids technical implementation details unless a separate implementation specification requires them.

## Start here

- [Product Decisions](product-decisions.md) — source of truth for established decisions and unresolved product questions.
- [Principal Attendance Flow](use-cases/principal-flow.md) — end-to-end flow from organization setup to attendance review.

## Core concepts

- [Person](core/person.md) — an individual whose attendance may be recorded.
- [Organization](core/organization.md) — the customer context in which attendance is managed.
- [Organization Profile](core/organization-profile.md) — the organization's configured attendance capabilities and registration paths.
- [Establishment](core/establishment.md) — a physical location associated with an organization.
- [Attendance](core/attendance.md) — a successful attendance record.
- [Attendance Method](core/attendance-method.md) — how a person is identified or a capture is initiated.
- [Attendance Event Type](core/attendance-event-type.md) — the purpose assigned to a successful capture when the configured workflow distinguishes event types.
- [Attendance Capture Attempt](core/attendance-attempt.md) — an unsuccessful attempt that did not produce a confirmed attendance.
- [Schedule](core/schedule.md) — each person's required planned or expected attendance time.
- [Schedule Applicability](core/schedule-applicability.md) — personal schedules as the reporting reference, distinct from establishment and opening-hour contexts.
- [Location Validation](core/location-validation.md) — optional validation of presence near an establishment.

## Platform roles

- [Platform Administrator](actors/platform-administrator.md) — administers the platform as a whole.
- [Organization Administrator](actors/organization-administrator.md) — administers one organization.

## Documentation principles

1. Keep the shared core focused on attendance.
2. Keep industry-specific concepts outside the core unless they are necessary to describe attendance consistently.
3. Keep a person's records isolated within each organization; do not share or automatically link person information across organizations.
4. Distinguish confirmed attendance from unsuccessful capture attempts.
5. Treat attendance methods, registration paths, event types, and location validation as configurable rather than universal assumptions. Every person must have an associated schedule; establishment opening hours do not replace it.
6. Keep identification method separate from attendance event type. Never infer entry, departure, or break type only from the sequence or count of captures.
7. All lateness and schedule-compliance reports use the person's applicable schedule. Schedule comparison is for reporting and does not block an otherwise successful attendance capture.
8. Organization configuration changes apply prospectively and must not rewrite historical attendance. A disabled method cannot be used for future successful captures.
9. Do not treat recording attendance as the same thing as deciding whether someone is allowed to enter.
10. When a business decision is unresolved, record it in `product-decisions.md` and clarify it before implementation. Do not silently invent product behavior.
