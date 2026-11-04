# Attendance

## Purpose

Attendance represents a confirmed business record created when a person's attendance capture succeeds through an organization's configured attendance process.

## Business Context

Attendance is used across different industries. A school may record a student's presence, a company may record an employee's arrival or departure, and a gym may record a member's visit. The organization's configured attendance workflow and the settings applicable to the person determine the meaning assigned to a successful capture when event types are used.

Examples of event types may include entry, departure, break start, and break end. An organization may also choose a simpler flow in which every successful capture is recorded as another attendance event without inferring entry or departure from the sequence.

## Business Purpose

Attendance provides a consistent record of successful captures that an organization can review over time. Depending on the configured workflow and business context, the record may support attendance reports, schedule comparisons, visit history, or other attendance-related analysis.

## Scope

A successful capture creates a confirmed Attendance event. An unsuccessful capture does not create a confirmed Attendance record; it belongs in the attendance capture attempt log so that the failed attempt can be reviewed separately.

The meaning of an event must come from the configured workflow or an explicit event-type selection, not from the number or order of previous captures alone. The identification method (such as PIN, fingerprint, or facial recognition) is distinct from the event type (such as entry, departure, or break start).

Attendance does not itself define the identification method, the establishment, the applicable schedule, or an industry-specific relationship such as student enrollment, employment, or gym membership. It also does not decide whether a person is entitled to enter a facility. Access control and membership validation are outside the shared core unless separately introduced as a distinct extension.

## Core Principle

Attendance must represent successfully registered attendance consistently across industries. Each successful capture is an event; the product must not invent its business meaning from capture sequence alone. A failed capture or a separate business decision must not be confused with a confirmed attendance record.
